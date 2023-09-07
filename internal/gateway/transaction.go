package gateway

import "github.com/jailtonjunior94/go-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
