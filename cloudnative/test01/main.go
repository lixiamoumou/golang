package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {

	rheader := r.Header
	for k, v := range rheader {
		w.Header().Set(k, v[0])
		fmt.Fprintln(w, k, ":", v[0])
	}
}

func versionHandle(w http.ResponseWriter, r *http.Request) {
	//ver:=os.Getenv("GOVERSION")
	ver, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, fmt.Sprintf("%s", ver))
}

func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/version", versionHandle)
	http.ListenAndServe(":8000", nil)
}
