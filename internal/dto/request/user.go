package request

type RegisterUser struct {
	Username       string `json:"username" form:"username" validate:"required,min=4,max=12"`
	Email          string `json:"email" form:"email" validate:"required,email"`
	Password       string `json:"password" form:"password" validate:"required,min=6,max=12"`
	RetypePassword string `json:"retype_password" form:"retype_password" validate:"required,min=6,max=12"`
	Role           string
}

type LoginUser struct {
	EmailOrUsername string `json:"email_or_username" form:"email_or_username" validate:"required"`
	Password        string `json:"password" form:"password" validate:"required"`
}

type UpdateUser struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Username    string `json:"username" form:"username" validate:"required,min=4,max=12"`
	Email       string `json:"email" form:"email" validate:"required,email"`
	Password    string `json:"password" form:"password" validate:"required,min=6,max=12"`
	Gender      string `json:"gender" form:"gender" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Picture     string `json:"picture" form:"picture"`
}
