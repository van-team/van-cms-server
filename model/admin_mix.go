package model

type AdminMix struct {
	ID         uint64 `json:"_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       Array  `json:"role"`
	Resource   Array  `json:"resource"`
	Acl        Array  `json:"acl"`
	Permission Array  `json:"permission"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Avatar     string `json:"avatar"`
	Status     *bool  `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}
