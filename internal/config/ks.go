package config

type Ks struct {
	A string `json:"a" validate:"required_without_all=B"`
	B string `json:"b" validate:"required_without_all=A"`
}
