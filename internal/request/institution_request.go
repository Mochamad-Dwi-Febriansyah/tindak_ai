package request

type InstitutionCreateRequest struct {
	Name        string `form:"name" binding:"required"`
	Address     string `form:"address" binding:"required"`
	PostalCode  string `form:"postal_code" binding:"required"`
	ContactPhone string `form:"contact_phone" binding:"required"`
	ContactEmail string `form:"contact_email" binding:"required,email"`
	Website     *string `form:"website" binding:"omitempty,url"` 
	Fax         *string `form:"fax" binding:"omitempty"`
	Latitude    float64 `form:"latitude" binding:"required"`
	Longitude   float64 `form:"longitude" binding:"required"`
}

type InstitutionUpdateRequest struct {
	Name         *string  `form:"name"`
	Address      *string  `form:"address"`
	PostalCode   *string  `form:"postal_code"`
	ContactPhone *string  `form:"contact_phone"`
	ContactEmail *string  `form:"contact_email" binding:"omitempty,email"`
	Website      *string `form:"website" binding:"omitempty,url"` 
	Fax          *string `form:"fax" binding:"omitempty"`
	Latitude     *float64 `form:"latitude"`
	Longitude    *float64 `form:"longitude"`
}