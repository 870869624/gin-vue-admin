
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type UsersSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    Username  *string `json:"username" form:"username" `
    Email  *string `json:"email" form:"email" `
    PhoneNumber  *string `json:"phoneNumber" form:"phoneNumber" `
    Age  *int `json:"age" form:"age" `
    Gender  *string `json:"gender" form:"gender" `
    request.PageInfo
}
