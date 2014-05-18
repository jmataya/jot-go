package restful

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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

func (API) methodNotFound(values url.Values) (int, interface{}) {
	data := map[string]string{"error": "Method Not Found"}
	return 404, data
}

func (api *API) matchAction(restfulType RestfulType, path string, method string) (bool, func(values url.Values) (int, interface{})) {
	if restfulType.IsCollectionMatch(path) {
		if method == GET {
			return true, restfulType.List
		} else if method == POST {
			return true, restfulType.Create
		} else {
			return false, api.methodNotFound
		}
	} else if restfulType.IsMemberMatch(path) {
		if method == GET {
			return true, restfulType.Show
		} else if method == PUT {
			return true, restfulType.Update
		} else if method == DELETE {
			return true, restfulType.Destroy
		} else {
			return false, api.methodNotFound
		}
	} else {
		return false, api.methodNotFound
	}
}

func (api *API) handleRequest(rw http.ResponseWriter, request *http.Request) {
	actualPath := request.URL.Path
	method := request.Method

	success, action := api.matchAction(api.registeredTypes["Note"], actualPath, method)

	if success {
		parseError := request.ParseForm()
		if parseError != nil {
			fmt.Fprintf(rw, "Error processing request %s", parseError)
			return
		} else {
			var statusCode int
			var data interface{}
			values := request.Form

			statusCode, data = action(values)

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
