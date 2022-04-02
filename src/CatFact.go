package main

import "fmt"
import (
//    "io/ioutil"
    "encoding/json"
    "net/http"
    "log"
)

func logError(error error, number int) {
    if error != nil {
        fmt.Println("erro")
        log.Fatalln(error)
    }
}
func main() {
    resp, err := http.Get("https://some-random-api.ml/facts/cat")
    logError(err, 1)
   
    defer resp.Body.Close()

//    body, err := ioutil.ReadAll(resp.Body)
//    logError(err, 2)

    var result map[string]interface{}

    json.NewDecoder(resp.Body).Decode(&result)
    fmt.Println(result["fact"])


}

// arrays sempre começam do zero
//para declarar a array é simples, []tipo{valores}
