package request

type AuthRegisterInput struct {
	FullName    string  `form:"full_name" binding:"required,min=3,max=255"`
    Email       string  `form:"email" binding:"required,email"`
    Password    string  `form:"password" binding:"required,min=6"`
    Gender      string  `form:"gender" binding:"required,oneof=male female other"`
    NumberPhone string  `form:"number_phone" binding:"required"`
    Address     string  `form:"address" binding:"required"` 
}