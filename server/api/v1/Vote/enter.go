package Vote

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ VoteApi }

var (
	usersService      = service.ServiceGroupApp.UsersServiceGroup.UsersService
	voteService       = service.ServiceGroupApp.VoteServiceGroup.VoteService
	voteRecordService = service.ServiceGroupApp.VoteRecordServiceGroup.VoteRecordService
)
