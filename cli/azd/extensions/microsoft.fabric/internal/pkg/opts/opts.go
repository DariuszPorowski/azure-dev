package opts

import "github.com/thediveo/enumflag/v2"

type RootOpts struct {
	Debug  bool
	Output Format
	Query  string
}

type Format enumflag.Flag

const (
	FormatJSON Format = iota
	FormatTable
	FormatEnvVars
	FormatNone
)

var Formats = map[Format][]string{
	FormatJSON:    {"json"},
	FormatTable:   {"table"},
	FormatEnvVars: {"dotenv"},
	FormatNone:    {"none"},
}

func GetFormats() []string {
	var modes []string
	for _, v := range Formats {
		modes = append(modes, v...)
	}

	return modes
}
