package observatory

type Project struct {
	ID, Name, Owner string
	Workspaces      []string
}

func NewProject(id, name, owner string) Project {
	return Project{ID: id, Name: NormalizeName(name), Owner: owner, Workspaces: []string{}}
}
func (p *Project) AddWorkspace(id string) {
	for _, x := range p.Workspaces {
		if x == id {
			return
		}
	}
	p.Workspaces = append(p.Workspaces, id)
}
func (p *Project) RemoveWorkspace(id string) {
	out := p.Workspaces[:0]
	for _, x := range p.Workspaces {
		if x != id {
			out = append(out, x)
		}
	}
	p.Workspaces = out
}
func (p Project) HasWorkspace(id string) bool {
	for _, x := range p.Workspaces {
		if x == id {
			return true
		}
	}
	return false
}
func (p Project) Count() int { return len(p.Workspaces) }
