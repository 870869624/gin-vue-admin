package Information

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ InformationApi }

var informationService = service.ServiceGroupApp.InformationServiceGroup.InformationService
