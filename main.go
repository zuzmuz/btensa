package main

import (
	"fmt"
	// "net/http"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	var router = chi.NewRouter()

	router.Use(middleware.Logger)

	// http.ListenAndServe("0.0.0.0:8000", router)

	secret1 := Secret{"user", "zuz", false}
	secret2 := Secret{"website", "gmail.com", false}
	secret3 := Secret{"password", "whatecs", true}

	capsule := Capsule{"personal", []Secret{secret1, secret2, secret3}}

	jsonObject, error := json.Marshal(capsule)

	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Printf("what are we talking about here %s\n", jsonObject)

}
