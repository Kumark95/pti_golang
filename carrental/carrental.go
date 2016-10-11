package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "io"
    "io/ioutil"
    "encoding/csv"
    "os"
	"strconv"
	"strings"
)


type RequestMessage struct {
    CarMaker string
    CarModel string
    NumberDays int
    NumberUnits int
}

var PRICE int = 500;

func main() {

router := mux.NewRouter().StrictSlash(true)
router.HandleFunc("/", Index)
router.HandleFunc("/new/", newFunc)
router.HandleFunc("/list/", listFunc)


log.Fatal(http.ListenAndServe(":8080", router))
}

//*****************************************************************************************

func writeToFile(w http.ResponseWriter, req RequestMessage) {
    file, err := os.OpenFile("rentals.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if err := json.NewEncoder(w).Encode(err); err != nil {
        panic(err)
    }
    writer := csv.NewWriter(file)

	var num string = strconv.Itoa(req.NumberDays)
	var units string = strconv.Itoa(req.NumberUnits)

	var data = []string{req.CarMaker, req.CarModel, num, units}
    writer.Write(data)
    writer.Flush()
    file.Close()
}

func readFromFile(w http.ResponseWriter){
	data, err := ioutil.ReadFile("rentals.csv")
	if err != nil {
		panic(err)
	}

	lineas := strings.Replace(string(data), "\n", " ", -1)
	lineasR := strings.Split(lineas, " ")


	for index,element := range lineasR {
		if index != len(lineasR)-1 {
			camps := strings.Split(element, ",")
			numI, _ := strconv.Atoi(camps[2])
			daysI, _ := strconv.Atoi(camps[3])
			finalPr := (PRICE * numI * daysI)
			fmt.Fprintln(w, "Car Maker:", camps[0], "\nCar Model:", camps[1], "\nNumber of days:", camps[2], "\nNumber of  units:", camps[3], "\nFINAL PRICE:", finalPr, "\n")
		}
	}
}

//*****************************************************************************************


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
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    } else {
        fmt.Fprintln(w, "Successfully received request with: \n· Car maker =", requestMessage.CarMaker, "\n· Car model = ", requestMessage.CarModel, "\n· Number of days = ", requestMessage.NumberDays, "\n· Number of units = ", requestMessage.NumberUnits)

		writeToFile(w, requestMessage)

    }
}

//curl -H "Content-Type: application/json" -d '{"CarMaker":"Audi", "CarModel":"A6", "NumberDays":3, "NumberUnits":2}' http://localhost:8080/new/


func listFunc(w http.ResponseWriter, r *http.Request) {
    readFromFile(w)
}
