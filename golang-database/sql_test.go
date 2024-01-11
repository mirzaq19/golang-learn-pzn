package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
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

func TestQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlQuery := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}

	fmt.Println("Success query data from database")
}

func TestQuerySQLComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlQuery := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float32
		var createdAt time.Time
		var birthDate sql.NullTime
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("===============")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth date:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created at:", createdAt)
	}

	fmt.Println("Success query data from database")
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	sqlQuery := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	fmt.Println(sqlQuery)
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login:", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	sqlQuery := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(sqlQuery)
	rows, err := db.QueryContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login:", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSQLCommandSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "mirzaq'; DROP TABLE user; #"
	password := "mirzaq"

	sqlQuery := "INSERT INTO user(username,password) VALUES (?,?)"
	_, err := db.ExecContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data to database")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "mirzaq@gmail.com"
	comment := "Test komen"

	sqlQuery := "INSERT INTO comments(email,comment) VALUES (?,?)"
	result, err := db.ExecContext(ctx, sqlQuery, username, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data to database with id:", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := "INSERT INTO comments(email,comment) VALUES (?,?)"
	statement, err := db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke-" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transaction
	sqlQuery := "INSERT INTO comments(email,comment) VALUES (?,?)"
	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke-" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, sqlQuery, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id", id)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
