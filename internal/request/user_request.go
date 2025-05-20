package request

type UserCreateRequest struct {
    FullName    string  `form:"full_name" binding:"required,min=3,max=255"`
    Email       string  `form:"email" binding:"required,email"`
    Password    string  `form:"password" binding:"required,min=6"`
    Gender      string  `form:"gender" binding:"required,oneof=male female other"`
    NumberPhone string  `form:"number_phone" binding:"required"`
    Address     string  `form:"address" binding:"required"`
    Location    *string `form:"location"`
    DeviceInfo  *string `form:"device_info"`
    AuthProvider *string `form:"auth_provider"`
} 

type UserUpdateRequest struct {
    FullName    *string `form:"full_name" binding:"omitempty,min=3,max=255"`
    Email       *string `form:"email" binding:"omitempty,email"`
    Password    *string `form:"password" binding:"omitempty,min=6"`
    Gender      *string `form:"gender" binding:"omitempty,oneof=male female other"`
    NumberPhone *string `form:"number_phone" binding:"omitempty"`
    Address     *string `form:"address" binding:"omitempty"`
    Location    *string `form:"location"`
    DeviceInfo  *string `form:"device_info"`
    AuthProvider *string `form:"auth_provider"`
}
