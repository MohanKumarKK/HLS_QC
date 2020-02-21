package main

import (
        "fmt"
	"encoding/json"
        "io/ioutil"
	"net/http"
	"regexp"
)

type Detail struct{
	ProbeId int64 `json:"probeId"`
	ProbeOwner string `json:"probeOwner"`
	Url string `json:"url"`
	Ts_Segment string `json:"ts_Segment"`
	ErrorStatus int64 `json:"errorStatus"`
	ErrorDescription string `json:"errorDescription"`
	Timestamp string `json:"timestamp"`
}
func addDetails(w http.ResponseWriter, r *http.Request){

        w.Header().Set("content-type", "application/json")

	if r.Method == "POST"{

		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}


		var detail Detail
		err = json.Unmarshal(body, &detail)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if detail.ProbeId != 0 && detail.ProbeOwner != "" && detail.Url != "" && detail.Ts_Segment != "" && detail.ErrorStatus != 0 && detail.ErrorDescription != "" && detail.Timestamp != ""{


			matched, err:= regexp.MatchString(`[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) (2[0-3]|[01][0-9]):[0-5][0-9]:[0-5][0-9]`,detail.Timestamp)
			if(err != nil){
				fmt.Println("Incorrect regular expression")
			}
			if(detail.ErrorStatus>=100 && detail.ErrorStatus<=599 && matched == true){
				w.Write([]byte(`{"message":"Successfully added"}`))
				fmt.Println(detail)
			}else{
				 w.WriteHeader(http.StatusBadRequest)
	                         w.Write([]byte(`{"message":"Unsuccessful-Incorrect feild data"}`))

			}

		}else{
			w.WriteHeader(http.StatusBadRequest)
	                w.Write([]byte(`{"message":"Unsuccessful-Fields missing"}`))
			//fmt.Println(detail)
		}

	}else{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"Bad request"}`))
	}
}

func main(){
	http.HandleFunc("/api/v1/addDetails",addDetails)
	http.ListenAndServe(":5000",nil)
}
