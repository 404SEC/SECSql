package main

import (
	"net/http"

	c "github.com/404SEC/SECSql/test/WebAppTest/controller"
	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	//r.NotFoundHandler = http.HandlerFunc(c.LoginHandler)

	// read app.pid
	//r.HandleFunc("/login", c.LoginHandler)

	r.HandleFunc("/", c.ConnHandler)
	r.HandleFunc("/api", c.Load)
	r.HandleFunc("/exec", c.ExecHandler)
	//pid:=os.Getpid()
	//write pid
	//kill pid
	return r
}

//var conf models.Config
func main() {
	r := newRouter()
	//Logger = c.Logger
	http.ListenAndServe("0.0.0.0:8881", r)
	//Logger.Info("Service Start On" + c.Conf.IP + ":" + c.Conf.Port)
}
