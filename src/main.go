package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Hola desde mi aplicación")

	//Abre base de datos
	db, err := sql.Open("mysql", "mondiz:Hola123.@tcp(192.168.1.94:3306)/mia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Exitosa conexion a MySQL")
}
