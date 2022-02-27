package wechat

import (
	"github.com/silenceper/wechat/v2/cache"
	miniProgramConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	officialConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	openPlatformConfig "github.com/silenceper/wechat/v2/openplatform/config"
	paymentConfig "github.com/silenceper/wechat/v2/pay/config"
	workConfig "github.com/silenceper/wechat/v2/work/config"
)

type OfficialAccountsConfig struct {
	Default string
	Apps    map[string]*officialConfig.Config
}

type PaymentsConfig struct {
	Default string
	Apps    map[string]*paymentConfig.Config
}

type MiniProgramsConfig struct {
	Default string
	Apps    map[string]*miniProgramConfig.Config
}

type OpenPlatformsConfig struct {
	Default string
	Apps    map[string]*openPlatformConfig.Config
}

type WorksConfig struct {
	Default string
	Apps    map[string]*workConfig.Config
}

type Config struct {
	Cache cache.Cache

	OfficialAccounts OfficialAccountsConfig
	Payments         PaymentsConfig
	MiniPrograms     MiniProgramsConfig
	OpenPlatforms    OpenPlatformsConfig
	Works            WorksConfig
}
