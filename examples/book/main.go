package main

import (
	"log"

	ws "github.com/aopoltorzhicky/go_kraken/websocket"
)

func main() {
	c := ws.New()
	err := c.Connect()
	if err != nil {
		log.Fatal("Error connecting to web socket : ", err)
	}

	// subscribe to BTCUSD, XLMUSD, ADACAD book
	err = c.SubscribeBook([]string{ws.QTUMCAD, ws.REPBTC, ws.XLMUSD}, ws.Depth100)
	if err != nil {
		log.Fatal(err)
	}

	for obj := range c.Listen() {
		switch obj.(type) {
		case error:
			log.Printf("channel closed: %s", obj)
		case ws.DataUpdate:
			data := obj.(ws.DataUpdate)
			log.Printf("MSG RECV: %#v", data)
		default:
		}
	}
}
