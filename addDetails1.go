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
		Log struct{
			Version string `json:"version"`
			Creator struct {
				Name string `json:"name"`
				Version string `json:"version"`
				Comment string `json:"comment"`
			} `json:"creator"`
			Browser struct{
				Name string `json:"name"`
				Version string `json:"version"`
				Comment string `json:"comment"`
			} `json:"browser"`
			Pages [] struct{
				StartedDateTime string `json:"startedDateTime"`
				Id string `json:"id"`
				Title string `json:"title"`
				PageTimings struct{
					OnContentLoad int64 `json:"onContentLoad"`
					OnLoad int64 `json:"onLoad"`
					Comment string `json:"comment"`
				} `json:"pageTimings"`
				Comment string `json:"comment"`
			} `json:"pages"`
			Entries [] struct{
				Pageref string `json:"pageref"`
				StartedDateTime string `json:"startedDateTime"`
				Time float64 `json:"time"`
				Request struct{
					Method string `json:"method"`
					Url string `json:"url"`
					HttpVersion string `json:"httpVersion"`
					Cookies [] struct{
						Name string `json:"name"`
						Value string `json:"value"`
						Path string `json:"path"`
						Domain string `json:"domain"`
						Expires string `json:"expires"`
						HttpOnly bool `json:"httpOnly"`
						Secure bool `json:"secure"`
						Comment string `json:"comment"`
					} `json:"cookies"`
					Headers [] struct{
						Name string `json:"name"`
						Value string `json:"value"`
						Comment string `json:"comment"`
					} `json:"headers"`
					QueryString [] struct{
						Name string `json:"name"`
						Value string `json:"value"`
						Comment string `json:"comment"`
					} `json:"queryString"`
					PostData struct{
						MimeType string `json:"mimeType"`
						Params [] struct{
							Name string `json:"name"`
							Value string `json:"value"`
							FileName string `json:"fileName"`
							ContentType string `json:"contentType"`
							Comment string `json:"comment"`
						} `json:"params"`
						Text string `json:"text"`
						Comment string `json:"comment"`
					} `json:"postData"`
					HeadersSize int64 `json:"headersSize"`
					BodySize int64 `json:"bodySize"`
					Comment string `json:"comment"`
				} `json:"request"`
				Response struct{
					Status int64 `json:"status"`
					StatusText string `json:"statusText"`
					HttpVersion string `json:"httpVersion"`
					Cookies [] struct{
						Name string `json:"name"`
						Value string `json:"value"`
						Path string `json:"path"`
						Domain string `json:"domain"`
						Expires string `json:"expires"`
						HttpOnly bool `json:"httpOnly"`
						Secure bool `json:"secure"`
						Comment string `json:"comment"`
					} `json:"cookies"`
					Headers [] struct{
						Name string `json:"name"`
						Value string `json:"value"`
						Comment string `json:"comment"`
					} `json:"headers"`
					Content struct{
						Size int64 `json:"size"`
						Compression int64 `json:"compression"`
						MimeType string `json:"mimeType"`
						Text string `json:"text"`
						Comment string `json:"comment"`
					} `json:"content"`
					RedirectURL string `json:"redirectURL"`
					HeadersSize int64 `json:"headersSize"`
					BodySize int64 `json:"bodySize"`
					Comment string `json:"comment"`
				} `json:"response"`
				Cache struct{
					BeforeRequest struct{
						Expires string `json:"expires"`
						LastAccess string `json:"lastAccess"`
						ETag string `json:"eTag"`
						HitCount int64 `json:"hitCount"`
						Comment string`json:"comment"`
					} `json:"beforeRequest"`
					AfterRequest struct{
						Expires string `json:"expires"`
						LastAccess string `json:"lastAccess"`
						ETag string `json:"eTag"`
						HitCount int64 `json:"hitCount"`
						Comment string`json:"comment"`
					} `json:"afterRequest"`
					Comment string `json:"comment"`
				} `json:"cache"`
				Timings struct{
					Blocked float64 `json:"blocked"`
					Dns float64 `json:"dns"`
					Connect float64 `json:"connect"`
					Send float64 `json:"send"`
					Wait float64 `json:"wait"`
					Receive float64 `json:"receive"`
					Ssl float64 `json:"ssl"`
					Comment string `json:"comment"`
				} `json:"timings"`
				ServerIPAddress string `json:"serverIPAddress"`
				Connection string `json:"connection"`
				Comment string `json:"comment"`
			} `json:"entries"`
			Comment string `json:"comment"`
		} `json:"log"`
	} `json:"details"`
//	Test [] struct{
		
	
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
		if detail.ProbeId != 0 && detail.ProbeOwner != "" && detail.Url != "" {

			
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
