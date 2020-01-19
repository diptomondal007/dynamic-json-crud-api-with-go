package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

 var data map[string]interface{}

func getAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	key := params["key"]
	log.Println(key)
	if len(key) >0{
		for d, _ := range data{
			log.Println(d)
			if string(d) == key{
				da := data[key].(interface{})
				json.NewEncoder(w).Encode(da)
			}
		}
	}else {
		json.NewEncoder(w).Encode(data)

	}

}

func addData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	responseBody, _ := ioutil.ReadAll(r.Body)
	//newData := new(interface{})
	log.Println(json.Unmarshal(responseBody, &data))
	//res , err := json.Marshal(&newData)

	//data = append(data, newData)
	json.NewEncoder(w).Encode(&data)


}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{key}", getAllData).Methods("GET")

	router.HandleFunc("/", addData).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
