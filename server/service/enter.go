package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/Airdrop"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Information"
	"github.com/flipped-aurora/gin-vue-admin/server/service/InvitationRecord"
	"github.com/flipped-aurora/gin-vue-admin/server/service/PublicChain"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Users"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Vote"
	"github.com/flipped-aurora/gin-vue-admin/server/service/VoteRecord"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup           system.ServiceGroup
	ExampleServiceGroup          example.ServiceGroup
	PublicChainServiceGroup      PublicChain.ServiceGroup
	InformationServiceGroup      Information.ServiceGroup
	UsersServiceGroup            Users.ServiceGroup
	InvitationRecordServiceGroup InvitationRecord.ServiceGroup
	AirdropServiceGroup          Airdrop.ServiceGroup
	VoteServiceGroup             Vote.ServiceGroup
	VoteRecordServiceGroup       VoteRecord.ServiceGroup
}
