package main

import (
	"github.com/kqbi/go-oxf/oxfTest/internal/test"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	t := new(test.TestS)
	t.Init()
	go t.StartBehavior()
	t.Send(&test.EvPoll{})
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Print("VIID exit")
			t.Destroy()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
