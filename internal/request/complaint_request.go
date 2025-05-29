package request

type ComplaintCreateRequest struct {
	// ComplaintNumber string  `form:"complaint_number" binding:"required"`
	UserID          string  `form:"user_id"`
	InstitutionID   *string  `form:"institution_id"  `
	Title           string  `form:"title" binding:"required"`
	Description     string  `form:"description" binding:"required"` 
	Location        *string `form:"location"`
	Latitude        *float64 `form:"latitude"`
	Longitude       *float64 `form:"longitude"`
	Status          string   `form:"status" binding:"required,oneof=done pending rejected in_progress withdrawn"`
	EstimatedAt     *string  `form:"estimated_at"` // parse ke time.Time nanti di handler
}

type ComplaintUpdateRequest struct {
	Title       *string  `form:"title" binding:"omitempty"`
	Description *string  `form:"description" binding:"omitempty"`
	Status      *string  `form:"status" binding:"omitempty,oneof=done pending rejected in_progress withdrawn"`
	VerifiedBy  *string  `form:"verified_by" binding:"omitempty"`
	StartedAt   *string  `form:"started_at" binding:"omitempty"`
	CompletedAt *string  `form:"completed_at" binding:"omitempty"`
}

type ComplaintRatingCreateRequest struct {
	UserID  string  `form:"user_id" `
	ComplaintID  string  `form:"complaint_id" binding:"required"`
	Rating  int     `form:"rating" binding:"required,min=1,max=5"`
	Comment *string `form:"comment"`
}

type ComplaintRatingUpdateRequest struct {
	Rating  *int    `form:"rating" binding:"omitempty,min=1,max=5"`
	Comment *string `form:"comment" binding:"omitempty"`
}
