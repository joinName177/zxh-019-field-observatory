package observatory

import (
	"fmt"
	"strings"
)

func ValidateWorkspace(w Workspace) error {
	if strings.TrimSpace(w.ID) == "" {
		return fmt.Errorf("workspace id is required")
	}
	if w.Layers == nil {
		return fmt.Errorf("workspace layers are required")
	}
	return nil
}
func ValidateLayer(l Layer) error {
	if !l.Valid() {
		return ErrInvalidLayer
	}
	for i, p := range l.Points {
		if p.X != p.X || p.Y != p.Y {
			return fmt.Errorf("point %d is not finite", i)
		}
	}
	return nil
}
func NormalizeName(name string) string { return strings.Join(strings.Fields(name), " ") }
func TagsValid(tags []string) bool {
	seen := map[string]bool{}
	for _, tag := range tags {
		t := NormalizeName(tag)
		if t == "" || seen[t] {
			return false
		}
		seen[t] = true
	}
	return true
}
