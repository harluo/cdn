package config

type Domain struct {
	// 主机
	Host string `json:"host" validate:"required,hostname|hostname_port"`
	// 模式
	Scheme string `default:"https" json:"scheme" validate:"oneof=http https"`
	// 匹配
	Pattern string `json:"pattern"`
	// 匹配列表
	Patterns []string `json:"patterns"`
	// 创世云
	Chuangcache *Chuangcache `json:"chuangcache" validate:"required_without_all=Tencent"`
	// 腾讯云
	Tencent *Tencent `json:"tencent" validate:"required_without_all=Chuangcache"`
}
