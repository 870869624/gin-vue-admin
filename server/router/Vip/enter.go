package Vip

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ VipRecordRouter }

var vipRecordApi = api.ApiGroupApp.VipApiGroup.VipRecordApi
