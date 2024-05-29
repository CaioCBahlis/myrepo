package mypackage

import (
	"database/sql"
	"fmt"
	"time"
	"os"
	_ "github.com/lib/pq"
)

type product struct {
	Title   string `json:"title"`
	Price   string `json:"price"`
	Reviews string `json:"reviews"`
	Imgurl  string `json:"imgurl"`
	Purl    string `json:"purl"`
	Lupdate string `json:"lupdate"`
	Seller  string `json:"seller "`
}

func OpenConn() *sql.DB{

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))


	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	return db

}

func DBinsert(MyProduct product) {

	db := OpenConn()
	defer db.Close()

	
	InsertSql := `INSERT INTO product (title, price, reviews, imgurl, purl, lupdate, seller) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	
	_, err := db.Exec(InsertSql, MyProduct.Title, MyProduct.Price, MyProduct.Reviews, MyProduct.Imgurl, MyProduct.Purl, time.Now(), MyProduct.Seller)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db, InsertSql)
	
	
}


func DB_Search_and_Update(MyProduct product)bool{
	db := OpenConn()
	defer db.Close()

	var DBSearch string

	QueryInput := "SELECT price FROM product WHERE imgurl=$1"
	err := db.QueryRow(QueryInput, MyProduct.Imgurl).Scan(&DBSearch)

	if err != nil{
		if err.Error() == "sql: no rows in result set"{
			return false
		}	
	}
	
	return true
	
}
