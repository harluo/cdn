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
	for _, cd := range config.Domains {
		if "" != cd.Pattern {
			cd.Patterns = append(cd.Patterns, cd.Pattern)
		}
		domain := builder.Domain().Host(cd.Host).Scheme(cd.Scheme).Pattern(cd.Patterns...).Ignore(cd.Ignores...)
		if nil != cd.Ks {
			signer := domain.Ks()
			if "" != cd.Ks.A {
				signer.A(cd.Ks.A)
			} else if "" != cd.Ks.B {
				signer.B(cd.Ks.B)
			}
			domain = signer.Build()
		} else if nil != cd.Tencent {
			signer := domain.Tencent()
			if "" != cd.Tencent.A {
				signer.A(cd.Tencent.A)
			} else if "" != cd.Tencent.B {
				signer.B(cd.Tencent.B)
			} else if "" != cd.Tencent.C {
				signer.C(cd.Tencent.C)
			} else if nil != cd.Tencent.D {
				signer.D(cd.Tencent.D.Key, cd.Tencent.D.Signature, cd.Tencent.D.Timestamp)
			}
			domain = signer.Build()
		}
		builder = domain.Build()
	}
	client = builder.Build()

	return
}
