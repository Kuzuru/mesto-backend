package user

type User struct {
	ID     int    `json:"id"`
	AuthID string `json:"auth_id"`
	Name   string `json:"name"`
	About  string `json:"about"`
	Avatar string `json:"avatar"`
}
