package wechat

import (
	"github.com/goal-web/contracts"
)

type ServiceProvider struct {
}

func (s ServiceProvider) Register(app contracts.Application) {
	app.Singleton("wechat", func() {

	})
}

func (s ServiceProvider) Start() error {
	return nil
}

func (s ServiceProvider) Stop() {
}
