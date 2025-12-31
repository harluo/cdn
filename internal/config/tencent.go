package config

type Tencent struct {
	A string `json:"a" validate:"required_without_all=B C D"`
	B string `json:"b" validate:"required_without_all=A C D"`
	C string `json:"c" validate:"required_without_all=A B D"`
	D *D     `json:"d" validate:"required_without_all=A B C"`
}
