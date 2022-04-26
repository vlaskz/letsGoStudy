package privateClassData

import (
	"encoding/json"
	"fmt"
)

type AccountDetails struct {
	id          string
	accountType string
}

type Account struct {
	details      *AccountDetails
	CustomerName string
}

func (account *Account) setDetails(id string, accountType string) {
	account.details = &AccountDetails{id, accountType}
}

func (account *Account) getId() string {
	return account.details.id
}
func (account *Account) getAccountType() string {
	return account.details.accountType
}

func Example() {
	var account *Account = &Account{CustomerName: "Agent Smith"}
	account.setDetails("3251", "privateSecure")

	jsonAccount, _ := json.Marshal(account)

	fmt.Println("Private Class Hidden", string(jsonAccount))
	fmt.Println("Account ID", account.getId())
	fmt.Println("Account Type ", account.getAccountType())
}
