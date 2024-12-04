package InvitationRecord

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ InvitationRecordRouter }

var IRApi = api.ApiGroupApp.InvitationRecordApiGroup.InvitationRecordApi
