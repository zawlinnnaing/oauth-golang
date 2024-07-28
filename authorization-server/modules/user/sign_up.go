package user

type SignUpBody struct {
	Email     string `json:"email" form:"email" binging:"required,email"`
	FirstName string `json:"first_name" form:"first_name" binding:"required"`
	LastName  string `json:"last_name" form:"last_name" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required,password"`
}
