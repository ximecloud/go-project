package dto

import "time"

type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id        uint64    `json:"id"`
	Username  string    `json:"username"`
	Status    bool      `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}
