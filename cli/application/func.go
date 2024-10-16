package application

import (
	"github.com/bdgroup/config"
	"github.com/bdgroup/service"
)

func getService() *service.Service {
	return config.Instance.ActiveService()
}
