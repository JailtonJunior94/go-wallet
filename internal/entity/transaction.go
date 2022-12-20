package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

var (
	ErrAmountMustBeGreaterThanZero = errors.New("amount must be greater than zero")
	ErrInsufficientFunds           = errors.New("insufficient funds")
)

func NewTransaction(accountFrom, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	if err := transaction.Validate(); err != nil {
		return nil, err
	}

	transaction.Transfer()
	return transaction, nil
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return ErrAmountMustBeGreaterThanZero
	}

	if t.AccountFrom.Balance < t.Amount {
		return ErrInsufficientFunds
	}

	return nil
}

func (t *Transaction) Transfer() {
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)
}
