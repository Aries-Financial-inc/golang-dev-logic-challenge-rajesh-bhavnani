package main

import (
    "log"
    "net/http"
    "github.com/Aries-Financial-inc/golang-dev-logic-challenge-rajesh-bhavnani/controllers"
)

func main() {
    http.HandleFunc("/analyze", controllers.AnalysisHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
