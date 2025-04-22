package models

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required,email"`
	Phone    int    `json:"phone"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required" `
	NewPassword string `json:"new_password" binding:"required" `
}

type CreateVerifyEmailRequest struct {
	Token uint `json:"token" binding:"required"`
}

type OptionsRequest struct {
	OptionName  string   `json:"option_name" binding:"required"`
	OptionValue []string `json:"option_values" binding:"required"`
}

type ExternalAuthRequest struct {
	OauthID         string `json:"oauth_id" binding:"required"`
	OauthProvider   string `json:"oauth_provider" binding:"required"`
	IsAgreeTerms    bool   `json:"is_agree_terms"`
	IsSexualDisease bool   `json:"is_sexual_disease"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}
type ValidateForgotPasswordRequest struct {
	Token       uint   `json:"token" binding:"required"`
	NewPassword string `json:"new_password"`
}
