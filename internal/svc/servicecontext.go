package svc

import (
	"sync"
	"zerok8scron/internal/config"
)

var (
	once   sync.Once
	svcCtx *ServiceContext
)

type ServiceContext struct {
	Config config.Config
}

func InitSvcCtx(c config.Config) {

	once.Do(func() {
		svcCtx = &ServiceContext{
			Config: c,
		}
	})

}

func GetSvcCtx() *ServiceContext {

	if svcCtx == nil {
		panic("svcCtx is nil")
	}

	return svcCtx
}
