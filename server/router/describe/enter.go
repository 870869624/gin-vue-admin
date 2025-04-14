package describe

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ DescribeRouter }

var DApi = api.ApiGroupApp.DescribeApiGroup.DescribeApi
