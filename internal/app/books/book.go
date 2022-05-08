package books

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func (b *Book) Bind(r *http.Request) error {
	bytes := []byte{}
	r.Body.Read(bytes)
	json.Unmarshal(bytes, b)
	return nil
}
