// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"flag"
	"log"
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
	"github.com/gorilla/websocket"
  "github.com/yookoala/realpath"
	"encoding/json"
	"math/rand"
)

var brands = readBrands()
var postCodes = readPostCodes()

var addr = flag.String("addr", "localhost:8888", "http service address")

var upgrader = websocket.Upgrader{
  ReadBufferSize:  1024,
  WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool { return true },
}

type Sale struct {
	IconUrl string
	Lat     float64
	Lon     float64
}

type Brand struct {
	Name    string
	IconUrl string
}

type PostCode struct {
	PostCode  string
	Lat       float64
	Lon       float64
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
    sleepTime := rand.Intn(10)
		time.Sleep(time.Duration(sleepTime) * time.Second)

		sales := getProducts(1, rand.Intn(10) + 1)
		salesJson, err := json.Marshal(sales)

		if err != nil {
			fmt.Println(err)
			return
		}

		err = c.WriteMessage(websocket.TextMessage, salesJson)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}


func getProducts(lastId int, numProducts int)(sales []Sale) {
	sales = []Sale{}

	for i := 0; i < numProducts; i++ {
    randBrand := rand.Intn(len(brands))
  	randZip := rand.Intn(len(postCodes))
		mySale := Sale {
			IconUrl: "http://127.0.0.1:8888/images/" + brands[randBrand].IconUrl,
			Lat:     postCodes[randZip].Lat,
			Lon:     postCodes[randZip].Lon,
		}
		sales = append(sales, mySale)
	}

	return
}

func readBrands()(brands []Brand) {
	brandJson, err := ioutil.ReadFile("brands.json")
	if err != nil {
		log.Print("Could not read brands data:", err)
		return
	}
	json.Unmarshal(brandJson, &brands)
	return
}

func readPostCodes()(postCodes []PostCode) {
	postCodesJson, err := ioutil.ReadFile("zip_point.json")
	if err != nil {
		log.Print("Could not read postCodes data:", err)
		return
	}
	json.Unmarshal(postCodesJson, &postCodes)
	return
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", echo)

  path, err := realpath.Realpath("images")
  if err != nil {
    log.Print("Could not resolve path:", err)
    return
  }
  fs := http.FileServer(http.Dir(path))
  http.Handle("/images/", http.StripPrefix("/images", fs))

	log.Fatal(http.ListenAndServe(*addr, nil))
}
