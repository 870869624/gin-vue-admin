package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Information"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/PublicChain"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Users"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup      system.ApiGroup
	ExampleApiGroup     example.ApiGroup
	PublicChainApiGroup PublicChain.ApiGroup
	InformationApiGroup Information.ApiGroup
	UsersApiGroup       Users.ApiGroup
}
