// system under test for http handlers
package handlers

import "net/http"

func HandleGetFoo(w http.ResponseWriter, r *http.Request) {
	// some code
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("FOO"))
}
