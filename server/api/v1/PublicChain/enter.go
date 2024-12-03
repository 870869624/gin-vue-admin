package PublicChain

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ PublicChainApi }

var PCService = service.ServiceGroupApp.PublicChainServiceGroup.PublicChainService
