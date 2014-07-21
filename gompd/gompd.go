package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ernado/gomp"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Google struct {
		Token string `json:"token"`
	} `json:"google"`
	Apple struct {
		Gateway string `json:"gateway"`
		Key     string `json:"key"`
		Cert    string `json:"cert"`
	} `json:"apple"`
}

var (
	configPath = flag.String("cfg", "conf.json", "configuration file path")
	port       = flag.Int("port", 3100, "port")
	host       = flag.String("host", "", "host")
	prefix     = flag.String("prefix", "/push", "uri prefix")
	client     *gomp.Sender
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := client.Handle(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		fmt.Fprint(w, "Sent")
	}
}

func main() {
	flag.Parse()
	cfg := new(Config)
	f, err := os.Open(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(cfg); err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", cfg)
	appleConfig := gomp.APNSConfig{cfg.Apple.Gateway, cfg.Apple.Cert, cfg.Apple.Key}
	googleConfig := gomp.GCMConfig{cfg.Google.Token}
	client = gomp.New(appleConfig, googleConfig)

	router := httprouter.New()
	router.GET(*prefix, Index)

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
