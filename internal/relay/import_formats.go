package observatory

import "strings"

type Format string

const (
	FormatJSON    Format = "json"
	FormatCSV     Format = "csv"
	FormatUnknown Format = "unknown"
)

func DetectFormat(name string) Format {
	switch strings.ToLower(strings.TrimPrefix(name[strings.LastIndex(name, "."):], ".")) {
	case "json":
		return FormatJSON
	case "csv":
		return FormatCSV
	default:
		return FormatUnknown
	}
}
func SupportedFormat(f Format) bool { return f == FormatJSON || f == FormatCSV }
func Extension(f Format) string {
	if f == FormatJSON {
		return ".json"
	}
	if f == FormatCSV {
		return ".csv"
	}
	return ""
}
func NormalizeFormat(f string) Format {
	return DetectFormat("x." + strings.TrimPrefix(strings.ToLower(strings.TrimSpace(f)), "."))
}
func FormatNames() []string { return []string{string(FormatJSON), string(FormatCSV)} }
