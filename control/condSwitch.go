package control

import (
	"fmt"
	"runtime"
)

func ShowDeviceRunningSwitch() {
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Is linux")
	case "windows":
		fmt.Println("Is windows")
	default:
		fmt.Printf("%s\n", os)
	}
}
