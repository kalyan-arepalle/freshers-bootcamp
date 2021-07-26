package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ex2/Config"
	"ex2/Controllers"
	"ex2/Models"
	"ex2/Routes"

	"github.com/jinzhu/gorm"
)


func TestGet(t *testing.T) {
	//Connect to SQL
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	//Setup router, send request
	router := Routes.SetupRouter()
	router.GET("/user-api/user/", Controllers.GetStudents)

	req, _ := http.NewRequest("GET", "/user-api/user/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	//Store the response
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":1,"first-name":"pqr","last-name":"abc","dob":"01-02-2003","address":"pilani","subject":"english","marks":20},{"id":3,"first-name":"pqr","last-name":"abc","dob":"01-02-2003","address":"pilani","subject":"english","marks":20},{"id":6,"first-name":"abd","last-name":"efg","dob":"31-03-2002","address":"delhi","subject":"bio","marks":25}]`
	if resp.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}
}

func TestPost(t *testing.T) {
	//Connect to SQL
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB.AutoMigrate(&Models.Student{})

	//Setup router, send request
	router := Routes.SetupRouter()
	router.POST("/user-api/user/",Controllers.CreateStudent)

	newStudent := Models.Student{
		FirstName: "xyz",
		LastName: "pqr",
		DOB: "30-01-2001",
		Address: "US",
		Subject: "Maths",
		Marks: 15,
	}

	responseBody,_ := json.Marshal(newStudent)
	req, _ := http.NewRequest("POST", "/user-api/user/", bytes.NewBuffer([]byte(responseBody)))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":10,"first-name":"xyz","last-name":"pqr","dob":"30-01-2001","address":"US","subject":"Maths","marks":15}`
	if resp.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}
}

func TestPut(t *testing.T) {
	//Connect to SQL
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB.AutoMigrate(&Models.Student{})

	//Setup router, send request
	router := Routes.SetupRouter()
	router.POST("/user-api/user/3",Controllers.UpdateStudent)

	newStudent := Models.Student{
		FirstName: "fasak",
		LastName: "hasak",
		Address: "bermuda",
		Marks: 1,
	}

	responseBody,_ := json.Marshal(newStudent)
	req, _ := http.NewRequest("POST", "/user-api/user/3", bytes.NewBuffer([]byte(responseBody)))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":3,"first-name":"fasak","last-name":"hasak","dob":"01-02-2003","address":"bermuda","subject":"english","marks":1}`
	if resp.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}
}

