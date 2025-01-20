package NavigationProject

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ NavigationProjectRouter }

var NPApi = api.ApiGroupApp.NavigationProjectApiGroup.NavigationProjectApi
