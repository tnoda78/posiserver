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
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.ListenAndServe(":3000", nil)
	fmt.Println("Listen port 3000.")
}
