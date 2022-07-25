package main

import (
	"fmt"
        "os"
	"encoding/binary"
	radio "github.com/DatanoiseTV/rf24-go"
)

func main() {
        uid := os.Getuid()
        if uid != 0 {
            fmt.Println("Must be running as root.")
            os.Exit(1)    
        } 

	var pipe uint64 = 0xF0F0F0F0E1
	r := radio.New(22, 0)
	defer r.Delete()
	r.Begin()
	r.SetRetries(15, 15)
	r.SetAutoAck(true)
	r.OpenReadingPipe(1, pipe)
	r.StartListening()
	r.PrintDetails()
	for {
		if r.Available() {
			data, _ := r.Read(4)
			fmt.Printf("data: %v\n", data)
			payload := binary.LittleEndian.Uint32(data)
			fmt.Printf("Received %v\n", payload)
		} else {
			//
		}
	}
}
