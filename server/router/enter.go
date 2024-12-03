package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Information"
	"github.com/flipped-aurora/gin-vue-admin/server/router/PublicChain"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Users"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System      system.RouterGroup
	Example     example.RouterGroup
	PublicChain PublicChain.RouterGroup
	Information Information.RouterGroup
	Users       Users.RouterGroup
}
