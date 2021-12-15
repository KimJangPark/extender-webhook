package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
	"flag"
	"github.com/julienschmidt/httprouter"
)

var t int = 0
var score1 int = 0
var score2 int = 0
var score3 int = 0

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	tVal := flag.Int("threshold",0, "an int")
	node1:= flag.Int("1", 0, "an int")
	node2:= flag.Int("2", 0, "an int")
	node3:= flag.Int("3", 0, "an int")

	flag.Parse()
	log.Print(*tVal)
	log.Print(*node1)
	log.Print(*node2)
	log.Print(*node3)
	t = *tVal
	score1 = *node1
	score2 = *node2
	score3 = *node3

	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/filter", Filter)
	router.POST("/prioritize", Prioritize)

	log.Print("info: server starting on the port :8888")
	if err := http.ListenAndServe(":8888", router); err != nil {
		log.Fatal(err)
	}

}
