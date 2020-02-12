package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	notedir := flag.String("notedir", "/tmp", "note dir")
	flag.Parse()
	*notedir += "/tempnotes"
	if _, err := os.Stat(*notedir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(*notedir, os.ModePerm); err != nil {
				log.Fatalf("Notedir %s can not write 1\n", *notedir)
			}
		} else {
			log.Fatalf("Notedir %s can not write 2\n", *notedir)
		}
	}
	if err := ioutil.WriteFile(*notedir+"/test", make([]byte, 0), os.ModePerm); err != nil {
		fmt.Println(err)
		log.Fatalf("Notedir %s can not write 3\n", *notedir)
	}

	h1 := func(w http.ResponseWriter, req *http.Request) {
		tpl := template.Must(template.New("letter").Parse(letter))
		type TplData struct {
			Data string
		}
		d := TplData{}

		key := req.URL.Path
		if _, err := os.Stat(*notedir + "/" + key); err == nil {
			bs, err := ioutil.ReadFile(*notedir + "/" + key)
			if err == nil {
				d.Data = string(bs)
			}
		}

		tpl.Execute(w, d)
	}
	h2 := func(w http.ResponseWriter, req *http.Request) {
		bs, err := ioutil.ReadAll(req.Body)
		req.Body.Close()
		if err != nil || len(bs) < 1 {
			return
		}
		h := md5.New()
		h.Write(bs)
		s := hex.EncodeToString(h.Sum(nil))

		key := string(s[:])
		ioutil.WriteFile(*notedir+"/"+key, bs, os.ModePerm)
		io.WriteString(w, key)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/save", h2)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, _ *http.Request) {

	})

	fmt.Printf("Note dir: %s\n", *notedir)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
