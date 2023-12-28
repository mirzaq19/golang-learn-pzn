package golang_database

import (
	"context"
	"fmt"
	"testing"
)

func TestSQLCommand(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO customer(id,name) VALUES ('eko','Eko')")
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data to database")
}
