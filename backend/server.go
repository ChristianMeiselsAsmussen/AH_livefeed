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
	"encoding/json"
)

var brands = readBrands()
var postCodes = readPostCodes()

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

type Sale struct {
	IconUrl string
	Lat     string
	Lon     string
}

type Brand struct {
	Name    string
	IconUrl string
}

type PostCode struct {
	PostCode  string
	Lat       string
	Lon       string
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		time.Sleep(2 * time.Second)

		sales := getProducts(1)
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

func getProducts(lastId int)(sales []Sale) {
	sales = []Sale{}

	mySale := Sale {
		IconUrl: "xxx",
		Lat:     "yyy",
		Lon:     "zzz",
	}

	sales = append(sales, mySale)

	mySale = Sale {
		IconUrl: "dfgdfsg",
		Lat:     "ydfgyy",
		Lon:     "dfgdfg",
	}

	sales = append(sales, mySale)
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
	log.Fatal(http.ListenAndServe(*addr, nil))
}
