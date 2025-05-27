package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": message,
		"data":    data,
	})
} 

// CreatedResponse mengirim status 201 dengan message dan data
func CreatedResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": message,
		"data":    data,
	})
}

// CreatedResponseNoContent mengirim status 201 dengan message saja tanpa data
func CreatedResponseNoContent(c *gin.Context, message string) {
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": message,
	})
}

// BadRequestResponse 400 Bad Request
func BadRequestResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"message": message,
	})
}

// NotFoundResponse 404 Not Found
func NotFoundResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  http.StatusNotFound,
		"message": message,
	})
}

// InternalServerErrorResponse 500 Internal Server Error
func InternalServerErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"message": message,
	})
}

func ValidationErrorResponse(c *gin.Context, err error) {
    errors := make(map[string]string)

    if errs, ok := err.(validator.ValidationErrors); ok {
        for _, e := range errs {
            field := e.Field()
            tag := e.Tag()
            var msg string
            switch tag {
            case "required":
                msg = field + " is required"
            case "email":
                msg = field + " must be a valid email"
            // Tambah case lain sesuai kebutuhan validasi
            default:
                msg = field + " is not valid"
            }
            errors[field] = msg
        }
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": "validation error",
            "errors":  errors,
        })
    } else {
        // Kalau error bukan dari validator.ValidationErrors
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": err.Error(),
        })
    }
}

func ValidationFieldErrorResponse(c *gin.Context, field string, message string) {
	errors := map[string]string{
		field: message,
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"message": "validation error",
		"errors":  errors,
	})
}


func UnauthorizedResponse(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  http.StatusUnauthorized,
		"message": message,
	})
}

func ForbiddenResponse(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, gin.H{
		"status":  http.StatusForbidden,
		"message": message,
	})
}