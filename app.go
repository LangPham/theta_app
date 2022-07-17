package main

import (
	"fmt"
	"sync"
	// "time"
)

func init() {
	sup := Suppervisor{
		wg:         new(sync.WaitGroup),
		isShutdown: make(chan bool),
	}
	Sup = &sup
}

var Sup *Suppervisor

type Actor interface {
	Start()
	Stop()
}

type Suppervisor struct {
	wg         *sync.WaitGroup
	childs     []Application
	isShutdown chan bool
}

type Application struct {
	name    string
	status  string
	typeApp string // alway, ono
	actor   Actor
}

func (app *Application) Start_link() {
	defer app.Shutdown()
	app.status = "running"
	fmt.Printf("CHAY APP: %+v", app)
	app.actor.Start()
}

func (app *Application) Shutdown() {
	app.status = "shutdown"
	fmt.Printf("Shutdown APP: %+v", app)
}

func (sup *Suppervisor) addChild(child Application) {
	sup.childs = append(sup.childs, child)
}

func (sup *Suppervisor) WaitingShutdown() {

	for {
		select {
		case close := <-sup.isShutdown:
			fmt.Println("hello", close)
			if close == true {
				sup.wg.Done()
			}
		}
	}
}

func (sup *Suppervisor) WriteShutdown(b bool) {
	sup.isShutdown <- b

}

func (sup *Suppervisor) Start_link() {
	sup.Start()
	fmt.Printf("Sup: %+v \n", sup)
	go sup.WaitingShutdown()

	fmt.Println("Cho shutdown")

	sup.wg.Wait()
}

func (sup *Suppervisor) Start() {

	fmt.Println("Start suppervisor")
	fmt.Printf("CHILDS: %+v \n", len(sup.childs))
	sup.wg.Add(1)
	for i := 0; i < len(sup.childs); i++ {

		fmt.Printf("===========%+v==========\n", i)
		i := i
		go func() {

			sup.childs[i].Start_link()
			// Do something
			fmt.Printf("===========trong goroutine %+v==========\n", i)

		}()
	}

}
