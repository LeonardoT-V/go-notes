package control

import (
	"fmt"
	"runtime"
)

func ShowDeviceRunning() {
	os := runtime.GOOS
	if os == "linux" {
		fmt.Println("you device is linux")
	} else {
		fmt.Println("you device is other")
	}
	fmt.Println(os)

	/*
		if osR := runtime.GOOS; osR == "linux" {
				fmt.Println("you device is linux")
			} else {
				fmt.Println("you device is other")
			}
			//fmt.Println(osR)
	*/
}
