package main

import (
	"database/sql"
	"fmt"
)

func foo() error {
	return sql.ErrConnDone
}
func bar() error {
	return sql.ErrNoRows
}
func baz() error {
	return fmt.Errorf("baz error", sql.ErrNoRows)
}
func main() {
	if err := foo(); err != nil {
		println(err)
	}
	if err := bar(); err != nil {
		println(err)
	}
	if err := baz(); err != nil {
		fmt.Println(err)
	}
}
