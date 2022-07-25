package main

import (
	"fmt"
        "os"
        "time"
	radio "github.com/DatanoiseTV/rf24-go"
)

func main() {
	fmt.Println("nRF24L01+ from GoLang")
        uid := os.Getuid()
        if uid != 0 {
            fmt.Println("Must be running as root.")
            os.Exit(1)    
        } 

	var pipe uint64 = 0x65646f4e32
	var readPipe uint64 = 0x65646f4e31
	r := radio.New(22, 0)
	defer r.Delete()
	r.Begin()
        r.SetChannel(110)
	r.SetRetries(3, 3)
	r.SetAutoAck(true)
	r.SetDataRate(radio.RATE_250KBPS)
        r.SetPALevel(radio.PA_LOW)
        r.EnableDynamicPayloads()
        r.OpenWritingPipe(pipe)
        r.OpenReadingPipe(1, readPipe)
	//r.OpenReadingPipe(1, pipe)
	r.StopListening()
	r.PrintDetails()


	for {
            resp := r.Write([]byte{0x42, 0x44}, 2)
            fmt.Println(resp) 

            time.Sleep(1000 * time.Millisecond)
	}
}
