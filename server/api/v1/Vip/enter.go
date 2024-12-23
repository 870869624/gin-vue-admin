package Vip

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ VipRecordApi }

var vipRecordService = service.ServiceGroupApp.VipServiceGroup.VipRecordService
