package Airdrop

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ AirdropRouter }

var airdropApi = api.ApiGroupApp.AirdropApiGroup.AirdropApi
