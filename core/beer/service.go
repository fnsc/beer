package beer

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Reader interface {
	GetAll() ([]*Beer, error)
	Get(id int64) (*Beer, error)
}

type Writer interface {
	Store(beer *Beer) error
	Update(beer *Beer) error
	Destroy(id int64) error
}

type ServiceInterface interface {
	Reader
	Writer
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (service *Service) GetAll() ([]*Beer, error) {
	result := []*Beer{}
	beers, err := service.DB.Query("SELECT * FROM beers")

	if err != nil {
		return nil, err
	}

	defer beers.Close()

	for beers.Next() {
		beer := Beer{}
		err = beers.Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Style)

		if err != nil {
			return nil, err
		}

		result = append(result, &beer)
	}

	return result, nil
}

func (service *Service) Get(id int64) (*Beer, error) {
	beer := Beer{}

	statement, err := service.DB.Prepare("SELECT * FROM beers WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Style)

	if err != nil {
		return nil, err
	}

	return &beer, nil
}

func (service *Service) Store(beer *Beer) error {
	transaction, err := service.DB.Begin()

	if err != nil {
		return err
	}

	statement, err := transaction.Prepare("INSERT INTO beers(id, name, type, style) VALUES (?,?,?,?)")

	if err != nil {
		return err
	}

	_, err = statement.Exec(beer.ID, beer.Name, beer.Type, beer.Style)

	if err != nil {
		transaction.Rollback()

		return nil
	}

	transaction.Commit()

	return nil
}

func (service *Service) Update(beer *Beer) error {
	if beer.ID == 0 {
		return fmt.Errorf("Invalid beer ID")
	}

	transaction, err := service.DB.Begin()

	if err != nil {
		return err
	}

	statement, err := transaction.Prepare("UPDATE beers SET name=?, type=?m style=? WHERE id=?")

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(beer.Name, beer.Type, beer.Style, beer.ID)

	if err != nil {
		transaction.Rollback()

		return err
	}

	transaction.Commit()

	return nil
}

func (service *Service) Destroy(id int64) error {
	if id == 0 {
		return fmt.Errorf("Invalid beer ID")
	}

	transaction, err := service.DB.Begin()

	if err != nil {
		return err
	}

	_, err = transaction.Exec("DELETE FROM beers WHERE id=?", id)

	if err != nil {
		transaction.Rollback()

		return err
	}

	transaction.Commit()

	return nil
}
