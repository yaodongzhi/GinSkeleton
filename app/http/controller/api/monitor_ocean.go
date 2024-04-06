package api

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/service/monitor_ocean"
	"goskeleton/app/utils/response"
)

type MonitorOcean struct {
}

func (u *MonitorOcean) Store(ctx *gin.Context) {
	if err := monitor_ocean.CreateMonitorOceanFactory().Store(ctx); err != nil {
		response.Fail(ctx, consts.CurdCreatFailCode, consts.CurdCreatFailMsg, gin.H{})
		return
	}

	response.Success(ctx, consts.CurdStatusOkMsg, gin.H{})
}
