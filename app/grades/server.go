package grades

import (
	"net/http"
)

func RegisterHandlers() {
	handler := new(studentsHandler)
	http.Handle("/students", handler)
}

type studentsHandler struct{}
