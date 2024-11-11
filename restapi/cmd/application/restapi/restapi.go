package restapi

/*

Testing:
Add element:		curl http://localhost:8000/item/key -X PUT -L -d '{"data":{"value":"value"}}'
Get element:		curl http://localhost:8000/item/key
Get statistics:		curl http://localhost:8000/stat
Delete:				curl http://localhost:8000/item/key -X DELETE
Increase:			curl http://localhost:8000/item/key/incr/2 -X POST
Reverse:			curl http://localhost:8000/item/key/reverse -X POST
Sort:				curl http://localhost:8000/item/key/sort -X POST
Dedupl/icate:		curl http://localhost:8000/item/key/dedup -X POST

Complex curl: /test.sh

*/

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mityamentor/utils"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var m sync.Mutex
var cache map[string]string
var methodsCounter map[string]int

type Item struct {
	Data ItemData `json:"data"`
}

type ItemData struct {
	Value string `json:"value"`
}

func StartAPI() {
	cache = make(map[string]string)
	methodsCounter = make(map[string]int)

	http.HandleFunc("/item/", parseItemMethod) // TODO: make handler for each method - make separate file for each handler 
	http.HandleFunc("/stat", stat)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseItemMethod(w http.ResponseWriter, r *http.Request) {
	httpMethod := r.Method

	urlArray := utils.ParseURL(strings.Split(r.URL.Path, "/"))

	if !checkURL(urlArray, httpMethod) {
		fmt.Fprintf(w, "Method not supported\n")
		return
	}

	var action string
	key := urlArray[1]

	switch len(urlArray) {
	case 2:  // /item/{key}

		switch httpMethod {

		case "GET":  // GET /item/{key}
			action = "GET element"

			if val, exists := cache[key]; !exists {
				fmt.Fprint(w, "No such key in cache\n")
				break
			} else {
				fmt.Fprintf(w, "%v\n", val) 
			}

		case "PUT": // PUT /item/{key}
			action = "PUT element"

			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Fprint(w, "Error while parsing request body\n")
				break
			}
			defer func() {r.Body.Close()}()

			var item Item
			err = json.Unmarshal(body, &item)
			if err != nil {
				fmt.Fprint(w, "Error while parsing request body to json\n")
				break
			}

			m.Lock()
			cache[key] = item.Data.Value
			m.Unlock()

		case "DELETE":  // DELETE /item/{key}
			action = "DELETE element"

			if _, exists := cache[key]; !exists {
				fmt.Fprint(w, "No such key in cache\n")
				break
			}

			m.Lock()
			delete(cache, key)
			m.Unlock()
		}

	case 3: // POST /item/{key}/...	
		var currentVal, newVal string

		key := urlArray[1]
		action = urlArray[2]

		if val, exists := cache[key]; !exists {
			fmt.Fprint(w, "No such key in cache\n")
			break
		} else {
			currentVal = val
		}

		switch action {
		case "reverse":
			newVal = utils.ReversreString(currentVal)
			fmt.Println(currentVal)
			fmt.Println(newVal)
		case "sort":
			newVal = utils.SortString(currentVal)
		case "dedup":
			newVal = utils.DeduplicateString(currentVal)
		}

		action = "POST " + action
		
		m.Lock()
		cache[key] = newVal
		m.Unlock()

	case 4: // POST /item/{key}/incr/{increment}
		action = "POST increase"

		var currentVal string
		key := urlArray[1]

		if val, exists := cache[key]; !exists {
			fmt.Fprint(w, "No such key in cache\n")
			break
		} else {
			currentVal = val
		}

		currentValInt, err := strconv.Atoi(currentVal)
		if err != nil {
			fmt.Fprint(w, "Key value should be int\n")
			break
		}

		incInt, err := strconv.Atoi(urlArray[3])
		if err != nil {
			fmt.Fprint(w, "Increment value should be int\n")
			break
		}

		m.Lock()
		cache[key] = fmt.Sprintf("%d", currentValInt + incInt)
		m.Unlock()
	}

	m.Lock()
	methodsCounter[action]++ // TODO make as middleware to decorate each function and not to duplicate
	m.Unlock()

}

func stat(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	methodsCounter["GET stat"]++
	m.Unlock()

	fmt.Fprintf(w, "%v", utils.SprintMapStringInt(methodsCounter))
}

func checkURL(urlArray []string, httpMethod string) bool {
	// TODO: such cases as .../item/sort/ are not processed correctly as "/" is trimmed => equal to .../item/sort (but it is not)
	if len(urlArray) == 2 && (httpMethod == "GET" || httpMethod == "PUT" || httpMethod == "DELETE") {
		return true
	}

	if len(urlArray) == 3 && httpMethod == "POST" {
		m := urlArray[2]
		if m == "reverse" || m == "sort" || m == "dedup" {
			return true
		}
	}

	if len(urlArray) == 4 && urlArray[2] == "incr" && httpMethod == "POST" {
		return true
	}
	return false
}