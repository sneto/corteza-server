package http

import (
	"fmt"

	"net/http"
)

type Fortune struct{}

func (*Fortune) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fortune := "Fortune favors the prepared mind. - Louis Pasteur"
	fmt.Fprintf(w, fortune)
}
