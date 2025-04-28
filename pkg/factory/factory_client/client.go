package factory_client

import (
	"github.com/pericles-luz/go-parcelamos-tudo/pkg/rest"
)

func NewClient(configPath string, rules []string) (*rest.Rest, error) {
	engine := rest.NewEngine(map[string]interface{}{"InsecureSkipVerify": true})
	restEntity, err := rest.NewRest(engine, configPath, rules)
	if err != nil {
		return nil, err
	}
	return restEntity, nil
}
