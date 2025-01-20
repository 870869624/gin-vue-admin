package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Airdrop"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Information"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/InvitationRecord"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/NavigationBar"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/NavigationProject"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/PlatformV1"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Presale"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/PublicChain"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Users"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Vip"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Vote"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/VoteRecord"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/platform"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup            system.ApiGroup
	ExampleApiGroup           example.ApiGroup
	PublicChainApiGroup       PublicChain.ApiGroup
	InformationApiGroup       Information.ApiGroup
	UsersApiGroup             Users.ApiGroup
	InvitationRecordApiGroup  InvitationRecord.ApiGroup
	AirdropApiGroup           Airdrop.ApiGroup
	VoteApiGroup              Vote.ApiGroup
	VoteRecordApiGroup        VoteRecord.ApiGroup
	PlatformApiGroup          platform.ApiGroup
	PlatformV1ApiGroup        PlatformV1.ApiGroup
	PresaleApiGroup           Presale.ApiGroup
	VipApiGroup               Vip.ApiGroup
	NavigationBarApiGroup     NavigationBar.ApiGroup
	NavigationProjectApiGroup NavigationProject.ApiGroup
}
