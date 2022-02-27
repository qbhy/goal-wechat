# goal-wechat
这是一个 goal 的微信 sdk，基于 [silenceper/wechat](https://github.com/silenceper/wechat) 封装，支持多应用

## 安装 - install
```shell
$ go get github.com/qbhy/goal-wechat
```

## 配置 - configuration
```go
// config/wechat.go
package config

import (
	"github.com/goal-web/contracts"
	wechat "github.com/qbhy/goal-wechat"
	"github.com/silenceper/wechat/v2/cache"
	miniProgramConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	officialConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	openPlatformConfig "github.com/silenceper/wechat/v2/openplatform/config"
	paymentConfig "github.com/silenceper/wechat/v2/pay/config"
	workConfig "github.com/silenceper/wechat/v2/work/config"
)

func init() {
	configs["wechat"] = func(env contracts.Env) interface{} {
		return &wechat.Config{
			// 默认 cache
			Cache: cache.NewMemory(),

			// 公众号配置
			OfficialAccounts: wechat.OfficialAccountsConfig{
				Default: env.StringOption("wechat.oc.default", "default"),
				Apps: map[string]*officialConfig.Config{
					"default": {
						AppID:          env.GetString("wechat.oc.appid"),
						AppSecret:      env.GetString("wechat.oc.secret"),
						Token:          env.GetString("wechat.oc.token"),
						EncodingAESKey: env.GetString("wechat.oc.aes"),
						//Cache:          nil,
					},
				},
			},

			// 支付配置
			Payments: wechat.PaymentsConfig{
				Default: env.StringOption("wechat.pay.default", "default"),
				Apps: map[string]*paymentConfig.Config{
					"default": {
						AppID:     env.GetString("wechat.pay.appid"),
						MchID:     env.GetString("wechat.pay.mch_id"),
						Key:       env.GetString("wechat.pay.key"),
						NotifyURL: env.GetString("wechat.pay.notify_url"),
					},
				},
			},

			// 微信小程序配置
			MiniPrograms: wechat.MiniProgramsConfig{
				Default: env.StringOption("wechat.mini.default", "default"),
				Apps: map[string]*miniProgramConfig.Config{
					"default": {
						AppID:     env.GetString("wechat.mini.appid"),
						AppSecret: env.GetString("wechat.mini.secret"),
						//Cache:     nil,
					},
				},
			},

			// 开放平台配置
			OpenPlatforms: wechat.OpenPlatformsConfig{
				Default: env.StringOption("wechat.op.default", "default"),
				Apps: map[string]*openPlatformConfig.Config{
					"default": {
						AppID:          env.GetString("wechat.op.appid"),
						AppSecret:      env.GetString("wechat.op.secret"),
						Token:          env.GetString("wechat.op.token"),
						EncodingAESKey: "",
						//Cache:          nil, // 不填默认使用默认 cache
					},
				},
			},

			// 企业微信配置
			Works: wechat.WorksConfig{
				Default: env.StringOption("wechat.work.default", "default"),
				Apps: map[string]*workConfig.Config{
					"default": {
						CorpID:         env.GetString("wechat.work.corp_id"),
						CorpSecret:     env.GetString("wechat.work.corp_secret"),
						AgentID:        env.GetString("wechat.work.agent_id"),
						RasPrivateKey:  env.GetString("wechat.work.ras_key"),
						Token:          env.GetString("wechat.work.token"),
						EncodingAESKey: env.GetString("wechat.work.aes_key"),
						//Cache:          nil,
					},
				},
			},
		}
	}
}
```

## 使用 - usage
注册服务
```go
// main.go
import (
    "github.com/goal-web/application"
    wechat "github.com/qbhy/goal-wechat"
)

func main()  {
    var app = application.Singleton()
    
    app.RegisterServices(
        // other service
        wechat.ServiceProvider{},
    )
	// ...
}
```
使用
```go
// app/controllers/account/login.go

import (
	"fmt"
    "github.com/goal-web/contracts"
    wechat "github.com/qbhy/goal-wechat"
    "github.com/silenceper/wechat/v2/miniprogram"
)

// ... 

// 通过注入微信工厂，你可以通过工厂获取定义好的应用
func LoginByWechatMini(wechat wechat.Factory, request contracts.HttpRequest) {
    var res, err = wechat.MiniProgram("default").GetAuth().Code2Session(request.GetString("code"))
    // todo: 完善登录逻辑
    fmt.Println(res, err)
}

// 你也可以直接注入 小程序、公众号、开放平台、企业微信、支付号 的实例，这种情况下会获取默认实例
func InjectInstanceExample(program *miniprogram.MiniProgram) {
	// do something ...
}

// ...
```

## 相关链接  
[goal-web](https://github.com/goal-web/goal)  
[qbhy/goal-wechat](https://github.com/qbhy/goal-wechat)  
[silenceper/wechat](https://github.com/silenceper/wechat)  
[微信官方文档](https://developers.weixin.qq.com/doc/)  
qbhy0715@qq.com
