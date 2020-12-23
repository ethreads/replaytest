package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (

)

func main()  {
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	select {
	case <-c:
		log.Println("exit")
		return
	}
}
