package monitor_ocean

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/model"
)

func CreateMonitorOceanFactory() *MonitorOcean {
	return &MonitorOcean{model.CreateMonitorOceanFactory("")}
}

type MonitorOcean struct {
	monitorOceanModel *model.MonitorOceanModel
}

func (u *MonitorOcean) Store(ctx *gin.Context) (err error) {
	return u.monitorOceanModel.CreateOneOcean(nil, ctx)
}
