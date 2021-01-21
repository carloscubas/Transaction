package account

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlRepository struct {
	db     *sql.DB
}

func NewMysqlRepository(db *sql.DB) (*MysqlRepository, error) {
	return &MysqlRepository{
		db:     db,
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

	rows, err := r.db.Query("select * from Accounts where Account_ID = ?;", id)
	if err != nil {
		return nil, err
	}

	var account = Account{}

	var isFind = false
	for rows.Next() {
		err = rows.Scan(&account.Id, &account.DocumentNumber)
		isFind = true
		if err != nil {
			return nil, err
		}
	}

	if isFind {
		return &account, nil
	}else{
		return nil, nil
	}

}

func (r *MysqlRepository) GetOperationType(id int64) (*OperationType, error) {

	rows, err := r.db.Query("select * from OperationsTypes where OperationsType_ID = ?;", id)
	if err != nil {
		return nil, err
	}

	var operationType = OperationType{}

	for rows.Next() {
		err = rows.Scan(&operationType.Id, &operationType.Description, &operationType.Type)
		if err != nil {
			return nil, err
		}
	}
	return &operationType, nil
}
