package Users

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ UsersApi }

var usersService = service.ServiceGroupApp.UsersServiceGroup.UsersService
