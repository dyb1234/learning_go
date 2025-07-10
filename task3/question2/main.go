package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

type Book struct {
	ID     uint
	Title  string
	Author string
	Price  float64
}

func main() {
	db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:13306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	defer db.Close()
	// question 2.1.1
	var employees []Employee
	err = db.Select(&employees, "select * from employees where department = ?", "技术部")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(employees)

	//question 2.1.2
	var employee Employee
	err = db.Get(&employee, "select * from employees where salary = (select max(salary) from employees)")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(employee)

	//question 2.2
	var books []Book
	err = db.Select(&books, "select * from books where price > ?", 50)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(books)
}
