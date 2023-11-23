package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron/v3"
)

func task() {
	fmt.Println(time.Now().String() + " - Start Task")
	time.Sleep(15 * time.Second)
	fmt.Println(time.Now().String() + " - End Task")
}

func listen() {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	fmt.Println(time.Now().String() + " - Closed")
}

func simpleCron() {
	fmt.Println(time.Now().String() + " - Start App - SimpleCron")
	c := cron.New()
	c.AddFunc("@every 1m10s", task)

	go c.Start()
}

func startImmediately() {
	fmt.Println(time.Now().String() + " - Start App - StartImmediately")
	c := cron.New()
	c.AddFunc("@every 1m10s", task)
	entry := c.Entries()
	entry[0].Job.Run()

	go c.Start()
}

func skipIfStillRunning() {
	fmt.Println(time.Now().String() + " - Start App - SkipIfStillRunning")

	c := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(cron.DefaultLogger),
	))
	c.AddFunc("@every 5s", task)
	entry := c.Entries()
	entry[0].Job.Run()

	go c.Start()
}

func main() {
	skipIfStillRunning()
	listen()
}
