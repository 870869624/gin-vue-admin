package Presale

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ PresaleRouter }

var presaleApi = api.ApiGroupApp.PresaleApiGroup.PresaleApi
