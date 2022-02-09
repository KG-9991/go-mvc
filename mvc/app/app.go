package app

import (
	"net/http"

	"github.com/go-mvc/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8079", nil); err != nil {
		panic(err)
	}
}
