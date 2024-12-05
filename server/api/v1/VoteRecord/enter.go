package VoteRecord

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ VoteRecordApi }

var voteRecordService = service.ServiceGroupApp.VoteRecordServiceGroup.VoteRecordService
