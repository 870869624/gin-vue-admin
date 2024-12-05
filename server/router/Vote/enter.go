package Vote

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ VoteRouter }

var voteApi = api.ApiGroupApp.VoteApiGroup.VoteApi
