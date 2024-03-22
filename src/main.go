package main

import (
	"fmt"
	"time"
)

func main() {
	//	todo: make config that is beeing read in by gonfig
	//	todo: call the weather and put data into struct
	//	todo: implement template file that gets that data and puts it on a website

	GetWeather()

	for { //ever
		GetWeather()
		time.Sleep(5 * time.Second)
	}
}

func count(x int) {
	for i := 0; i < x; i++ {
		fmt.Println(i)
		time.Sleep(300 * time.Millisecond)
	}
}
