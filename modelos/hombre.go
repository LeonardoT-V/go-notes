package modelos

type Hombre struct {
	edad      int
	altura    float32
	peso      float32
	caminando bool
	pensando  bool
	comiendo  bool
	vivo      bool
}

func (h *Hombre) Walking()       { h.caminando = true }
func (h *Hombre) Eating()        { h.comiendo = true }
func (h *Hombre) Thinking()      { h.pensando = true }
func (h *Hombre) Genrer() string { return "hombre" }
