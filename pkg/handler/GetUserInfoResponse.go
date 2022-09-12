package handler

type GetUserInfoResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	About  string `json:"about"`
	Avatar string `json:"avatar"`
}
