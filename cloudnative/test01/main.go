package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func getinfo(w http.ResponseWriter, r *http.Request) {
	n := fmt.Sprintf("client IP :%s, status code: %s", r.Host, w.Header().Get("status"))
	fmt.Println(n)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	rheader := r.Header
	io.WriteString(w, "========>Details of the http response header<========\n")
	for k, v := range rheader {
		w.Header().Set(k, v[0])
		//fmt.Fprintln(w, k, ":", v[0])
		io.WriteString(w, fmt.Sprintf("%s:%s\n", k, v))
	}
	getinfo(w, r)
}

func versionHandle(w http.ResponseWriter, r *http.Request) {
	//ver:=os.Getenv("GOVERSION")
	ver, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, fmt.Sprintf("%s", ver))
	getinfo(w, r)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
	w.WriteHeader(200)
	getinfo(w, r)
}

func main() {
	flag.Set("v", "4")
	//flag.Parse()
	glog.V(0).Info("Starting http server...")
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/version", versionHandle)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":8000", nil)
}
