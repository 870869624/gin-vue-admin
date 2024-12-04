package InvitationRecord

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ InvitationRecordApi }

var IRService = service.ServiceGroupApp.InvitationRecordServiceGroup.InvitationRecordService
