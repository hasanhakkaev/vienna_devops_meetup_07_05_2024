package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
)

var b = `
██╗   ██╗██╗███████╗███╗   ██╗███╗   ██╗ █████╗
██║   ██║██║██╔════╝████╗  ██║████╗  ██║██╔══██╗
██║   ██║██║█████╗  ██╔██╗ ██║██╔██╗ ██║███████║
╚██╗ ██╔╝██║██╔══╝  ██║╚██╗██║██║╚██╗██║██╔══██║
 ╚████╔╝ ██║███████╗██║ ╚████║██║ ╚████║██║  ██║
  ╚═══╝  ╚═╝╚══════╝╚═╝  ╚═══╝╚═╝  ╚═══╝╚═╝  ╚═╝

██████╗ ███████╗██╗   ██╗ ██████╗ ██████╗ ███████╗
██╔══██╗██╔════╝██║   ██║██╔═══██╗██╔══██╗██╔════╝
██║  ██║█████╗  ██║   ██║██║   ██║██████╔╝███████╗
██║  ██║██╔══╝  ╚██╗ ██╔╝██║   ██║██╔═══╝ ╚════██║
██████╔╝███████╗ ╚████╔╝ ╚██████╔╝██║     ███████║
╚═════╝ ╚══════╝  ╚═══╝   ╚═════╝ ╚═╝     ╚══════╝

███╗   ███╗███████╗███████╗████████╗██╗   ██╗██████╗
████╗ ████║██╔════╝██╔════╝╚══██╔══╝██║   ██║██╔══██╗
██╔████╔██║█████╗  █████╗     ██║   ██║   ██║██████╔╝
██║╚██╔╝██║██╔══╝  ██╔══╝     ██║   ██║   ██║██╔═══╝
██║ ╚═╝ ██║███████╗███████╗   ██║   ╚██████╔╝██║
╚═╝     ╚═╝╚══════╝╚══════╝   ╚═╝    ╚═════╝ ╚═╝
`

func main() {

	banner.InitString(colorable.NewColorableStdout(), true, true, b)
	fmt.Print("\n")

	http.HandleFunc("/", handler)
	http.HandleFunc("/time", timeHandler)
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", logRequest(http.DefaultServeMux)))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Vienna DevOps Meetup")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "Current time: %s", currentTime)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
