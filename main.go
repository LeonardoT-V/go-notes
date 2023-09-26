package main

import (
	"github.com/leonardoT-V/go-notes/promesas"
)

func main() {
	// go no espera a que termine la ejecucion de una promesa
	chan1 := make(chan bool)
	go promesas.DelayShowName("Leonardo", chan1)
	/*
		Con esto se puede cordinar para que todas las promesas se completen antes de terminar
		defer func ()  {
			<-chan1
		}() */
	// esto es un await
	/* res := <-chan1

	fmt.Println(res) */
	defer func() {
		<-chan1
	}()
}
