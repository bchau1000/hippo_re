package model

type UserToCreate struct {
	User     User
	Password string `json:"password"`
}
