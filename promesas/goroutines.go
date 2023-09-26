package promesas

import (
	"fmt"
	"strings"
	"time"
)

func DelayShowName(n string, chan1 chan bool) {
	letras := strings.Split(n, "")

	for _, letras := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print(letras)
	}

	chan1 <- true

}
