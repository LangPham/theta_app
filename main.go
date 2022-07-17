package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main start ================")
	// Sup.addChild()
	go func() {
		time.Sleep(8 * time.Second)
		Sup.WriteShutdown(true)

		Web.Stop()
	}()

	app := Application{name: "Fiber", status: "stop", actor: Web}
	Sup.addChild(app)

	Sup.Start_link()

	fmt.Println("Main stop =================")

}
