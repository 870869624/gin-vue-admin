package Users

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ UsersApi }

var (
	casbinService = service.ServiceGroupApp.UsersServiceGroup.CasbinService
	jwtService    = service.ServiceGroupApp.SystemServiceGroup.JwtService
	usersService  = service.ServiceGroupApp.UsersServiceGroup.UsersService
)
