package account

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlRepository struct {
	config Config
	db     *sql.DB
}

func NewMysqlRepository(config Config) (*MysqlRepository, error) {
	db, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		return nil, err
	}

	return &MysqlRepository{
		config: config,
		db:     db,
	}, nil
}

func (r *MysqlRepository) InsertAccount(account Account) error {
	stmt, err := r.db.Prepare("insert into Accounts(Document_Number) values (?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(account.DocumentNumber)
	if err != nil {
		return err
	}

	_, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (r *MysqlRepository) InsertTransactions(transaction Transaction) error {
	stmt, err := r.db.Prepare("insert into Transactions(Account_ID, OperationType_ID, Amount, EventDate) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(transaction.AccountId, transaction.OperationTypeId, transaction.Amount, time.Now())
	if err != nil {
		return err
	}

	_, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (r *MysqlRepository) GetAccount(id int64) (*Account, error) {

	rows, err := r.db.Query("select * from Accounts where Account_ID = ?;", id)
	if err != nil {
		return nil, err
	}

	var account = Account{}

	for rows.Next() {
		err = rows.Scan(&account.Id, &account.DocumentNumber)
		if err != nil {
			return nil, err
		}
	}
	return &account, nil
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
