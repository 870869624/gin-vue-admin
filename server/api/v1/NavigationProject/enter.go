package NavigationProject

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ NavigationProjectApi }

var NPService = service.ServiceGroupApp.NavigationProjectServiceGroup.NavigationProjectService
