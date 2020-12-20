package server

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Counter struct {
	Value int
}

func (c *Counter) Increment(v int) {
	c.Value += v
}

func (c *Counter) Decrement(v int) {
	c.Value -= v
}

type CounterBody struct {
	Value int
}

type CounterResponse struct {
	Value int
}

func HandleCounterIncrement(c *Counter) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var body CounterBody
		err := json.NewDecoder(request.Body).Decode(&body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		c.Increment(body.Value)

		err = json.NewEncoder(writer).Encode(CounterResponse{Value: c.Value})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func HandleCounterDecrement(c *Counter) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var body CounterBody
		err := json.NewDecoder(request.Body).Decode(&body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		c.Decrement(body.Value)

		err = json.NewEncoder(writer).Encode(CounterResponse{Value: c.Value})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func HandleCounterValue(c *Counter) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		err := json.NewEncoder(writer).Encode(CounterResponse{Value: c.Value})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func AddCounterApiRoutes(router *httprouter.Router, baseRoute string, counter *Counter) {
	router.GET(fmt.Sprintf("%s/counter", baseRoute), HandleCounterValue(counter))
	router.POST(fmt.Sprintf("%s/counter/increment", baseRoute), HandleCounterIncrement(counter))
	router.POST(fmt.Sprintf("%s/counter/decrement", baseRoute), HandleCounterDecrement(counter))
}
