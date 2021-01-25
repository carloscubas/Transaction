package account

import (
	"database/sql"
	"fmt"
	"time"
)

type RepositoryDB struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) (*RepositoryDB, error) {
	return &RepositoryDB{
		db: db,
	}, nil
}

// InsertAccount insert the new account in database
func (r *RepositoryDB) InsertAccount(account Account) (*Account, error) {
	stmt, err := r.db.Prepare("insert into Accounts(Document_Number) values (?)")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(account.DocumentNumber)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	account.Id = id

	return &account, nil
}

// InsertTransaction insert the new transaction in database
func (r *RepositoryDB) InsertTransaction(transaction Transaction) (*Transaction, error) {
	stmt, err := r.db.Prepare("insert into Transactions(Account_ID, OperationsType_ID, Amount, EventDate) values (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(transaction.AccountId, transaction.OperationTypeId, transaction.Amount, time.Now())
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	transaction.Id = id

	return &transaction, nil
}

// GetAccount brings the account from within the database.
func (r *RepositoryDB) GetAccount(id int64) (*Account, error) {

	row := r.db.QueryRow(fmt.Sprintf("select * from Accounts where Account_ID = %d;", id))

	var account = Account{}
	err := row.Scan(&account.Id, &account.DocumentNumber)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// GetOperationType brings the operation type from within the database.
func (r *RepositoryDB) GetOperationType(id int64) (*OperationType, error) {

	row := r.db.QueryRow(fmt.Sprintf("select * from OperationsTypes where OperationsType_ID = %d;", id))

	var operationType = OperationType{}
	err := row.Scan(&operationType.Id, &operationType.Description, &operationType.Type)

	if err != nil {
		return nil, err
	}

	return &operationType, nil
}

// GetOperationTypes brings all operation type from within the database.
func (r *RepositoryDB) GetOperationTypes() ([]OperationType, error) {

	var operationsTypes []OperationType

	rows, err := r.db.Query("select * from OperationsTypes")
	if err != nil {
		return nil, err
	}

	var operationType OperationType
	for rows.Next() {
		err := rows.Scan(&operationType.Id, &operationType.Description, &operationType.Type)
		if err != nil {
			return nil, err
		}
		operationsTypes = append(operationsTypes, operationType)
	}
	return operationsTypes, nil
}

func (r *RepositoryDB) GetTransactions() ([]TransactionResponse, error) {
	var transactions []TransactionResponse

	rows, err := r.db.Query("select a.Document_Number as DocumentNumber, " +
		"t.Amount as Amount, ot.Description as Description from Transactions t join Accounts a " +
		"ON a.Account_ID = t.Account_ID join OperationsTypes ot  ON t.OperationsType_ID = ot.OperationsType_ID ")
	if err != nil {
		return nil, err
	}

	var transactionResponse TransactionResponse
	for rows.Next() {
		err := rows.Scan(&transactionResponse.DocumentNumber, &transactionResponse.Amount, &transactionResponse.Description)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transactionResponse)
	}
	return transactions, nil
}
