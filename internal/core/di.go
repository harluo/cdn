package core

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newClient,
	).Build().Apply()
}
