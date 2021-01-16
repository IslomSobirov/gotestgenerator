package db

import (
	"log"
	"testing"
)

func TestConnect(t *testing.T) {
	db := Connect()

	t.Run("Tesing connection to database", func(t *testing.T) {
		_, err := db.Query("SHOW TABLES")

		if err != nil {
			log.Fatal(err)
			t.Fail()
		}

	})

	defer db.Close()

}
