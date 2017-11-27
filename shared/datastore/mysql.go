package datastore

import (
	"database/sql"
	"log"

	"github.com/okeyonyia123/cityrescue/models"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDS struct {
	*sql.DB
}

func NewMySQLDS(dataSourceName string) (*MySQLDS, error) {

	connection, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &MySQLDS{
		DB: connection}, nil
}

func (m *MySQLDS) CreateUser(user *models.User) error {

	begin, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer begin.Rollback()

	stmt, err := begin.Prepare("INSERT INTO user(uuid, username, first_name, last_name, email, password_hash) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.UUID, user.Username, user.FirstName, user.LastName, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	err = begin.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (m *MySQLDS) GetUser(username string) (*models.User, error) {

	stmt, err := m.Prepare("SELECT uuid, username, first_name, last_name, email, password_hash, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM user WHERE username = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(username)
	u := models.User{}
	err = row.Scan(&u.UUID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.PasswordHash, &u.TimestampCreated, &u.TimestampModified)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &u, err
}

func (m *MySQLDS) Close() {
	m.Close()
}
