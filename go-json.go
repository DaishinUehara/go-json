package main

import (
    "net/http"
	"encoding/json"
)

type BookJson struct {
    ID int `json:"id"`
    Title string `json:"title"`
    Good int `json:"good"`
}

func booksJson(w http.ResponseWriter, r *http.Request) {
	books:= []BookJson{}
	books=append(books,BookJson{0,"本0",5} )
	books=append(books,BookJson{1,"本1",5} )
	books=append(books,BookJson{2,"本2",5} )
	books=append(books,BookJson{3,"本3",5} )
	books=append(books,BookJson{4,"本3",5} )

	res, err := json.Marshal(books)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main(){
  	// 「/a」に対して処理を追加
  	http.HandleFunc("/books", booksJson)
  	// 8080ポートで起動
	http.ListenAndServe(":18080", nil)
}