package rest

import (
	"errors"
	"net/http"
	"os"

	"github.com/pericles-luz/go-parcelamos-tudo/internal/model"
	"github.com/pericles-luz/go-parcelamos-tudo/internal/model/response"
)

var (
	ErrMissingEngine      = errors.New("missing rest engine")
	ErrSubscriptionFailed = errors.New("subscription failed")
	ErrPlanCreationFailed = errors.New("plan creation failed")

	ErrAuthenticationRequired   = errors.New("authentication required")
	ErrMissingAutenticationData = errors.New("missing authentication data")
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
	Get(payload map[string]interface{}, link string) (IResponse, error)
	Delete(link string) (IResponse, error)
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

func (r *Rest) getLink(link string) string {
	return r.baseLink + link
}

func (r *Rest) SetToken(token *Token) error {
	return r.engine.SetToken(token)
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
	result, err := r.engine.PostWithHeaderNoAuth(r.authentication.ToMap(), r.getLink("/auth/token"), map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	})
	if err != nil {
		return err
	}
	authenticationResponse := response.NewAuthentication()
	if err := authenticationResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
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
	result, err := r.engine.Post(subscription.ToMap(), r.getLink("/api/subscription"))
	if err != nil {
		return nil, err
	}
	if result.GetCode() != http.StatusCreated {
		return nil, ErrSubscriptionFailed
	}
	subscriptionResponse := response.NewSubscription()
	if err := subscriptionResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		return nil, err
	}
	return subscriptionResponse, nil
}

// scope: subscription.delete
func (r *Rest) Unsubscribe(subscriptionID string) (*response.SubscriptionDelete, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	result, err := r.engine.Delete(r.getLink("/api/subscription/" + subscriptionID))
	if err != nil {
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		return nil, ErrSubscriptionFailed
	}
	subscriptionDeleteResponse := response.NewSubscriptionDelete()
	if err := subscriptionDeleteResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
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
	result, err := r.engine.Post(plan.ToMap(), r.getLink("/plans"))
	if err != nil {
		return nil, err
	}
	if result.GetCode() != http.StatusCreated {
		return nil, ErrPlanCreationFailed
	}
	planResponse := response.NewPlan()
	if err := planResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		return nil, err
	}
	return planResponse, nil
}

// scope: plan.search
func (r *Rest) ListPlan(page, pageSize uint16) (*response.PlanList, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	if pageSize == 0 {
		pageSize = 10
	}
	result, err := r.engine.Get(map[string]interface{}{
		"page":      page,
		"page_size": pageSize,
	}, r.getLink("/api/plan"))
	if err != nil {
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		return nil, ErrSubscriptionFailed
	}
	planListResponse := response.NewPlanList()
	if err := planListResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		return nil, err
	}
	return planListResponse, nil
}

// scope: plan.read
func (r *Rest) GetPlan(planID string) (*response.Plan, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	result, err := r.engine.Get(nil, r.getLink("/api/plan/"+planID))
	if err != nil {
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		return nil, ErrSubscriptionFailed
	}
	planResponse := response.NewPlan()
	if err := planResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		return nil, err
	}
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
	result, err := r.engine.Post(card.ToMap(), r.getLink("/cards"))
	if err != nil {
		return nil, err
	}
	if result.GetCode() != http.StatusCreated {
		return nil, ErrSubscriptionFailed
	}
	cardResponse := model.NewCard()
	if err := cardResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		return nil, err
	}
	return cardResponse, nil
}

// scope: card.read
func (r *Rest) GetCard(cardID string) (*model.Card, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	result, err := r.engine.Get(nil, r.getLink("/api/card/"+cardID))
	if err != nil {
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		return nil, ErrSubscriptionFailed
	}
	cardResponse := model.NewCard()
	if err := cardResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		return nil, err
	}
	return cardResponse, nil
}
