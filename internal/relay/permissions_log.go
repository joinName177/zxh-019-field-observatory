package observatory

type Access struct {
	User, Workspace string
	Role            Role
}
type AccessList struct{ items []Access }

func NewAccessList() *AccessList { return &AccessList{items: []Access{}} }
func (a *AccessList) Grant(v Access) {
	for i, x := range a.items {
		if x.User == v.User && x.Workspace == v.Workspace {
			a.items[i] = v
			return
		}
	}
	a.items = append(a.items, v)
}
func (a *AccessList) Revoke(user, wid string) {
	out := a.items[:0]
	for _, x := range a.items {
		if x.User != user || x.Workspace != wid {
			out = append(out, x)
		}
	}
	a.items = out
}
func (a *AccessList) RoleOf(user, wid string) Role {
	for _, x := range a.items {
		if x.User == user && x.Workspace == wid {
			return x.Role
		}
	}
	return RoleViewer
}
func (a *AccessList) All(wid string) []Access {
	out := []Access{}
	for _, x := range a.items {
		if x.Workspace == wid {
			out = append(out, x)
		}
	}
	return out
}
