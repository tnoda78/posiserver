package posiserver

import (
	"fmt"
	"net/http"

	"github.com/tnoda78/posiserver/configuration"
	"github.com/tnoda78/posiserver/environment"
	"github.com/tnoda78/posiserver/writer"
)

// Posiserver is simple struct
type Posiserver struct {
}

func NewPosiserver() *Posiserver {
	return &Posiserver{}
}

func (server *Posiserver) Run() {
	env := environment.NewEnvironment()
	config := configuration.NewConfiguration()
	config.LoadConfigFile(env)
	go writer.WriteTargetsLoop(config)

	// Listen port 3000
	fs := http.FileServer(http.Dir(config.RootPath))
	http.Handle("/", handleCORS(fs))
	fmt.Println("Listen port 3000.")
	http.ListenAndServe(":3000", nil)
}

func handleCORS(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
