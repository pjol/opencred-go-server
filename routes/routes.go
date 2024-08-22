package routes

import (
	"fmt"
	"net/http"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Success!")
}
