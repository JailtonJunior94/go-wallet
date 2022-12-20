package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("Jailton Junior", "j@j")
	account, err := NewAccount(client)

	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWhenArgsAreInvalid(t *testing.T) {
	account, err := NewAccount(nil)

	assert.NotNil(t, err)
	assert.Nil(t, account)
	assert.Equal(t, ErrClientIsRequired, err)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("Jailton Junior", "j@j")
	account, _ := NewAccount(client)

	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("Jailton Junior", "j@j")
	account, _ := NewAccount(client)

	account.Credit(100)
	account.Debit(50)
	assert.Equal(t, float64(50), account.Balance)
}
