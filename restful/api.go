package restful

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type API struct {
	registeredTypes map[string]RestfulType
}

func (api *API) Abort(rw http.ResponseWriter, statusCode int) {
	rw.WriteHeader(statusCode)
}

func (api *API) handleRequest(rw http.ResponseWriter, request *http.Request) {
	actualPath := request.URL.Path
	method := request.Method

	success, action := api.registeredTypes["Note"].ActionMatch(actualPath, method)

	if success {
		parseError := request.ParseForm()
		if parseError != nil {
			fmt.Fprintf(rw, "Error processing request %s", parseError)
			return
		} else {
			var statusCode int
			var data interface{}
			values := request.Form

			switch action {
			case LIST:
				statusCode, data = api.registeredTypes["Note"].List(values)
			case SHOW:
				statusCode, data = api.registeredTypes["Note"].Show(values)
			case CREATE:
				statusCode, data = api.registeredTypes["Note"].Create(values)
			case UPDATE:
				statusCode, data = api.registeredTypes["Note"].Update(values)
			case DESTROY:
				statusCode, data = api.registeredTypes["Note"].Destroy(values)
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
	} else {
		content, err := json.Marshal(map[string]string{"message": "Not found"})
		if err != nil {
			api.Abort(rw, 500)
			return
		}
		rw.WriteHeader(200)
		rw.Write(content)
	}
}

// func (api *API) requestHandler(resource Resource) http.HandlerFunc {
// 	return func(rw http.ResponseWriter, request *http.Request) {
// 		method := request.Method
// 		var parseError = request.ParseForm()
// 		if parseError != nil {
// 			fmt.Fprintf(rw, "Error processing request %s", parseError)
// 		} else {
//
// 			content, err := json.Marshal(data)
// 			if err != nil {
// 				api.Abort(rw, 500)
// 				return
// 			}
// 			rw.WriteHeader(statusCode)
// 			rw.Write(content)
// 		}
// 	}
// }

// func (api *API) AddResource(resource Resource, path string) {
// 	http.HandleFunc(path, api.requestHandler(resource))
// }

func (api *API) RegisterRestfulType(name string, restfulType RestfulType) {
	if api.registeredTypes == nil {
		api.registeredTypes = map[string]RestfulType{}
	}
	api.registeredTypes[name] = restfulType
}

func (api *API) Start(port int) {
	portString := fmt.Sprintf(":%d", port)
	http.HandleFunc("/", api.handleRequest)
	err := http.ListenAndServe(portString, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
