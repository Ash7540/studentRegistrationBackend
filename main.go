package main

import (
	"fmt"
	"net/http"
	"studentregist/router"
)

func main() {
	fmt.Println("Student Registration Form")
	r := router.Router()
	fmt.Println("Server is getting started ...")
	http.ListenAndServe(":8000", r)
	fmt.Println("Listening at port 8000...")
}
