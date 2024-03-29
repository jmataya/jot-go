package restful

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type API struct {
	registeredTypes map[string]RestfulController
}

func (api *API) Abort(rw http.ResponseWriter, statusCode int) {
	rw.WriteHeader(statusCode)
}

func (API) methodNotFound(values url.Values, params map[string]string) (int, interface{}) {
	data := map[string]string{"error": "Method Not Found"}
	return 404, data
}

func (API) getCollectionPath(basePath string) string {
	/*
	 * This essentially works by looking at the base path and stripping off an
	 * ID parameter if it exists on the end of the string.
	 */
	lastIdRegexStr := "{[a-zA-Z0-9_-]+}/?$"
	lastIdRegexMatcher := regexp.MustCompile(lastIdRegexStr)
	return lastIdRegexMatcher.ReplaceAllString(basePath, "")
}

func (API) pathIsMatch(base string, actual string) bool {
	// First, interpolate the placeholders that are in for the strings.
	keyMatcher := regexp.MustCompile("{[a-zA-Z0-9_-]+}")
	interpolatedStr := keyMatcher.ReplaceAllString(base, "[a-zA-Z0-9_-]+")

	// Second, clean up the string:
	// 1. Make sure that the last "/" is optional
	// 2. Make sure that nothing can come after this string.
	slashRegexMatcher := regexp.MustCompile("/$")
	interpolatedStr = slashRegexMatcher.ReplaceAllString(interpolatedStr, "")
	interpolatedStr += "/?$"

	// Finally, do the actual match
	valueMatcher := regexp.MustCompile(interpolatedStr)
	return valueMatcher.MatchString(actual)
}

func (API) getPathParams(basePath string, actualPath string) map[string]string {
	// First break each path into its individual pieces.
	basePieces := strings.Split(basePath, "/")
	actualPieces := strings.Split(actualPath, "/")

	if len(basePieces) != len(actualPieces) {
		return map[string]string{"error": "base and actual paths do not match"}
	}

	keyRegexString := "{[a-zA-Z0-9_-]+}"
	keyRegexMatcher := regexp.MustCompile(keyRegexString)
	params := map[string]string{}

	for i := 0; i < len(basePieces); i++ {
		if keyRegexMatcher.MatchString(basePieces[i]) {
			key := strings.Trim(basePieces[i], " {}")
			value := strings.Trim(actualPieces[i], " ")
			params[key] = value
		}
	}

	return params
}

func (api *API) IsCollectionMatch(resourcePath string, path string) bool {
	collectionPath := api.getCollectionPath(resourcePath)
	return api.pathIsMatch(collectionPath, path)
}

func (api *API) IsMemberMatch(resourcePath string, path string) bool {
	return api.pathIsMatch(resourcePath, path)
}
func (api *API) matchAction(restfulType RestfulController, resourcePath string, path string, method string) (bool, func(values url.Values, params map[string]string) (int, interface{})) {
	if api.IsCollectionMatch(resourcePath, path) {
		if method == GET {
			return true, restfulType.List
		} else if method == POST {
			return true, restfulType.Create
		} else {
			return false, api.methodNotFound
		}
	} else if api.IsMemberMatch(resourcePath, path) {
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

	var success bool
	var action func(values url.Values, params map[string]string) (int, interface{})

	for resourcePath, regType := range api.registeredTypes {
		success, action = api.matchAction(regType, resourcePath, actualPath, method)
		if success {
			parseError := request.ParseForm()
			if parseError != nil {
				api.Abort(rw, 500)
				return
			} else {
				var statusCode int
				var data interface{}
				values := request.Form
				params := api.getPathParams(resourcePath, actualPath)

				statusCode, data = action(values, params)

				content, err := json.Marshal(data)
				if err != nil {
					api.Abort(rw, 500)
					return
				}

				rw.WriteHeader(statusCode)
				rw.Write(content)
				return
			}
		}
	}

	// If we hit this point, the resource hasn't been found.
	content, err := json.Marshal(map[string]string{"message": "Not found"})
	if err != nil {
		api.Abort(rw, 500)
		return
	}
	rw.WriteHeader(404)
	rw.Write(content)
}

func (api *API) RegisterRestfulController(basePath string, restfulType RestfulController) {
	if api.registeredTypes == nil {
		api.registeredTypes = map[string]RestfulController{}
	}
	api.registeredTypes[basePath] = restfulType
}

func (api *API) Start(port int) {
	portString := fmt.Sprintf(":%d", port)
	http.HandleFunc("/", api.handleRequest)
	err := http.ListenAndServe(portString, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
