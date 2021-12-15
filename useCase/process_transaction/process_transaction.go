package process_transaction

import "github.com/BrunoCesarAngst/golang/entity"

type ProcessTransaction struct {
	Repository entity.TransactionRepository
}

func NewProcessTransaction(repository entity.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{repository}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	invalidTransaction := transaction.IsValid()
	if invalidTransaction != nil {
		return p.rejectedTransaction(transaction, invalidTransaction)
	}

	return p.approvedTransaction(transaction)
}

func (p *ProcessTransaction) rejectedTransaction(transaction *entity.Transaction, invalidTransaction error) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "rejected", invalidTransaction.Error())
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       "rejected",
		ErrorMassage: invalidTransaction.Error(),
	}

	return output, nil
}

func (p *ProcessTransaction) approvedTransaction(transaction *entity.Transaction) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "approved", "")
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       "approved",
		ErrorMassage: "",
	}

	return output, nil
}
