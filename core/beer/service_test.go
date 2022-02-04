package beer_test

import (
	"database/sql"
	"testing"

	"github.com/fnsc/beer/core/beer"
	_ "github.com/mattn/go-sqlite3"
)

func TestStore(test *testing.T) {
	newBeer := beer.Beer{
		ID:    1,
		Name:  "Heineken",
		Type:  beer.Lager,
		Style: beer.Pale,
	}

	db, err := sql.Open("sqlite3", "../../data/beer_test.db")

	if err != nil {
		test.Fatalf("Error while connecting to database %s", err.Error())
	}

	err = refreshDatabase(db)

	if err != nil {
		test.Fatalf("Error while refreshing the database %s", err.Error())
	}

	defer db.Close()

	service := beer.NewService(db)

	err = service.Store(&newBeer)

	if err != nil {
		test.Fatalf("Error while saving in database: %s", err.Error())
	}

	storedBeer, err := service.Get(1)

	if err != nil {
		test.Fatalf("Error while getting from database: %s", err.Error())
	}

	if storedBeer.ID != 1 {
		test.Fatalf("Wrong data. Expecting %d, Got: %d", 1, storedBeer.ID)
	}
}

func refreshDatabase(db *sql.DB) error {
	transaction, err := db.Begin()

	if err != nil {
		return err
	}

	_, err = transaction.Exec("DELETE FROM beers")

	if err != nil {
		return err
	}

	transaction.Commit()

	return nil
}
