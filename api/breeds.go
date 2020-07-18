package api

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:thecat_miau@tcp(172.23.148.237:3306)/CatBreeds?multiStatements=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func SearchCatBreads(c *gin.Context) {
	search := c.Query("name")

	if search == "" {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	responseData, err := searchOnDb(&search)
	if err != nil {
		log.Fatal(err)
		return
	}

	if responseData != "" {
		c.Data(http.StatusOK, "application/json", []byte(responseData))
	} else {
		responseData, err := searchOnCatApi(&search)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			insertBreeds(&search, &responseData)
			c.Data(http.StatusOK, "application/json", responseData)
		}
	}
}

func searchOnCatApi(term *string) ([]byte, error) {
	response, err := http.Get("https://api.thecatapi.com/v1/breeds/search?q=" + *term)

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)

	return responseData, err
}

func searchOnDb(term *string) (string, error) {
	db := dbConn()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	rows, err := db.Query("select JsonData from breeds where search = ?", term)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var jsonExtract string
	if rows.Next() {
		err = rows.Scan(&jsonExtract)
		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}
	}

	defer db.Close()
	return jsonExtract, nil
}

func insertBreeds(search *string, breeds *[]byte) error {
	db := dbConn()

	insert, err := db.Prepare("insert into Breeds(Search, JsonData) values (?, ?)")
	if err != nil {
		return err
	}
	_, err = insert.Exec(search, breeds)
	if err != nil {
		return err
	}

	defer insert.Close()
	defer db.Close()
	return nil
}
