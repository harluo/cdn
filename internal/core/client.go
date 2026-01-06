package core

import (
	"github.com/goexl/cdn"
	"github.com/harluo/cdn/internal/config"
)

type Client = cdn.Client

func newClient(config *config.Config) (client *Client, err error) {
	builder := cdn.New()
	if nil != config.Domain {
		config.Domains = append(config.Domains, config.Domain)
	}
	for _, domain := range config.Domains {
		if "" != domain.Pattern {
			domain.Patterns = append(domain.Patterns, domain.Pattern)
		}
		_domain := builder.Domain().Host(domain.Host).Scheme(domain.Scheme).Pattern(domain.Patterns...)
		if nil != domain.Ks {
			signer := _domain.Ks()
			if "" != domain.Ks.A {
				signer.A(domain.Ks.A)
			} else if "" != domain.Ks.B {
				signer.B(domain.Ks.B)
			}
			_domain = signer.Build()
		} else if nil != domain.Tencent {
			signer := _domain.Tencent()
			if "" != domain.Tencent.A {
				signer.A(domain.Tencent.A)
			} else if "" != domain.Tencent.B {
				signer.B(domain.Tencent.B)
			} else if "" != domain.Tencent.C {
				signer.C(domain.Tencent.C)
			} else if nil != domain.Tencent.D {
				signer.D(domain.Tencent.D.Key, domain.Tencent.D.Signature, domain.Tencent.D.Timestamp)
			}
			_domain = signer.Build()
		}
		builder = _domain.Build()
	}
	client = builder.Build()

	return
}
