package server

import (
	"calemaric/mini-home-controller/internal/actions"
	"calemaric/mini-home-controller/internal/controls"
	"flag"
	"html/template"
	"log"
	"net/http"
)

var configFile = flag.String("config-path", ".config.yaml", "config file path")

var addr = flag.String("addr", "0.0.0.0:8080", "address to serve")

func Serve() {
	flag.Parse()
	log.Print(configFile)
	actions.InitializeActions()
	controls.InitializeControls()

	connections := SetupConnections()
	go connections.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		HandleConnection(connections, w, r)
	})

	templ, templateError := template.ParseFiles("./web/src/index.html", "./web/templates/controls/card.html")

	if templateError != nil {
		log.Fatal("Cannot parse index", templateError)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templ.Execute(w, controls.GetAllControls())

		if err != nil {
			log.Fatal("Cannot parse template", err)
		}
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
