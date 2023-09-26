package deferpanic

import "log"

func PanicExample() {
	// se debe llevar un recover con un defer -> el defer siempre se ejecuta a pesar del panic
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error de panic \n %v", reco)
		}
		// si es llamado como funcion hace falta ()
	}()

	if a := 1; a == 1 {
		panic("Se encontro el valor 1")
	}
}
