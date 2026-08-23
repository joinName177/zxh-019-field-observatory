package observatory

type Clipboard struct {
	layer Layer
	has   bool
}

func (c *Clipboard) Copy(l Layer) { c.layer = CloneLayer(l); c.has = true }
func (c *Clipboard) Clear()       { c.layer = Layer{}; c.has = false }
func (c *Clipboard) Paste() (Layer, bool) {
	if !c.has {
		return Layer{}, false
	}
	return CloneLayer(c.layer), true
}
func (c *Clipboard) Has() bool { return c.has }
