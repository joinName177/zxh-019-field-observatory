package observatory

import "errors"

var ErrPermission = errors.New("permission denied")

type Role string

const (
	RoleViewer Role = "viewer"
	RoleEditor Role = "editor"
	RoleOwner  Role = "owner"
)

func CanWrite(r Role) bool  { return r == RoleEditor || r == RoleOwner }
func CanDelete(r Role) bool { return r == RoleOwner }
func CheckWrite(r Role) error {
	if !CanWrite(r) {
		return ErrPermission
	}
	return nil
}
func CheckDelete(r Role) error {
	if !CanDelete(r) {
		return ErrPermission
	}
	return nil
}
