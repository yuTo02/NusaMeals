package model

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name" form:"name" gorm:"not null"`
	Username string `json:"username" form:"username" gorm:"unique;not null"`
	Email    string `json:"email" form:"email" gorm:"unique;not null"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role" gorm:"type:enum('user','admin');default:'user'"`
}

//type Token struct {
//	Token string `gorm:"-"`
//}
