package PublicChain

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ PublicChainRouter }

var PCApi = api.ApiGroupApp.PublicChainApiGroup.PublicChainApi
