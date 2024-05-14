package wechat

import (
	"github.com/goal-web/supports/logs"
	"github.com/goal-web/supports/utils"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/openplatform"
	"github.com/silenceper/wechat/v2/pay"
	"github.com/silenceper/wechat/v2/work"
	"sync"
)

type factory struct {
	config *Config

	wechat *wechat.Wechat

	openPlatforms    sync.Map
	officialAccounts sync.Map
	works            sync.Map
	miniPrograms     sync.Map
	payments         sync.Map
}

func (factory *factory) Wechat() *wechat.Wechat {
	return factory.wechat
}

func NewFactory(config *Config) Factory {
	var wx = wechat.NewWechat()
	wx.SetCache(config.Cache)
	return &factory{
		config:           config,
		wechat:           wx,
		openPlatforms:    sync.Map{},
		officialAccounts: sync.Map{},
		works:            sync.Map{},
		miniPrograms:     sync.Map{},
		payments:         sync.Map{},
	}
}

func (factory *factory) OpenPlatform(name ...string) *openplatform.OpenPlatform {
	var (
		key = utils.DefaultString(name, factory.config.OpenPlatforms.Default)

		instance, exists = factory.openPlatforms.Load(key)
	)
	if exists {
		return instance.(*openplatform.OpenPlatform)
	}

	var config = factory.config.OpenPlatforms.Apps[key]
	if config == nil {
		logs.WithField("name", key).Warn("wechat.factory.OpenPlatform: app is not defined")
		return nil
	}

	var app = factory.wechat.GetOpenPlatform(config)
	factory.openPlatforms.Store(key, app)

	return app
}

func (factory *factory) OfficialAccount(name ...string) *officialaccount.OfficialAccount {
	var (
		key = utils.DefaultString(name, factory.config.OfficialAccounts.Default)

		instance, exists = factory.officialAccounts.Load(key)
	)
	if exists {
		return instance.(*officialaccount.OfficialAccount)
	}

	var config = factory.config.OfficialAccounts.Apps[key]
	if config == nil {
		logs.WithField("name", key).Warn("wechat.factory.OfficialAccount: app is not defined")
		return nil
	}

	var app = factory.wechat.GetOfficialAccount(config)
	factory.officialAccounts.Store(key, app)

	return app
}

func (factory *factory) Work(name ...string) *work.Work {
	var (
		key = utils.DefaultString(name, factory.config.Works.Default)

		instance, exists = factory.works.Load(key)
	)
	if exists {
		return instance.(*work.Work)
	}

	var config = factory.config.Works.Apps[key]
	if config == nil {
		logs.WithField("name", key).Warn("wechat.factory.Work: app is not defined")
		return nil
	}

	var app = factory.wechat.GetWork(config)
	factory.works.Store(key, app)

	return app
}

func (factory *factory) MiniProgram(name ...string) *miniprogram.MiniProgram {
	var (
		key = utils.DefaultString(name, factory.config.MiniPrograms.Default)

		instance, exists = factory.miniPrograms.Load(key)
	)
	if exists {
		return instance.(*miniprogram.MiniProgram)
	}

	var config = factory.config.MiniPrograms.Apps[key]
	if config == nil {
		logs.WithField("name", key).Warn("wechat.factory.MiniProgram: app is not defined")
		return nil
	}

	var app = factory.wechat.GetMiniProgram(config)
	factory.miniPrograms.Store(key, app)

	return app
}

func (factory *factory) Payment(name ...string) *pay.Pay {
	var (
		key = utils.DefaultString(name, factory.config.Payments.Default)

		instance, exists = factory.payments.Load(key)
	)
	if exists {
		return instance.(*pay.Pay)
	}

	var config = factory.config.Payments.Apps[key]
	if config == nil {
		logs.WithField("name", key).Warn("wechat.factory.Payment: app is not defined")
		return nil
	}

	var app = factory.wechat.GetPay(config)
	factory.payments.Store(key, app)

	return app
}
