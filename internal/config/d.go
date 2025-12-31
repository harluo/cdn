package config

type D struct {
	// 密钥
	Key string `json:"key" validate:"required"`
	// 签名字段
	Signature string `default:"sign" json:"signature"`
	// 时间字段
	Timestamp string `default:"t" json:"timestamp"`
}
