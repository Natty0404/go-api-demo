package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"go-api-demo/interfaceadapter/controller"
)

func GetBookHandler(goDemoAPI *controller.GoApiDemo) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bookID := r.URL.Query().Get("book_id")
		book, err := goDemoAPI.GetBook(r.Context(), bookID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}
