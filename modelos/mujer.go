package modelos

type Mujer struct {
	Hombre
	esMadre bool
}

func (m *Mujer) Walking()       { m.caminando = true }
func (m *Mujer) Eating()        { m.comiendo = true }
func (m *Mujer) Thinking()      { m.pensando = true }
func (m *Mujer) Genrer() string { return "mujer" }
