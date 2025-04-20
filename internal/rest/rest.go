package rest

import (
	"errors"

	"github.com/pericles-luz/go-parcelamos-tudo/internal/model"
)

var (
	ErrMissingEngine      = errors.New("missing rest engine")
	ErrSubscriptionFailed = errors.New("subscription failed")
	ErrPlanCreationFailed = errors.New("plan creation failed")
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
	Post(payload map[string]interface{}, link string) (IResponse, error)
	PostWithHeaderNoAuth(payload map[string]interface{}, link string, header map[string]string) (IResponse, error)
	Delete(link string) (IResponse, error)
}

type Rest struct {
	engine   IEngine
	baseLink string
}

func NewRest(engine IEngine) *Rest {
	return &Rest{
		engine: engine,
	}
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

func (r *Rest) Authenticate(auth *model.Authentication) error {
	if err := auth.Validate(); err != nil {
		return err
	}
	if r.engine == nil {
		return ErrMissingEngine
	}
	response, err := r.engine.PostWithHeaderNoAuth(auth.ToMap(), r.getLink("/auth/token"), map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	})
	if err != nil {
		return err
	}
	authResponse := model.NewAuthenticationReturn()
	if err := authResponse.Unmarshal([]byte(response.GetRaw())); err != nil {
		return err
	}
	token := NewToken(authResponse)
	return r.SetToken(token)
}

func (r *Rest) Subscribe(subscription *model.Subscription) (*model.SubscriptionResponse, error) {
	if r.engine == nil {
		return nil, ErrMissingEngine
	}
	if err := subscription.Validate(); err != nil {
		return nil, err
	}
	response, err := r.engine.Post(subscription.ToMap(), r.getLink("/subscriptions"))
	if err != nil {
		return nil, err
	}
	if response.GetCode() != 201 {
		return nil, ErrSubscriptionFailed
	}
	subscriptionResponse := model.NewSubscriptionResponse()
	if err := subscriptionResponse.Unmarshal([]byte(response.GetRaw())); err != nil {
		return nil, err
	}
	return subscriptionResponse, nil
}

func (r *Rest) Unsubscribe(subscriptionID string) (*model.SubscriptionDeleteResponse, error) {
	if r.engine == nil {
		return nil, ErrMissingEngine
	}
	response, err := r.engine.Delete(r.getLink("/subscriptions/" + subscriptionID))
	if err != nil {
		return nil, err
	}
	if response.GetCode() != 200 {
		return nil, ErrSubscriptionFailed
	}
	subscriptionDeleteResponse := model.NewSubscriptionDeleteResponse()
	if err := subscriptionDeleteResponse.Unmarshal([]byte(response.GetRaw())); err != nil {
		return nil, err
	}
	return subscriptionDeleteResponse, nil
}

func (r *Rest) CreatePlan(plan *model.Plan) (*model.PlanResponse, error) {
	if r.engine == nil {
		return nil, ErrMissingEngine
	}
	if err := plan.Validate(); err != nil {
		return nil, err
	}
	response, err := r.engine.Post(plan.ToMap(), r.getLink("/plans"))
	if err != nil {
		return nil, err
	}
	if response.GetCode() != 201 {
		return nil, ErrPlanCreationFailed
	}
	planResponse := model.NewPlanResponse()
	if err := planResponse.Unmarshal([]byte(response.GetRaw())); err != nil {
		return nil, err
	}
	return planResponse, nil
}
