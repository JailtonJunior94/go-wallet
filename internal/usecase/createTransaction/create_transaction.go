package createtransaction

import (
	"github.com/jailtonjunior94/go-wallet/internal/entity"
	"github.com/jailtonjunior94/go-wallet/internal/gateway"
)

type CreateTransactionInput struct {
	AccountIDFrom string
	AccountIdTo   string
	Amount        float64
}

type CreateTransactionOutput struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(t gateway.TransactionGateway, a gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: t,
		AccountGateway:     a,
	}
}

func (uc *CreateTransactionUseCase) Execute(input *CreateTransactionInput) (*CreateTransactionOutput, error) {
	accountFrom, err := uc.AccountGateway.FindByID(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := uc.AccountGateway.FindByID(input.AccountIdTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = uc.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutput{ID: transaction.ID}, nil
}
