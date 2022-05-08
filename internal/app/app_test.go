package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

func TestApp_PostBook_Returns201Created(t *testing.T) {
	app := Create()
	w := httptest.NewRecorder()
	json := strings.NewReader(`{ "author": "Robin Hobb", "title": "Assasin's Apprentice" }`)
	r, _ := http.NewRequest("POST", "/books", json)
	r.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	app.ServeHTTP(w, r)

	if w.Code != 201 {
		t.Errorf("expected '%d' but received '%d'", 201, w.Code)
	}
}

func TestApp_PostBook_InsertsBook(t *testing.T) {
	db := createDb()
	defer db.Close()
	db.Exec("DELETE FROM book;")

	author := "Robin Hobb"
	title := "Assasin's Apprentice"

	app := Create()
	w := httptest.NewRecorder()
	json := fmt.Sprintf(`{ "author": "%v", "title": "%v" }`, author, title)
	body := strings.NewReader(json)
	r, _ := http.NewRequest(http.MethodPost, "/books", body)
	r.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	app.ServeHTTP(w, r)

	sqlQuery := `SELECT author, title FROM book WHERE author=$1 AND title=$2;`
	row := db.QueryRow(sqlQuery, author, title)
	var authorInDb string
	var titleInDb string
	err := row.Scan(&authorInDb, &titleInDb)

	if err != nil {
		t.Errorf(err.Error())
	}

	if authorInDb != author {
		t.Errorf("expected author to be '%v' but found '%v'", author, authorInDb)
	}

	if titleInDb != title {
		t.Errorf("expected title to be '%v' but found '%v'", title, titleInDb)
	}
}

func createDb() *sql.DB {
	psqlInfo := "host=localhost port=5432 user=postgres password=password dbname=book sslmode=disable"

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	return db
}
