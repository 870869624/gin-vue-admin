package Information

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ InformationRouter }

var informationApi = api.ApiGroupApp.InformationApiGroup.InformationApi
