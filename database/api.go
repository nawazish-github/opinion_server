package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nawazish-github/opinion_server/model"
)

func SaveOpinion(opinion model.Opinion) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to DB!")

	insertOpinion := `insert into opinion (qid, oid, ip_addr) values ($1, $2, $3)`
	res, err := db.Exec(insertOpinion, opinion.QID, opinion.OptionID, opinion.IPAddress)
	if err != nil {
		fmt.Println("Failed to save opinion in DB")
		panic(err)
	}
	fmt.Println("successfully saved opinion in DB ", res)
	return nil
}
