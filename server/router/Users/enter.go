package Users

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ UsersRouter }

var usersApi = api.ApiGroupApp.UsersApiGroup.UsersApi
