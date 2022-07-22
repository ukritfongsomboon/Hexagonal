package models

import "time"

type UserOauthModel struct {
	Provider string `json:"provider" bson:"provider" db:"provider"`
	Id       string `json:"id" bson:"id" db:"id"`
	Email    string `json:"email" bson:"email" db:"email"`
	Password string `json:"password" bson:"password"`
}

type UserModel struct {
	UserID      string           `json:"user_id" bson:"user_id" db:"user_id"`
	Email       string           `json:"email" bson:"email" db:"email"`
	Password    string           `json:"password" bson:"password"`
	Name        string           `json:"name" bson:"name" db:"name"`
	CreatedDate time.Time        `json:"create_date" bson:"create_date" db:"create_date"`
	LastUpdate  time.Time        `json:"update_date" bson:"update_date" db:"update_date"`
	Role        int              `json:"role" bson:"role" db:"role"`
	Status      bool             `json:"status" bson:"status" db:"status"`
	Oauth       []UserOauthModel `json:"oauth" bson:"oauth" db:"oauth"`
}

type UserPaginationModel struct {
	Page  int `json:"page" bson:"page"`
	Row   int `json:"row" bson:"row"`
	Total int `json:"total" bson:"total"`
}

type UserCreateModel struct {
	Email    string `json:"email" bson:"email" db:"email"`
	Password string `json:"password" bson:"password"`
	Name     string `json:"name" bson:"name" db:"name"`
	Role     int    `json:"role" bson:"role" db:"role"`
	Status   bool   `json:"status" bson:"status" db:"status"`
}

type UserResModel struct {
	UserID string `json:"user_id" bson:"user_id" db:"user_id"`
	Email  string `json:"email" bson:"email" db:"email"`
	Name   string `json:"name" bson:"name" db:"name"`
	Role   int    `json:"role" bson:"role" db:"role"`
}

type UserResGetAllModel struct {
	Items      []UserResModel      `json:"item" bson:"item"`
	Pagination UserPaginationModel `json:"pagination" bson:"pagination"`
}

type SignInReqModel struct {
	Username string `json:"username" bson:"username" db:"username"`
	Password string `json:"password" bson:"password" db:"password"`
}

type SignInResModel struct {
	Accesstoken string `json:"accesstoken" bson:"accesstoken" db:"accesstoken"`
	Email       string `json:"email" bson:"email" db:"email"`
	Name        string `json:"name" bson:"name" db:"name"`
	Status      bool   `json:"status" bson:"status" db:"status"`
	Role        int    `json:"role" bson:"role" db:"role"`
}

type SignUpModel struct {
	Email    string `json:"email" bson:"email" db:"email"`
	Password string `json:"password" bson:"password" db:"password"`
	Name     string `json:"name" bson:"name" db:"name"`
	Status   bool   `json:"status" bson:"status" db:"status"`
	Role     int    `json:"role" bson:"role" db:"role"`
}

type SignUpReqModel struct {
	Email    string `json:"email" bson:"email" db:"email"`
	Password string `json:"password" bson:"password"`
	Name     string `json:"name" bson:"name" db:"name"`
}

type SignUpResModel struct {
	Email  string `json:"email" bson:"email" db:"email"`
	Name   string `json:"name" bson:"name" db:"name"`
	Status bool   `json:"status" bson:"status" db:"status"`
	Role   int    `json:"role" bson:"role" db:"role"`
}

type UserUpdateReqModel struct {
	UserID string `json:"user_id" bson:"user_id" db:"user_id"`
	Email  string `json:"email" bson:"email" db:"email"`
	Name   string `json:"name" bson:"name" db:"name"`
}
