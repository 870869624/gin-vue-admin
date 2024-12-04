package Users

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ UsersApi }

var (
	casbinService = service.ServiceGroupApp.UsersServiceGroup.CasbinService
	usersService  = service.ServiceGroupApp.UsersServiceGroup.UsersService
	jwtService    = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
