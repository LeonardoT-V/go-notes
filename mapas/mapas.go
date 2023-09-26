package mapas

import "fmt"

func Mapas() {
	paises := make(map[string]string)

	paises["mexico"] = "D.F"
	paises["ecuador"] = "quito"

	fmt.Println(paises)
	fmt.Println(paises["ecuador"])

	campeonato := map[string]int{"Barcelona": 19, "R.Madrid": 12}
	fmt.Println(campeonato)
	fmt.Println(campeonato["ecuador"])

	for equipo, puntaje := range campeonato {
		fmt.Printf("equipo: %s\npuntaje: %d\n", equipo, puntaje)
	}
	delete(campeonato, "R.Madrid")
	fmt.Println(campeonato)
	puntaje, exist := campeonato["Barcelona"]
	fmt.Printf("Los puntos son %d y existe %t", puntaje, exist)
}
