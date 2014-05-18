package restful

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type API struct{}

func (api *API) Abort(rw http.ResponseWriter, statusCode int) {
	rw.WriteHeader(statusCode)
}

func (api *API) requestHandler(resource Resource) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		method := request.Method
		var parseError = request.ParseForm()
		if parseError != nil {
			fmt.Fprintf(rw, "Error processing request %s", parseError)
		} else {
			var statusCode int
			var data interface{}
			values := request.Form

			switch method {
			case "GET":
				statusCode, data = resource.Get(values)
			case "POST":
				statusCode, data = resource.Post(values)
			case "PUT":
				statusCode, data = resource.Put(values)
			case "DELETE":
				statusCode, data = resource.Delete(values)
			default:
				api.Abort(rw, 405)
				return
			}

			content, err := json.Marshal(data)
			if err != nil {
				api.Abort(rw, 500)
				return
			}
			rw.WriteHeader(statusCode)
			rw.Write(content)
		}
	}
}

func (api *API) AddResource(resource Resource, path string) {
	http.HandleFunc(path, api.requestHandler(resource))
}

func (api *API) Start(port int) {
	portString := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(portString, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
