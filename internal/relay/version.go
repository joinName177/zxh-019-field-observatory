package observatory

type Version struct{ Major, Minor, Patch int }

func CurrentVersion() Version    { return Version{1, 0, 0} }
func (v Version) String() string { return fmtVersion(v) }
func fmtVersion(v Version) string {
	return string(rune('0'+v.Major)) + "." + string(rune('0'+v.Minor)) + "." + string(rune('0'+v.Patch))
}
func (v Version) Compatible(other Version) bool { return v.Major == other.Major }
