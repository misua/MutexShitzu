package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Account struct {
	accountID     string
	balance       int
	transactions  []string
	totalOutgoing int
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
	a.transactions = append(a.transactions, "deposit"+strconv.Itoa(amount))
}

func (a *Account) Withdraw(amount int) {
	if a.balance >= amount {
		a.balance -= amount
		a.transactions = append(a.transactions, "withdraw"+strconv.Itoa(amount))
		a.totalOutgoing += amount
	}
}

type BankingSystem struct {
	accounts map[string]*Account
}

func NewBankingSystem() *BankingSystem {
	return &BankingSystem{accounts: make(map[string]*Account)}
}

func (b *BankingSystem) TopSpenders(n int) string {
	type accountInfo struct {
		id    string
		spend int
	}
	accountList := []accountInfo{}

	for id, account := range b.accounts {
		accountList = append(accountList, accountInfo{id: id, spend: account.totalOutgoing})
	}

	sort.SliceStable(accountList, func(i, j int) bool {
		if accountList[i].spend == accountList[j].spend {
			return accountList[i].id < accountList[j].id
		}
		return accountList[i].spend > accountList[j].spend
	})

	topN := accountList[:n]
	result := ""
	for _, acc := range topN {
		if result != "" {
			result += ", "
		}
		result += fmt.Sprintf("%s(%d)", acc.id, acc.spend)
	}
	return result
}

func (b *BankingSystem) CreateAccount(accountID string) string {
	if _, ok := b.accounts[accountID]; ok {
		return "false"
	}
	b.accounts[accountID] = &Account{accountID: accountID}
	return "true"
}

func (b *BankingSystem) Deposit(accountID string, amount int) string {
	account, ok := b.accounts[accountID]
	if !ok {
		return ""
	}
	account.Deposit(amount)
	return strconv.Itoa(account.balance)
}

func (b *BankingSystem) Transfer(sourceAccountID, targetAccountID string, amount int) string {
	sourceAccount, ok1 := b.accounts[sourceAccountID]
	targetAccount, ok2 := b.accounts[targetAccountID]

	if !ok1 || !ok2 || sourceAccountID == targetAccountID || sourceAccount.balance < amount {
		return ""
	}

	sourceAccount.Withdraw(amount)
	targetAccount.Deposit(amount)

	return strconv.Itoa(sourceAccount.balance)
}

func Solution(queries [][]interface{}) []string {
	bankingSystem := NewBankingSystem()
	results := []string{}

	for _, query := range queries {
		action := query[0].(string)

		switch action {
		case "CREATE_ACCOUNT":
			accountID := query[1].(string)
			result := bankingSystem.CreateAccount(accountID)
			results = append(results, result)
		case "DEPOSIT":
			accountID := query[1].(string)
			amount := query[2].(int)
			result := bankingSystem.Deposit(accountID, amount)
			results = append(results, result)
		case "TRANSFER":
			sourceAccountID := query[1].(string)
			targetAccountID := query[2].(string)
			amount := query[3].(int)
			result := bankingSystem.Transfer(sourceAccountID, targetAccountID, amount)
			results = append(results, result)
		case "TOP_SPENDERS":
			n := query[1].(int)
			result := bankingSystem.TopSpenders(n)
			results = append(results, result)
		}
	}

	return results
}

func main() {
	queries := [][]interface{}{
		{"CREATE_ACCOUNT", "account1"},
		{"CREATE_ACCOUNT", "account2"},
		{"CREATE_ACCOUNT", "account3"},
		{"DEPOSIT", "account1", 1000},
		{"DEPOSIT", "account2", 1000},
		{"DEPOSIT", "account3", 1000},
		{"TRANSFER", "account2", "account3", 100},
		{"TRANSFER", "account2", "account1", 100},
		{"TRANSFER", "account3", "account1", 100},
		{"TOP_SPENDERS", 3},
	}

	results := Solution(queries)
	fmt.Println(results) // Output: ["true", "true", "true", "1000", "1000", "1000", "900", "800", "1000", "account2(200), account3(100), account1(0)"]
}
