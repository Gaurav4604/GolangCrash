package main

import (
	"encoding/json"
	"net/http"
)

func JSONError(res http.ResponseWriter, errVal map[string]string, code int) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("X-Content-Type-Options", "nosniff")
	res.WriteHeader(code)

	jsonStr, _ := json.Marshal(errVal)
	res.Write(jsonStr)
}
