package NavigationBar

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ NavigationBarApi }

var NAService = service.ServiceGroupApp.NavigationBarServiceGroup.NavigationBarService
