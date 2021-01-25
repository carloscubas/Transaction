package account

import (
	"database/sql"
	"fmt"
	"time"
)

type MysqlRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) (*MysqlRepository, error) {
	return &MysqlRepository{
		db: db,
	}, nil
}

func (r *MysqlRepository) InsertAccount(account Account) (*Account, error) {
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

func (r *MysqlRepository) InsertTransactions(transaction Transaction) (*Transaction, error) {
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

func (r *MysqlRepository) GetAccount(id int64) (*Account, error) {

	row := r.db.QueryRow(fmt.Sprintf("select * from Accounts where Account_ID = %d;", id))

	var account = Account{}
	err := row.Scan(&account.Id, &account.DocumentNumber)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *MysqlRepository) GetOperationType(id int64) (*OperationType, error) {

	row := r.db.QueryRow(fmt.Sprintf("select * from OperationsTypes where OperationsType_ID = %d;", id))

	var operationType = OperationType{}
	err := row.Scan(&operationType.Id, &operationType.Description, &operationType.Type)

	if err != nil {
		return nil, err
	}

	return &operationType, nil
}
