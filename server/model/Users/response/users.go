package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/Users"

type UsersResponse struct {
	User Users.Users `json:"user"`
}

type LoginResponse struct {
	User      Users.Users `json:"user"`
	Token     string      `json:"token"`
	ExpiresAt int64       `json:"expiresAt"`
}
