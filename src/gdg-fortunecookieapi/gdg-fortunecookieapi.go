package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type Fortune struct {
	Fortune string `json:"fortune"`
	Text    string `json:"text"`
}

var initTime = time.Now()

/*
 Main API server.
 Handles / and /api/fortune
*/
func main() {

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/api/fortune", apiHandler)
	appengine.Main()

}

/*
 API handler for endpoint: /api/fortune
 Exec's /usr/games/fortune and puts the result into a Fortune
 struct, jsonifies that struct, and returns.
 TODO: eventually allow for POST to create new fortunes
*/
func apiHandler(w http.ResponseWriter, r *http.Request) {

	var out []byte
	out, err := exec.Command("/usr/games/fortune").Output()
	if err != nil {
		out = []byte("This is not the fortune you are looking for.")
	}

	c := appengine.NewContext(r)
	log.Infof(c, string(out))
	fortuneString := simpleTextStrip(string(out))

	f := Fortune{Fortune: fortuneString, Text: string(out)}

	json, err := json.Marshal(f)

	w.Header().Add("Content-Type", "application/json")
	w.Write(json)
}

/*
 Convenience method for unescaping a string
 Ugly. Should find a way to unescape rather than iteration.
*/
func simpleTextStrip(s string) string {
	text := fmt.Sprintf("%s", s)

	replaces := map[string]string{
		"\n":     " ",
		"\t":     " ",
		"\"":     "'",
		`\u003c`: "<",
		`\u003e`: ">",
	}

	for k, v := range replaces {
		text = strings.Replace(text, k, v, -1)
	}

	return strings.TrimSpace(text)
}

/*
 Handles the index, returning a Golang template,
 filling in the uptime of the API server for fun.
*/
func handleIndex(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	c := appengine.NewContext(r)
	log.Infof(c, "Serving main page.")

	tmpl, _ := template.ParseFiles("web/tmpl/index.tmpl")

	tmpl.Execute(w, time.Since(initTime))
}
