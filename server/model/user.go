package model

import "firebase.google.com/go/auth"

type User struct {
	Id          int    `db:"id" json:"id"`
	UID         string `db:"uid" json:"uid"`
	Email       string `db:"email" json:"email"`
	DisplayName string `db:"display_name" json:"displayName"`
}

func UserRecordToUser(userRecord *auth.UserRecord) *User {
	return &User{
		UID:         userRecord.UID,
		Email:       userRecord.Email,
		DisplayName: userRecord.DisplayName,
	}
}
