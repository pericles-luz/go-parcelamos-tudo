package rest

import (
	"errors"

	"github.com/pericles-luz/go-parcelamos-tudo/internal/model"
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
	SetToken(token IToken) error
	Post(payload map[string]interface{}, link string) (IResponse, error)
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

func (r *Rest) SetToken(token IToken) error {
	return r.engine.SetToken(token)
}

func (r *Rest) Authenticate(auth *model.Authentication) error {
	if err := auth.Validate(); err != nil {
		return err
	}
	if r.engine == nil {
		return errors.New("missing rest engine")
	}
	response, err := r.engine.Post(auth.ToMap(), r.getLink("/auth/token"))
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
		return nil, errors.New("missing rest engine")
	}
	if err := subscription.Validate(); err != nil {
		return nil, err
	}
	response, err := r.engine.Post(subscription.ToMap(), r.getLink("/subscriptions"))
	if err != nil {
		return nil, err
	}
	if response.GetCode() != 201 {
		return nil, errors.New("subscription failed")
	}
	subscriptionResponse := model.NewSubscriptionResponse()
	if err := subscriptionResponse.Unmarshal([]byte(response.GetRaw())); err != nil {
		return nil, err
	}
	return subscriptionResponse, nil
}

func (r *Rest) Unsubscribe(subscriptionID string) (*model.SubscriptionDeleteResponse, error) {
	if r.engine == nil {
		return nil, errors.New("missing rest engine")
	}
	response, err := r.engine.Delete(r.getLink("/subscriptions/" + subscriptionID))
	if err != nil {
		return nil, err
	}
	if response.GetCode() != 200 {
		return nil, errors.New("unsubscription failed")
	}
	subscriptionDeleteResponse := model.NewSubscriptionDeleteResponse()
	if err := subscriptionDeleteResponse.Unmarshal([]byte(response.GetRaw())); err != nil {
		return nil, err
	}
	return subscriptionDeleteResponse, nil
}

func (r *Rest) CreatePlan(plan *model.Plan) (*model.PlanResponse, error) {
	if r.engine == nil {
		return nil, errors.New("missing rest engine")
	}
	if err := plan.Validate(); err != nil {
		return nil, err
	}
	response, err := r.engine.Post(plan.ToMap(), r.getLink("/plans"))
	if err != nil {
		return nil, err
	}
	if response.GetCode() != 201 {
		return nil, errors.New("plan creation failed")
	}
	planResponse := model.NewPlanResponse()
	if err := planResponse.Unmarshal([]byte(response.GetRaw())); err != nil {
		return nil, err
	}
	return planResponse, nil
}
