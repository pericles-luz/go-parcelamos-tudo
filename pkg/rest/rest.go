package rest

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/pericles-luz/go-base/pkg/utils"
	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model"
	"github.com/pericles-luz/go-parcelamos-tudo/pkg/model/response"
)

var (
	ErrMissingEngine       = errors.New("missing rest engine")
	ErrSubscriptionFailed  = errors.New("subscription failed")
	ErrPlanCreationFailed  = errors.New("plan creation failed")
	ErrPlanNotFound        = errors.New("plan not found")
	ErrCardCreationFailed  = errors.New("card creation failed")
	ErrCardRetrievalFailed = errors.New("card retrieval failed")

	ErrAuthenticationRequired   = errors.New("authentication required")
	ErrMissingAutenticationData = errors.New("missing authentication data")
	ErrSubscriptionIDIsRequired = errors.New("subscription id is required")
)

type IResponse interface {
	GetCode() int
	GetRaw() string
}

type IToken interface {
	SetValidity(validity string) error
	SetKey(key string)
	IsValid() bool
	GetValidity() string
	GetKey() string
}

type IEngine interface {
	SetToken(token *Token) error
	NeedAutenticate() bool
	Post(payload map[string]interface{}, link string) (IResponse, error)
	PostWithHeaderNoAuth(payload map[string]interface{}, link string, header map[string]string) (IResponse, error)
	PostWithHeader(payload map[string]interface{}, link string, header map[string]string) (IResponse, error)
	Get(payload map[string]interface{}, link string) (IResponse, error)
	GetWithHeader(payload map[string]interface{}, link string, header map[string]string) (IResponse, error)
	Delete(link string) (IResponse, error)
	DeleteWithHeader(link string, header map[string]string) (IResponse, error)
}

type Rest struct {
	engine   IEngine
	baseLink string

	authentication *model.Authentication
}

func NewRest(engine IEngine, credentials string, scope []string) (*Rest, error) {
	if engine == nil {
		return nil, ErrMissingEngine
	}
	configJSON, err := os.ReadFile(credentials)
	if err != nil {
		return nil, err
	}
	autentication := model.NewAuthentication()
	for _, s := range scope {
		autentication.AddScope(s)
	}
	if err := autentication.Unmarshal(configJSON); err != nil {
		return nil, err
	}
	return &Rest{
		engine:         engine,
		authentication: autentication,
		baseLink:       autentication.Link,
	}, nil
}

func (r *Rest) SetBaseLink(baseLink string) {
	r.baseLink = baseLink
}

func (r *Rest) GetLink(link string) string {
	return r.baseLink + link
}

func (r *Rest) SetToken(token *Token) error {
	return r.engine.SetToken(token)
}

func (r *Rest) showResponse(result IResponse) {
	if result == nil {
		fmt.Println("Result is nil")
		return
	}
	fmt.Println("Code: ", result.GetCode())
	fmt.Println("Raw: ", result.GetRaw())
}

func (r *Rest) Authenticate() error {
	if r.engine == nil {
		return ErrMissingEngine
	}
	if !r.engine.NeedAutenticate() {
		return nil
	}
	if r.authentication == nil {
		return ErrMissingAutenticationData
	}
	if err := r.authentication.Validate(); err != nil {
		return err
	}
	result, err := r.engine.PostWithHeaderNoAuth(r.authentication.ToMap(), r.GetLink("/auth/token"), map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	})
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	authenticationResponse := response.NewAuthentication()
	if err := authenticationResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return err
	}
	token := NewToken(authenticationResponse)
	return r.SetToken(token)
}

// scope: subscription.create
func (r *Rest) Subscribe(subscription *model.Subscription) (*response.Subscription, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	if err := subscription.Validate(); err != nil {
		return nil, err
	}
	if r.engine.NeedAutenticate() {
		return nil, ErrAuthenticationRequired
	}
	fmt.Println("Subscription: ", string(utils.MapInterfaceToBytes(subscription.ToMap())))
	fmt.Println("Link: ", r.GetLink("/api/subscription"))
	result, err := r.engine.PostWithHeader(subscription.ToMap(), r.GetLink("/api/subscription"), map[string]string{
		"api-version": "1",
	})
	if err != nil {
		return nil, err
	}
	if result.GetCode() != http.StatusCreated {
		r.showResponse(result)
		return nil, ErrSubscriptionFailed
	}
	subscriptionResponse := response.NewSubscription()
	if err := subscriptionResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	return subscriptionResponse, nil
}

// scope: subscription.read
func (r *Rest) GetSubscription(subscriptionID string) (*model.Subscription, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	if subscriptionID == "" {
		return nil, ErrSubscriptionIDIsRequired
	}
	result, err := r.engine.GetWithHeader(nil, r.GetLink("/api/subscription/"+subscriptionID), map[string]string{
		"api-version": "1",
	})
	if err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		r.showResponse(result)
		return nil, ErrSubscriptionFailed
	}
	subscription := model.NewSubscription()
	if err := subscription.Unmarshal([]byte(result.GetRaw())); err != nil && !errors.Is(err, model.ErrCustomerIsRequired) {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	return subscription, nil
}

// scope: subscription.delete
func (r *Rest) Unsubscribe(subscriptionID string) (*response.SubscriptionDelete, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	result, err := r.engine.DeleteWithHeader(r.GetLink("/api/subscription/"+subscriptionID), map[string]string{
		"api-version": "1",
	})
	if err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		r.showResponse(result)
		return nil, ErrSubscriptionFailed
	}
	subscriptionDeleteResponse := response.NewSubscriptionDelete()
	if err := subscriptionDeleteResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	return subscriptionDeleteResponse, nil
}

// scope: plan.create
func (r *Rest) CreatePlan(plan *model.Plan) (*response.Plan, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	if err := plan.Validate(); err != nil {
		return nil, err
	}
	result, err := r.engine.PostWithHeader(plan.ToMap(), r.GetLink("/api/plan"), map[string]string{
		"api-version": "1",
	})
	if err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	if result.GetCode() != http.StatusCreated {
		r.showResponse(result)
		return nil, ErrPlanCreationFailed
	}
	planResponse := response.NewPlan()
	if err := planResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	return planResponse, nil
}

// scope: plan.search
func (r *Rest) ListPlan(page, pageSize uint16, externalID string) (*response.PlanList, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	if pageSize == 0 {
		pageSize = 10
	}
	params := map[string]interface{}{
		"page":      page,
		"page_size": pageSize,
	}
	if externalID != "" {
		params["external_reference_id"] = externalID
	}
	result, err := r.engine.GetWithHeader(params, r.GetLink("/api/plan"), map[string]string{
		"api-version": "1",
		"Accept":      "application/json",
	})
	if err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		r.showResponse(result)
		return nil, ErrSubscriptionFailed
	}
	planListResponse := response.NewPlanList()
	if err := planListResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	return planListResponse, nil
}

// scope: plan.read
func (r *Rest) GetPlan(planID string) (*response.Plan, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	result, err := r.engine.GetWithHeader(nil, r.GetLink("/api/plan/"+planID), map[string]string{
		"api-version": "1",
	})
	if err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		r.showResponse(result)
		return nil, ErrSubscriptionFailed
	}
	planResponse := response.NewPlan()
	planResponse.Success = true
	if err := planResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	return planResponse, nil
}

// scope: plan.read
func (r *Rest) GetPlanByExternalID(planID string) (*response.Plan, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	mapping := map[string]interface{}{
		"external_reference_id": planID,
	}
	result, err := r.engine.GetWithHeader(mapping, r.GetLink("/api/plan"), map[string]string{
		"api-version": "1",
	})
	if err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		r.showResponse(result)
		return nil, ErrSubscriptionFailed
	}
	planList := response.NewPlanList()
	if err := planList.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	if len(planList.Data) == 0 {
		return nil, ErrPlanNotFound
	}
	if len(planList.Data) > 1 {
		return nil, ErrPlanNotFound
	}
	planResponse := response.NewPlan()
	planResponse.Success = true
	planResponse.ID = planList.Data[0].ID
	return planResponse, nil
}

// scope: card.create
func (r *Rest) CreateCard(card *model.Card) (*model.Card, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	if err := card.Validate(); err != nil {
		return nil, err
	}
	result, err := r.engine.PostWithHeader(card.ToMap(), r.GetLink("/api/card"), map[string]string{
		"api-version": "1",
	})
	if err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	if result.GetCode() != http.StatusCreated {
		r.showResponse(result)
		return nil, ErrCardCreationFailed
	}
	cardResponse := model.NewCard()
	cardResponse.Number = card.Number
	cardResponse.CVV = card.CVV
	cardResponse.ExpirationMonth = card.ExpirationMonth
	cardResponse.ExpirationYear = card.ExpirationYear
	if err := cardResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	return cardResponse, nil
}

// scope: card.read
func (r *Rest) GetCard(cardID string) (*model.Card, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	result, err := r.engine.GetWithHeader(nil, r.GetLink("/api/card/"+cardID), map[string]string{
		"api-version": "1",
	})
	if err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		r.showResponse(result)
		return nil, ErrCardRetrievalFailed
	}
	cardResponse := model.NewCard()
	if err := cardResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	return cardResponse, nil
}

// scope: webhook.create
func (r *Rest) CreateWebhook(webhook *model.Webhook) (*response.Webhook, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	if err := webhook.Validate(); err != nil {
		return nil, err
	}
	result, err := r.engine.PostWithHeader(webhook.ToAPI(), r.GetLink("/api/webhook"), map[string]string{
		"api-version": "1",
	})
	if err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	if result.GetCode() != http.StatusCreated {
		r.showResponse(result)
		return nil, ErrSubscriptionFailed
	}
	webhookResponse := response.NewWebhook()
	if err := webhookResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	return webhookResponse, nil
}
