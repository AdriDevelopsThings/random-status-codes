package pkg

import "net/http"

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(getRandomStatusCode())
}
