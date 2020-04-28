package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HelloRequest represents body of Hello request.
// swagger:parameters idOfHelloEndpoint
type HelloParamsWrapper struct {
	// in: query
	Name string `json:"name"`
}

// This text will appear as description of your response body.
// swagger:response helloResponse
type HelloResponse struct {
	// in:body
	Payload []*Hello `json:"body,omitempty"`
}

type Hello struct {
	Message *Message `json:"message"`
}

type Message struct {
	Body string `json:"words"`
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	// swagger:route POST /hello hello-tag idOfHelloEndpoint
	// Hello returns hello world string with given input.
	// responses:
	//   200: helloResponse
	// Consumes:
	// - application/json
	// Produces:
	// - application/json

	case http.MethodPost:
		queryValues := r.URL.Query()
		params := HelloParamsWrapper{
			Name: queryValues.Get("name"),
		}
		words := fmt.Sprintf("Hello %v", params.Name)
		message := &Message{
			Body: words,
		}
		hello := &Hello{
			Message: message,
		}
		resp := HelloResponse{}
		resp.Payload = append(resp.Payload, hello)
		outJSON, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "error marshalling response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(outJSON)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
