package config

type Chuangcache struct {
	C string `json:"c" validate:"required,min=6,max=40"`
}
