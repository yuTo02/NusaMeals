package request

type RegisterUser struct {
	Name           string `json:"name" form:"name" validate:"required"`
	Username       string `json:"username" form:"username" validate:"required,min=4,max=12"`
	Email          string `json:"email" form:"email" validate:"required,email"`
	Password       string `json:"password" form:"password" validate:"required,min=6,max=12"`
	RetypePassword string `json:"retype_password" form:"retype_password" validate:"required,min=6,max=12"`
}

type LoginUser struct {
	Username string `json:"username" form:"username" validate:"required,min=4,max=12"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateUser struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Username string `json:"username" form:"username" validate:"required,min=4,max=12"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6,max=12"`
}
