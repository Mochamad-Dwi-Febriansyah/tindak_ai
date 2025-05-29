package http

import (
	"errors"
	"net/http"
	"os"
	"tindak_ai/config"
	"tindak_ai/internal/domain"
	"tindak_ai/internal/middleware"
	"tindak_ai/internal/request"
	"tindak_ai/internal/usecase"
	"tindak_ai/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/api/oauth2/v2"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
	logUsecase *usecase.LoggerUsecase
}

func NewAuthHandler(router *gin.RouterGroup, uc *usecase.AuthUsecase, logUc *usecase.LoggerUsecase, jwtSecret string) {
	handler := &AuthHandler{
		usecase: uc,
		logUsecase: logUc,
	}

	publicAuth := router.Group("/auth")
	{
		publicAuth.POST("/register", handler.Register)
		publicAuth.POST("/login", handler.Login) 

		publicAuth.GET("/google/login", handler.LoginByGoogle)
		publicAuth.GET("/google/callback", handler.HandleGoogleCallback)

		publicAuth.POST("/forgot-password", handler.ForgotPassword)
		publicAuth.POST("/reset-password", handler.ResetPassword)

	} 

	protectedAuth := router.Group("/auth")
	protectedAuth.Use(middleware.JWTMiddleware(jwtSecret))
	{
		protectedAuth.POST("/logout", handler.Logout)
		protectedAuth.GET("/me", handler.Me)
		protectedAuth.GET("/permissions", handler.GetUserPermissions)
		protectedAuth.GET("/permissions/:action/:resource", handler.HasPermission)
	} 
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req request.AuthRegisterInput

	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}
	input := domain.AuthRegisterInput{
        FullName:    req.FullName,
        Email:       req.Email,
        Password:    req.Password,
        Gender:      req.Gender,
        NumberPhone: req.NumberPhone, 
    }

	if err := h.usecase.Register(&input); err != nil {
		if err == usecase.ErrEmailExists {
			h.logUsecase.Log(
				domain.MethodTypePost,
				domain.LogLevelWarn,
				"email already exists: "+req.Email,
				"auth.register",
				helper.PtrUUID(uuid.Nil),
				c.ClientIP(),
				c.Request.UserAgent(), 
				helper.PtrString("{}"),
			)
			helper.ValidationFieldErrorResponse(c,"Email", err.Error())
			return
		}else{	
			h.logUsecase.Log(
				domain.MethodTypePost,
				domain.LogLevelError,
				"failed to register user: "+err.Error(),
				"auth.register",
				helper.PtrUUID(uuid.Nil),
				c.ClientIP(),
				c.Request.UserAgent(), 
				helper.PtrString("{}"),
			)
			helper.InternalServerErrorResponse(c, err.Error())
			return
		}
	}
	h.logUsecase.Log(
		domain.MethodTypePost,
		domain.LogLevelInfo,
		"user registered successfully: "+req.Email,
		"auth.register",
		helper.PtrUUID(uuid.Nil),
		c.ClientIP(),
		c.Request.UserAgent(), 
		helper.PtrString("{}"),
	)
	helper.SuccessResponse(c, "user registered successfully", nil)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request.AuthLoginInput
	if err := c.ShouldBind(&req); err != nil {  
		helper.ValidationErrorResponse(c, err)
		return
	}

	input := domain.AuthLoginInput{
		Email: req.Email,
		Password: req.Password,
	}

	token, user, err := h.usecase.Login(&input)
	if err != nil {
		if err == usecase.ErrInvalidCredentials {
			helper.BadRequestResponse(c, usecase.ErrInvalidCredentials.Error())
			return
		}
		if  err == usecase.ErrInactive  || err == usecase.ErrVerification {
			h.logUsecase.Log(
				domain.MethodTypePost,
				domain.LogLevelWarn,
				"failed login attempt for email: "+req.Email+" error: "+err.Error(),
				"auth.login",
				helper.PtrUUID(uuid.Nil),
				c.ClientIP(),
				c.Request.UserAgent(), 
				helper.PtrString("{}"),
			)
			helper.ValidationFieldErrorResponse(c, "Email", err.Error())
			return
		} else {

			h.logUsecase.Log(
				domain.MethodTypePost,
				domain.LogLevelError,
				"error during login for email: "+req.Email+" error: "+err.Error(),
				"auth.login",
				helper.PtrUUID(uuid.Nil),
				c.ClientIP(),
				c.Request.UserAgent(), 
				helper.PtrString("{}"),
			)
			helper.InternalServerErrorResponse(c, err.Error())
			return
		}
		
	}
	// response := gin.H{
	// 	"token": token,
	// 	"user": user,
	// }
	c.SetCookie(
		"jwt_token",      // cookie name
		token,            // value
		3600*1,          // maxAge 24 jam (detik)
		"/",              // path
		"",               // domain (kosong = current domain)
		true,             // secure (true kalau pakai https)
		true,             // httpOnly (tidak bisa diakses JS)
	)

	h.logUsecase.Log(
		domain.MethodTypePost,
		domain.LogLevelInfo,
		"user logged in successfully",
		"auth.login",
		helper.PtrUUID(user.ID),
		c.ClientIP(),
		c.Request.UserAgent(), 
		helper.PtrString("{}"),
	)

	helper.SuccessResponse(c, "User logged in successfully", user)
}

func (h *AuthHandler) Logout(c *gin.Context){
	c.SetCookie(
		"jwt_token",      // cookie name
		"",               // value
		-1,              // maxAge -1 untuk menghapus cookie
		"/",              // path
		"",               // domain (kosong = current domain)
		true,             // secure (true kalau pakai https)
		true,             // httpOnly (tidak bisa diakses JS)
	)
	h.logUsecase.Log(
		domain.MethodTypePost,
		domain.LogLevelInfo,
		"user logged out successfully",
		"auth.logout",
		helper.PtrUUID(uuid.Nil),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString("{}"),
	)
	helper.SuccessResponse(c, "User logged out successfully", nil)
}

func (h *AuthHandler) Me(c *gin.Context) {
	userIDRaw, exist := c.Get(domain.ContextKeyUserID)
	// log.Print(userIDRaw)
	if !exist {
		helper.UnauthorizedResponse(c, "User ID not found in context")
		return
	}

	userID, ok := userIDRaw.(string)
	if !ok {
		helper.UnauthorizedResponse(c, "Invalid User ID type")
		return
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		helper.BadRequestResponse(c, "Invalid User ID format")
		return
	}

	user, err := h.usecase.GetProfile(uid)
	metaStr := `{
		"user_id":"` + uid.String() + `",
		"email":"` + user.Email + `"
	}`
	if err != nil {
		h.logUsecase.Log(
		domain.MethodTypeGet,
		domain.LogLevelError,
		"error retrieving user profile: "+err.Error(),
		"auth.me",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(), 
		helper.PtrString(metaStr),
	)
		helper.InternalServerErrorResponse(c, err.Error())
		return
	}
	h.logUsecase.Log(
		domain.MethodTypeGet,
		domain.LogLevelInfo, 
		"user profile retrieved successfully",
		"auth.me",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(), 
		helper.PtrString(metaStr),
	)
	helper.SuccessResponse(c, "User profile retrieved successfully", user)
}

func (h *AuthHandler) HasPermission(c *gin.Context) {
	userIDRaw, exist := c.Get(domain.ContextKeyUserID)
	if !exist {
		helper.UnauthorizedResponse(c, "User ID not found in context")
		return
	}

	userID, ok := userIDRaw.(string)
	if !ok {
		helper.UnauthorizedResponse(c, "Invalid User ID type")
		return
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		helper.BadRequestResponse(c, "Invalid User ID format")
		return
	}

	action := c.Param("action")
	resource := c.Param("resource")

	hasPermission, err := h.usecase.HasPermission(uid, action, resource)
	metaStr := `{
		"user_id":"` + uid.String() + `",
		"action":"` + action + `",
		"resource":"` + resource + `"
	}`
	if err != nil {
		h.logUsecase.Log(
		domain.MethodTypeGet,
		domain.LogLevelError,
		"error checking user permission: "+err.Error(),
		"auth.has_permission",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(), 
		helper.PtrString(metaStr),
	)
		helper.InternalServerErrorResponse(c, err.Error())
		return
	}

	if hasPermission { 
		c.JSON(200, gin.H{"message": true})
	} else { 
		c.JSON(403, gin.H{"message": false})
	}
}

func (h *AuthHandler) GetUserPermissions(c *gin.Context) {
	userIDRaw, exist := c.Get(domain.ContextKeyUserID)
	if !exist {
		helper.UnauthorizedResponse(c, "User ID not found in context")
		return
	}

	userID, ok := userIDRaw.(string)
	if !ok {
		helper.UnauthorizedResponse(c, "Invalid User ID type")
		return
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		helper.BadRequestResponse(c, "Invalid User ID format")
		return
	}

	permissions, err := h.usecase.GetUserPermissions(uid)
	metaStr := `{
		"user_id":"` + uid.String() + `"
	}`
	if err != nil {
			h.logUsecase.Log(
		domain.MethodTypeGet,
		domain.LogLevelError,
		"error retrieving user permissions: "+err.Error(),
		"auth.get_user_permissions",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(), 
		helper.PtrString(metaStr),
	)
		helper.InternalServerErrorResponse(c, err.Error())
		return
	}

	appName := os.Getenv("APP_ENV")
	parsedRespose := make([]string, len(permissions))
	for i := range permissions {
		parsedRespose[i] = appName + "."+ permissions[i].Resource + "." + permissions[i].Action
	} 

	c.JSON(200, gin.H{"permissions": parsedRespose})
}

func (h *AuthHandler) LoginByGoogle(c *gin.Context) {
	url := config.GoogleOAuthConfig.AuthCodeURL("state_string") 
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *AuthHandler) HandleGoogleCallback(c *gin.Context) {
	ctx := c.Request.Context()
 
	code := c.Query("code")
	if code == "" {
		helper.BadRequestResponse(c, "code parameter is missing")
		return
	}

	token, err := config.GoogleOAuthConfig.Exchange(ctx, code)
	if err != nil {
		h.logUsecase.Log(
				domain.MethodTypeGet,
				domain.LogLevelError,
				"OAuth2 code exchange failed: "+err.Error(),
				"auth-google.login",
				helper.PtrUUID(uuid.Nil),
				c.ClientIP(),
				c.Request.UserAgent(), 
				helper.PtrString("{}"),
			)
		helper.InternalServerErrorResponse(c, "code exchange failed") 
		return
	}

	client := config.GoogleOAuthConfig.Client(ctx, token)
	oauth2Service, err := oauth2.New(client)
	if err != nil { 
			h.logUsecase.Log(
				domain.MethodTypeGet,
				domain.LogLevelError,
				"Failed to create OAuth2 client: "+err.Error(),
				"auth-google.login",
				helper.PtrUUID(uuid.Nil),
				c.ClientIP(),
				c.Request.UserAgent(), 
				helper.PtrString("{}"),
			)
		helper.InternalServerErrorResponse(c, "internal error") 
		return
	} 
	userInfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
			h.logUsecase.Log(
				domain.MethodTypeGet,
				domain.LogLevelError,
				"Failed to fetch user info from Google API: "+err.Error(),
				"auth-google.login",
				helper.PtrUUID(uuid.Nil),
				c.ClientIP(),
				c.Request.UserAgent(), 
				helper.PtrString("{}"),
			)
		helper.InternalServerErrorResponse(c, "failed to fetch user info") 
		return
	}

	jwtToken, err := h.usecase.LoginWithGoogle(ctx, userInfo)
	if err != nil {
		helper.UnauthorizedResponse(c, "login failed") 
		return
	}

		c.SetCookie(
		"jwt_token",      // cookie name
		jwtToken,            // value
		3600*1,          // maxAge 24 jam (detik)
		"/",              // path
		"",               // domain (kosong = current domain)
		true,             // secure (true kalau pakai https)
		true,             // httpOnly (tidak bisa diakses JS)
	)

	h.logUsecase.Log(
		domain.MethodTypePost,
		domain.LogLevelInfo,
		"user logged in successfully",
		"auth-google.login",
		nil,
		c.ClientIP(),
		c.Request.UserAgent(), 
		helper.PtrString("{}"),
	)
 
	// responseData := gin.H{
	// 	"token": jwtToken,
	// 	"user":  userInfo,
	// }

	// helper.SuccessResponse(c, "User logged in successfully", responseData)
	c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_URL")+"/auth/google/callback")

}


func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var req request.ForgotPasswordInput
	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}

	err := h.usecase.SendResetPasswordEmail(req.Email)
	if err != nil {
		if errors.Is(err, domain.ErrResetTokenNotFound){
			helper.NotFoundResponse(c, domain.ErrResetTokenNotFound.Error())
		}
		if errors.Is(err, domain.ErrResetTokenExpired){
			helper.BadRequestResponse(c, domain.ErrResetTokenExpired.Error())
		}
		if errors.Is(err, domain.ErrUserNotFound){
			helper.BadRequestResponse(c, domain.ErrUserNotFound.Error())
		}
		h.logUsecase.Log(
			domain.MethodTypePost,
			domain.LogLevelError,
			"failed to send reset password email: "+err.Error(),
			"auth.forgot_password",
			helper.PtrUUID(uuid.Nil),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString("{}"),
		)
		helper.InternalServerErrorResponse(c, err.Error())
		return
	}

	h.logUsecase.Log(
		domain.MethodTypePost,
		domain.LogLevelInfo,
		"reset password email sent to "+req.Email,
		"auth.forgot_password",
		helper.PtrUUID(uuid.Nil),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString("{}"),
	)
	helper.SuccessResponse(c, "Reset password link sent to email", nil)
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req request.ResetPasswordInput
	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}

	err := h.usecase.ResetPassword(req.Token, req.NewPassword)
	if err != nil {
		h.logUsecase.Log(
			domain.MethodTypePost,
			domain.LogLevelError,
			"failed to reset password: "+err.Error(),
			"auth.reset_password",
			helper.PtrUUID(uuid.Nil),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString("{}"),
		)
		helper.BadRequestResponse(c, err.Error())
		return
	}

	h.logUsecase.Log(
		domain.MethodTypePost,
		domain.LogLevelInfo,
		"password reset successfully",
		"auth.reset_password",
		helper.PtrUUID(uuid.Nil),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString("{}"),
	)
	helper.SuccessResponse(c, "Password has been reset", nil)
}