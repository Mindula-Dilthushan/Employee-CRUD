package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


var employees [] employee

type employee struct {
	Nic string `json:"nic"`
	FullName string `json:"fullname"`
	Address string `json:"address"`
	Mobile int `json:"mobile"`
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

//Save Section------------------------------------------------------------------
func addEmployee(w http.ResponseWriter, r *http.Request)  {

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	nic := r.FormValue("nic")
	fullname :=r.FormValue("fullname")
	address :=r.FormValue("address")
	mobile :=r.FormValue("mobile")

	db, _ := sql.Open("mysql", "root:1023@tcp(127.0.0.1:3306)/employeegolang")
	insert, err := db.Query("INSERT INTO employee VALUES (?, ?, ?, ?)", nic,fullname, address,mobile)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

}


//Delete Section------------------------------------------------------------------


func main()  {
	fmt.Println("Welcome To Employee Management System....")
	fmt.Println("Server Starting")
	r := mux.NewRouter()
	fmt.Println("Server Running...")
	r.HandleFunc("/api/employee", addEmployee).Methods("POST","OPTION")
	log.Fatal(http.ListenAndServe(":8001",r))

}
