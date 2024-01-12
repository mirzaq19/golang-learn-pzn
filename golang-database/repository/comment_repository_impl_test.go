package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	golangdatabase "golang-database"
	"golang-database/entity"
	"testing"
)

func TestCommentRepository(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "code.mirzaq@gmail.com",
		Comment: "Ini test comment repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, 90)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindByAll(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
