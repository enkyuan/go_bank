package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type storage interface {
	createAccount(*account)  error
	deleteAccount(int) error
	updateAccount(*account) error
	getAccountByID(int) (*account, error)
	getAccounts() ([] *account, error)
}

type postgresStore struct {
	db *sql.DB
}

func newPostgresStore() (*postgresStore, error) {
	connStr := "user=postgres dbname=postgres password=go_bank sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	} 

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &postgresStore{
		db: db,
	}, nil
}

func (s *postgresStore) init() error {
	return s.createAccountTable()
}

func (s *postgresStore) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}


func (s *postgresStore) createAccount(account *account) error {
	query := `
		insert into account (
			first_name,
			last_name,
			number, 
			balance,
			created_at
		)
		values (
			$1, $2, $3, $4, $5
		)
	`

	response, err := s.db.Query(
		query, 
		account.FirstName, 
		account.LastName, 
		account.Number, 
		account.Balance,
		account.CreatedAt,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v/n", response)

	return nil
}

func (s *postgresStore)  updateAccount(*account) error {
	return nil
}

func (s *postgresStore)  deleteAccount(id int) error {
	return nil
}

func (s *postgresStore)  getAccountByID(id int) (*account, error) {
	return nil, nil
}

func (s *postgresStore) getAccounts() ([] *account, error)  {
	rows, err := s.db.Query("select * from account")

	if err != nil {
		return nil, err
	}

	accounts := [] *account {}

	for rows.Next() {
		account := new(account)		
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
		); 
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	
	return accounts, nil
}