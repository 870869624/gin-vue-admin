
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type InvitationRecordSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    User_id  *int `json:"user_id" form:"user_id" `
    Username  *string `json:"username" form:"username" `
    NewUserId  *int `json:"newUserId" form:"newUserId" `
    NewUsername  *string `json:"newUsername" form:"newUsername" `
    request.PageInfo
}
