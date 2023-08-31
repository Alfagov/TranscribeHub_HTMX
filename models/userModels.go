package models

type LoginUser struct {
	Email    string `binding:"required" json:"email" form:"email"`
	Password string `binding:"required" json:"password" form:"password"`
}

type RegisterUser struct {
	Id       string `json:"id"`
	Username string `binding:"required" json:"username" form:"username"`
	Email    string `binding:"required" json:"email" form:"email"`
	Password string `binding:"required" json:"password" form:"password"`
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
