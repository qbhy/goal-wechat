package wechat

import (
	"github.com/goal-web/contracts"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/openplatform"
	"github.com/silenceper/wechat/v2/pay"
	"github.com/silenceper/wechat/v2/work"
)

type ServiceProvider struct {
}

func (s ServiceProvider) Register(app contracts.Application) {
	app.Singleton("wechat", func(config contracts.Config) Factory {
		return NewFactory(config.Get("wechat").(*Config))
	})
	app.Singleton("wechat.oc", func(wechat Factory) *officialaccount.OfficialAccount {
		return wechat.OfficialAccount()
	})
	app.Singleton("wechat.pay", func(wechat Factory) *pay.Pay {
		return wechat.Payment()
	})
	app.Singleton("wechat.op", func(wechat Factory) *openplatform.OpenPlatform {
		return wechat.OpenPlatform()
	})
	app.Singleton("wechat.mini", func(wechat Factory) *miniprogram.MiniProgram {
		return wechat.MiniProgram()
	})
	app.Singleton("wechat.work", func(wechat Factory) *work.Work {
		return wechat.Work()
	})
}

func (s ServiceProvider) Start() error {
	return nil
}

func (s ServiceProvider) Stop() {
}
