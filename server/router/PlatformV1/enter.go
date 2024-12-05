package PlatformV1

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ PlatformRouter }

var platformApi = api.ApiGroupApp.PlatformV1ApiGroup.PlatformApi
