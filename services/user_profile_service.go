package services

import (
	"regexp"
	"strings"

	"pweb-api.abdanhafidz.com/models"
	"pweb-api.abdanhafidz.com/repositories"
)

type UserProfileService struct {
	Service[models.AccountDetails, models.UserProfileResponse]
}

// SanitizePhoneNumber membersihkan dan menormalkan nomor telepon ke format +62
func SanitizePhoneNumber(input string) string {
	// Hilangkan semua spasi dan strip
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "-", "")
	input = strings.ReplaceAll(input, "(", "")
	input = strings.ReplaceAll(input, ")", "")

	// Hilangkan semua karakter non-digit kecuali +
	re := regexp.MustCompile(`[^0-9\+]`)
	input = re.ReplaceAllString(input, "")

	// Handle nomor diawali 0 (contoh: 0812...) menjadi +62812...
	if strings.HasPrefix(input, "0") {
		input = "+62" + input[1:]
	}

	// Handle jika diawali dengan 62 tanpa + (contoh: 62812...)
	if strings.HasPrefix(input, "62") && !strings.HasPrefix(input, "+62") {
		input = "+" + input
	}

	// Handle jika tidak ada awalan +62 sama sekali (contoh: 8123456789)
	if !strings.HasPrefix(input, "+62") {
		if strings.HasPrefix(input, "8") {
			input = "+62" + input
		}
	}

	return input
}
func (s *UserProfileService) Create() {
	userProfile := repositories.CreateAccountDetails(s.Constructor)
	s.Error = userProfile.RowsError
	if userProfile.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "There is no account with given credentials!"
		return
	}
	s.Result = models.UserProfileResponse{
		Account: repositories.GetAccountById(s.Constructor.AccountId).Result,
		Details: userProfile.Result,
	}
}
func (s *UserProfileService) Retrieve() {
	userProfile := repositories.GetDetailAccountById(s.Constructor.AccountId)
	s.Error = userProfile.RowsError
	if userProfile.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "There is no account with given credentials!"
		return
	}
	s.Result = models.UserProfileResponse{
		Account: repositories.GetAccountById(s.Constructor.AccountId).Result,
		Details: userProfile.Result,
	}
	s.Result.Account.Password = "SECRET"
}

func (s *UserProfileService) Update() {
	if s.Constructor.PhoneNumber != nil {
		phoneNumber := *s.Constructor.PhoneNumber
		*s.Constructor.PhoneNumber = SanitizePhoneNumber(phoneNumber)
	}
	userProfile := repositories.UpdateAccountDetails(s.Constructor)
	s.Error = userProfile.RowsError
	if userProfile.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "There is no account with given credentials!"
		return
	}
	account := repositories.GetAccountById(s.Constructor.AccountId)
	account.Result.IsDetailCompleted = (userProfile.Result.InitialName != "" &&
		userProfile.Result.FullName != nil &&
		userProfile.Result.PhoneNumber != nil &&
		userProfile.Result.University != nil)
	repositories.UpdateAccount(account.Result)
	s.Result = models.UserProfileResponse{
		Account: account.Result,
		Details: userProfile.Result,
	}
	s.Result.Account.Password = "SECRET"
}
