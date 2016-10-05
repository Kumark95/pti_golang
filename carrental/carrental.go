package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "io"
    "io/ioutil"
)

type ResponseMessage struct {
    Car_maker string
    Car_model string
    Number_days int
    Number_units int
    Final_price int
}

type RequestMessage struct {
    CarMaker string
    CarModel string
    NumberDays int
    NumberUnits int
}

func main() {

router := mux.NewRouter().StrictSlash(true)
router.HandleFunc("/", Index)
router.HandleFunc("/new/", newFunc)
router.HandleFunc("/list/{param}", listFunc)


log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Car rental")
}

func newFunc(w http.ResponseWriter, r *http.Request) {
    var requestMessage RequestMessage
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &requestMessage); err != nil {
        fmt.Fprintln(w, "Prueba")
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        fmt.Fprintln(w, "Prueba2")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    } else {
        fmt.Fprintln(w, "Successfully received request with Car maker =", requestMessage.CarMaker, "Car model", requestMessage.CarModel, "()Number of days", requestMessage.NumberDays, "Number of units", requestMessage.NumberUnits)
    }
}

func listFunc(w http.ResponseWriter, r *http.Request) {
    /*
    var requestMessage RequestMessage
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &requestMessage); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    } else {
        fmt.Fprintln(w, "Successfully received request with Field1 =", requestMessage.Field1)
        fmt.Println(r.FormValue("queryparam1"))
    }
    */
}