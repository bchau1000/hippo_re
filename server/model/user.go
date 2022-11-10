package model

type User struct {
	Id          int    `db:"id" json:"id"`
	Email       string `db:"email" json:"email"`
	DisplayName string `db:"display_name" json:"displayName"`
}
