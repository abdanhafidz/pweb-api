package repositories

import (
	"pweb-api.abdanhafidz.com/models"
)

func GetAccountbyEmail(email string) Repository[models.Account, models.Account] {
	repo := Construct[models.Account, models.Account](
		models.Account{Email: email},
	)
	repo.Transactions(
		WhereGivenConstructor[models.Account, models.Account],
		Find[models.Account, models.Account],
	)
	return *repo
}

func GetAllAccount() Repository[models.Account, []models.Account] {
	repo := Construct[models.Account, []models.Account](
		models.Account{},
	)
	repo.Transactions(
		Find[models.Account, []models.Account],
	)
	return *repo
}
func GetAccountById(AccountId uint) Repository[models.Account, models.Account] {
	repo := Construct[models.Account, models.Account](
		models.Account{Id: AccountId},
	)
	repo.Transactions(
		WhereGivenConstructor[models.Account, models.Account],
		Find[models.Account, models.Account],
	)
	return *repo
}

func UpdateAccount(account models.Account) Repository[models.Account, models.Account] {
	repo := Construct[models.Account, models.Account](
		account,
	)
	repo.Transaction.Save(&repo.Constructor)
	repo.Result = repo.Constructor
	return *repo
}

func GetDetailAccountById(AccountId uint) Repository[models.AccountDetails, models.AccountDetails] {
	repo := Construct[models.AccountDetails, models.AccountDetails](
		models.AccountDetails{AccountId: AccountId},
	)

	// fmt.Println("Account ID:", repo.Constructor.AccountId)
	repo.Transactions(
		WhereGivenConstructor[models.AccountDetails, models.AccountDetails],
		Find[models.AccountDetails, models.AccountDetails],
	)
	return *repo
}

func CreateAccount(account models.Account) Repository[models.Account, models.Account] {
	repo := Construct[models.Account, models.Account](
		account,
	)
	Create(repo)
	return *repo
}

func CreateAccountDetails(accountDetails models.AccountDetails) Repository[models.AccountDetails, models.AccountDetails] {
	repo := Construct[models.AccountDetails, models.AccountDetails](
		accountDetails,
	)
	Create(repo)
	return *repo
}

func UpdateAccountDetails(accountDetails models.AccountDetails) Repository[models.AccountDetails, models.AccountDetails] {
	repo := Construct[models.AccountDetails, models.AccountDetails](
		models.AccountDetails{AccountId: accountDetails.AccountId},
	)
	repo.Transaction.Where("account_id = ?", accountDetails.AccountId).First(&repo.Constructor)
	accountDetails.ID = repo.Constructor.ID
	// fmt.Println(repo.Constructor)
	// fmt.Println(accountDetails)
	repo.Transaction.Updates(accountDetails)
	repo.Result = accountDetails
	return *repo
}
