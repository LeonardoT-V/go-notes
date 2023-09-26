package modeluser

import (
	"fmt"
	"time"

	"github.com/leonardoT-V/go-notes/modelos"
)

func AltaUsuario() {
	us := new(modelos.User)
	us.AddUser(10, "pablo", time.Now(), true)

	fmt.Println(us)
}
