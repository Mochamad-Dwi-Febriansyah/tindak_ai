package request
 

type UserInstitutionCreateRequest struct {
	UserID        string  `form:"user_id" binding:"required"`
	InstitutionID string  `form:"institution_id" binding:"required"`
	Position string  `form:"position" binding:"required"`
}

type UserInstitutionUpdateRequest struct {
	UserID        *string  `form:"user_id" binding:"omitempty"`
	InstitutionID *string  `form:"institution_id" binding:"omitempty"`
	Position *string  `form:"position" binding:"omitempty"`
}