package facade

import (
	"fmt"
)

type Account struct {
	id          string
	accountType string
}

func (account *Account) Create(accountType string) *Account {
	fmt.Println("Account creation with type " + accountType)
	account.accountType = accountType
	return account
}

func (account *Account) GetById(id string) *Account {
	fmt.Println("Got Account by id " + id)
	return account
}

func (account *Account) DeleteById(id string) {
	fmt.Println("Account has been deleted")
}

type Customer struct {
	name string
	id   int
}

func (customer *Customer) Create(name string) *Customer {
	fmt.Println("Creating Customer")
	customer.name = name
	return customer
}

type Transaction struct {
	id            string
	amount        float64
	srcAccountId  string
	destAccountId string
}

func (facade *BranchManagerFacade) CreateCustomerAccount(customerName string, accountType string) (*Customer, *Account) {
	var customer = facade.customer.Create(customerName)
	var account = facade.account.Create(accountType)

	return customer, account
}

func (trs *Transaction) Create(srcAccId string, destAccId string, amount float64) *Transaction {
	fmt.Println("Creating Transaction")
	trs.srcAccountId = srcAccId
	trs.destAccountId = destAccId
	trs.amount = amount

	return trs
}

type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

func (facade *BranchManagerFacade) CreateTransaction(srcAccId string, destAccId string, amount float64) *Transaction {
	var transaction = facade.transaction.Create(srcAccId, destAccId, amount)
	return transaction
}

func Example() {
	var facade = NewBranchManagerFacade()
	var customer *Customer
	var account *Account

	customer, account = facade.CreateCustomerAccount("Isaias Velasquez", "Porquinho da BMW")

	fmt.Println(customer.name)
	fmt.Println(account.accountType)

	var transaction = facade.CreateTransaction("21456", "87345", 1000)

	fmt.Println(transaction.amount)
}
