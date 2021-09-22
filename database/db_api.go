package database

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/nawazish-github/opinion_server/io"
	"github.com/nawazish-github/opinion_server/model"
	"net/http"
)

func SaveOpinion(opinion model.Opinion, ipAddr string) error {
	db := connect()
	defer db.Close()
	insertOpinion := `insert into opinion (qid, oid, ip_addr, date_time) values ($1, $2, $3, $4)`
	res, err := db.Exec(insertOpinion, opinion.QID, opinion.OptionID, ipAddr, opinion.DateTime)
	if err != nil {
		fmt.Println("Failed to save opinion in DB")
		panic(err)
	}
	fmt.Println("successfully saved opinion in DB ", res)
	return nil
}

func GetQuestion(c *gin.Context, date string) (model.QuestionAndOptions, error) {
	db := connect()
	defer db.Close()

	getQuestion := `select q.id, q.question_prompt, o.id, o.option_prompt from question q inner join option o on  q.id = o.qid where q.date = $1`

	res, err := db.Query(getQuestion, date)

	if err != nil {
		fmt.Println("Failed to fetch question for the day [which day] from DB: ", err)
		io.ErrResponse(c, http.StatusInternalServerError, err)
		return model.QuestionAndOptions{}, err
	}

	r, err := parseRows(res)
	fmt.Println("successfully fetched question: ", r)
	return r, nil
}

func parseRows(res *sql.Rows) (model.QuestionAndOptions, error) {
	m := make(map[string]model.QuestionAndOptions)
	defer res.Close()
	var qid, qPrompt, oid, oPrompt string
	for ; res.Next(); {

		err := res.Scan(&qid, &qPrompt, &oid, &oPrompt)
		if err != nil {
			fmt.Errorf("Error while fetching question for the day: %v", err)
			return model.QuestionAndOptions{}, err
		}

		v, exist := m[qid]

		//smelly code begins from here
		if exist {
			option := model.Option{
				OID:          oid,
				OptionPrompt: oPrompt,
			}
			v.Options = append(v.Options, option)
			m[qid] = v
		} else {
			var resp model.QuestionAndOptions
			option := model.Option{
				OID:          oid,
				OptionPrompt: oPrompt,
			}
			resp.Question.QID = qid
			resp.Question.QuestionPrompt = qPrompt
			resp.Options = append(resp.Options, option)
			m[qid] = resp
		}
		//smelly code ends here
	}
	return m[qid], nil
}

func connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to DB!")
	return db
}
