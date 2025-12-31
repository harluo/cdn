package config

import (
	"github.com/harluo/config"
)

type Config struct {
	// 域名
	Domain *Domain `json:"domain,omitempty" validate:"required_without=Domains"`
	// 域名
	Domains []*Domain `json:"domains,omitempty" validate:"required_without=Domain"`
}

func newConfig(getter config.Getter) (config *Config, err error) {
	config = new(Config)
	err = getter.Get(&struct {
		CDN *Config `json:"cdn,omitempty" validate:"required"`
	}{
		CDN: config,
	})

	return
}
