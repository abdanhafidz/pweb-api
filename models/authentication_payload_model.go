package models

type AccountData struct {
	UserID       uint
	VerifyStatus string
	ErrVerif     error
}
