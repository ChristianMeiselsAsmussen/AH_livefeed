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
	"github.com/gorilla/websocket"
	"encoding/json"
	"math/rand"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

type Sale struct {
	IconUrl string
	Lat     string
	Lon     string
}
var brands = map[string]string {
	"Nike":"http://pngimg.com/uploads/nike/nike_PNG12.png",
	"Adidas":"http://pngimg.com/uploads/adidas/adidas_PNG21.png",
	"Hummel":"http://www.freelogovectors.net/wp-content/uploads/2018/03/hummel-logo.png",
	"NewBalance":"https://seeklogo.com/images/N/New_Balance-logo-F34722CB97-seeklogo.com.png",
	"UnderArmour":"https://upload.wikimedia.org/wikipedia/commons/thumb/4/44/Under_armour_logo.svg/2000px-Under_armour_logo.svg.png",
}
//var post

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
	numProducts := rand.Intn(10)
	randBrand := rand.Intn(len(brands))
	randPostcode := rand.Intn(9999-1000)+1000
	log.Print("numProducts:", numProducts)
	log.Print("randNum:", randPostcode)
	log.Print("randBrand:", randBrand)
	mySale := Sale {
		IconUrl: brands,
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

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
