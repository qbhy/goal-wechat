package wechat

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/openplatform"
	"github.com/silenceper/wechat/v2/pay"
	"github.com/silenceper/wechat/v2/work"
)

type Factory interface {
	Wechat() *wechat.Wechat
	OpenPlatform(name ...string) *openplatform.OpenPlatform
	OfficialAccount(name ...string) *officialaccount.OfficialAccount
	Work(name ...string) *work.Work
	MiniProgram(name ...string) *miniprogram.MiniProgram
	Payment(name ...string) *pay.Pay
}
