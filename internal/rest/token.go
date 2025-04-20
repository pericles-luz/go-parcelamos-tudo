package rest

import (
	"time"

	"github.com/pericles-luz/go-parcelamos-tudo/internal/model"
)

type Token struct {
	validity time.Time
	key      string
}

func (t *Token) SetValidity(validity string) error {
	dtValidity, err := time.Parse("2006-01-02 15:04:05", validity)
	if nil == err {
		t.validity = dtValidity
	}
	return err
}

func (t *Token) SetExpiresIn(seconds int) {
	t.validity = time.Now().UTC().Add(time.Duration(seconds) * time.Second)
}

func (t *Token) SetKey(key string) {
	t.key = key
}

func (t *Token) IsValid() bool {
	if len(t.key) == 0 {
		return false
	}
	if t.validity.IsZero() {
		return false
	}
	isValid := time.Now().UTC().Before(t.validity)
	return isValid
}

func (t *Token) GetValidity() string {
	return t.validity.Format("2006-01-02 15:04:05")
}

func (t *Token) GetKey() string {
	return t.key
}

func NewToken(auth *model.AuthenticationReturn) *Token {
	result := &Token{}
	result.SetKey(auth.AccessToken)
	result.SetExpiresIn(auth.ExpiresIn)
	return result
}
