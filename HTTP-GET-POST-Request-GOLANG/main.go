package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
)
	
type Employee struct{
	ID string `json:"id"`
	Name string `json::"name"`
	Salary float64 `json:"salary"`
}
	
var emps []Employee


func getfetch(){

	url := "http://localhost:8080/employees"

	response,err := http.Get(url)

	if err != nil{
		log.Fatal(err)
	}

	defer response.Body.Close()

	responseData,err := ioutil.ReadAll(response.Body)

	if err != nil{
		log.Fatal(err)
	}

	err = json.Unmarshal(responseData,&emps)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(emps)

}

func postEmployee(emp Employee) {
	url := "http://localhost:8080/employees"

	
	jsonData, err := json.Marshal(emp)
	if err != nil {
		log.Fatal(err)
	}

	
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(responseData,&emps)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(emps)

	
}
func main(){
	getfetch()

	newEmp := Employee{
		ID:     "E003",
		Name:   "JB",
		Salary: 60000.00,
	}
	postEmployee(newEmp)
}