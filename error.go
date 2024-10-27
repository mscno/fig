package fig

import (
	"errors"
	"sort"
	"strings"
)

// ErrFileNotFound is returned as a wrapped error by `Load` when the config file is
// not found in the given search dirs.
var ErrFileNotFound = errors.New("file not found")

// fieldErrors collects errors for fields of config struct.
type fieldErrors map[string]error

// growFactor is the factor by which the string builder grows when it needs more space.
const growFactor = 10

// Error formats all fields errors into a single string.
func (fe fieldErrors) Error() string {
	keys := make([]string, 0, len(fe))
	for key := range fe {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var sb strings.Builder
	sb.Grow(len(keys) * growFactor)

	for _, key := range keys {
		sb.WriteString(key)
		sb.WriteString(": ")
		sb.WriteString(fe[key].Error())
		sb.WriteString(", ")
	}

	return strings.TrimSuffix(sb.String(), ", ")
}
