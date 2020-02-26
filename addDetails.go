package main

import (
        "fmt"
	"encoding/json"
        "io/ioutil"
	"net/http"
	
)

// structure to store data
type Detail struct{
	ProbeId int64 `json:"probeId"`
	ProbeOwner string `json:"probeOwner"`
	Url string `json:"url"`
	Details struct{
		Content string `json:"_content"`
		StatusCode int64 `json:"status_code"`
		Headers_ struct{
			ContentType string `json:"Content-Type"`
			TransferEncoding string `json:"Transfer-Encoding"`
			Connection_ string `json:"Connection"`
			Date_ string `json:"Date"`
			Server_ string `json:"Server"`
			CacheControl string `json:"Cache-Control"`
			LastModified string `json:"Last-Modified"`
			AccessControlAllowCredentials string `json:"Access-Control-Allow-Credentials"`
			ContentEncoding string `json:"Content-Encoding"`
			Vary_ string `json:"Vary"`
			XCache string `json:"X-Cache"`
			Via_ string `json:"Via"`
			XAmzCfPop string `json:"X-Amz-Cf-Pop"`
			XAmzCfId string `json:"X-Amz-Cf-Id"`
		} `json:"headers"`
		Url string `json:"url"`
		Encoding [] string `json:"encoding"`
		History [] string `json:"history"`
		Cookies [] string `json:"cookies"`
		Elapsed [] string `json:"elapsed"`
		Request [] string `json:"request"`
		Reason string `json:"reason"`
		
	} `json:"details"`
}


func addDetails(w http.ResponseWriter, r *http.Request){

	//set response header
        w.Header().Set("content-type", "application/json")

	if r.Method == "POST"{

		//read post request body
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}


		var detail Detail
		//parse request body into structure of type Detail
		err = json.Unmarshal(body, &detail)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		//check for missing feild
		if detail.ProbeId != 0 && detail.ProbeOwner != "" && detail.Url != ""{

				w.Write([]byte(`{"message":"Successfully added"}`))
				fmt.Println(detail)
			

		}else{
			w.WriteHeader(http.StatusBadRequest)
	                w.Write([]byte(`{"message":"Unsuccessful-Fields missing"}`))
			//fmt.Println(detail)
		}

	}else{ //for request other than POST type
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"Bad request"}`))
	}
}

func main(){
	http.HandleFunc("/api/v1/Details",addDetails)
	http.ListenAndServe(":5000",nil)
}
