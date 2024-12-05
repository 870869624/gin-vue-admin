package PlatformV1

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ PlatformApi }

var platformService = service.ServiceGroupApp.PlatformV1ServiceGroup.PlatformService
