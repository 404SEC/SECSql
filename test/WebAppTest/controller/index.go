package controller

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("http://127.0.0.1:8881/hiveExec"))
}
