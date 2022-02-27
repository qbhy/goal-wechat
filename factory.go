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

func (this *factory) Wechat() *wechat.Wechat {
	return this.wechat
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

func (this *factory) OpenPlatform(name ...string) *openplatform.OpenPlatform {
	var (
		key = utils.DefaultString(name, this.config.OpenPlatforms.Default)

		instance, exists = this.openPlatforms.Load(key)
	)
	if exists {
		return instance.(*openplatform.OpenPlatform)
	}

	var config = this.config.OpenPlatforms.Apps[key]
	if config == nil {
		logs.WithField("name", key).Warn("wechat.factory.OpenPlatform: app is not defined")
		return nil
	}

	var app = this.wechat.GetOpenPlatform(config)
	this.openPlatforms.Store(key, app)

	return app
}

func (this *factory) OfficialAccount(name ...string) *officialaccount.OfficialAccount {
	var (
		key = utils.DefaultString(name, this.config.OfficialAccounts.Default)

		instance, exists = this.officialAccounts.Load(key)
	)
	if exists {
		return instance.(*officialaccount.OfficialAccount)
	}

	var config = this.config.OfficialAccounts.Apps[key]
	if config == nil {
		logs.WithField("name", key).Warn("wechat.factory.OfficialAccount: app is not defined")
		return nil
	}

	var app = this.wechat.GetOfficialAccount(config)
	this.officialAccounts.Store(key, app)

	return app
}

func (this *factory) Work(name ...string) *work.Work {
	var (
		key = utils.DefaultString(name, this.config.Works.Default)

		instance, exists = this.works.Load(key)
	)
	if exists {
		return instance.(*work.Work)
	}

	var config = this.config.Works.Apps[key]
	if config == nil {
		logs.WithField("name", key).Warn("wechat.factory.Work: app is not defined")
		return nil
	}

	var app = this.wechat.GetWork(config)
	this.works.Store(key, app)

	return app
}

func (this *factory) MiniProgram(name ...string) *miniprogram.MiniProgram {
	var (
		key = utils.DefaultString(name, this.config.MiniPrograms.Default)

		instance, exists = this.miniPrograms.Load(key)
	)
	if exists {
		return instance.(*miniprogram.MiniProgram)
	}

	var config = this.config.MiniPrograms.Apps[key]
	if config == nil {
		logs.WithField("name", key).Warn("wechat.factory.MiniProgram: app is not defined")
		return nil
	}

	var app = this.wechat.GetMiniProgram(config)
	this.miniPrograms.Store(key, app)

	return app
}

func (this *factory) Payment(name ...string) *pay.Pay {
	var (
		key = utils.DefaultString(name, this.config.Payments.Default)

		instance, exists = this.payments.Load(key)
	)
	if exists {
		return instance.(*pay.Pay)
	}

	var config = this.config.Payments.Apps[key]
	if config == nil {
		logs.WithField("name", key).Warn("wechat.factory.Payment: app is not defined")
		return nil
	}

	var app = this.wechat.GetPay(config)
	this.payments.Store(key, app)

	return app
}
