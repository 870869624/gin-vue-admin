package VoteRecord

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ VoteRecordRouter }

var voteRecordApi = api.ApiGroupApp.VoteRecordApiGroup.VoteRecordApi
