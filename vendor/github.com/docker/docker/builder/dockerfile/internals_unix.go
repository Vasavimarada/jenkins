// +build !windows

package dockerfile

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/pkg/system"
)

// normaliseDest normalises the destination of a COPY/ADD command in a
// platform semantically consistent way.
func normaliseDest(cmdName, workingDir, requested string) (string, error) {
	dest := filepath.FromSlash(requested)
	endsInSlash := strings.HasSuffix(requested, string(os.PathSeparator))
	if !system.IsAbs(requested) {
		dest = filepath.Join(string(os.PathSeparator), filepath.FromSlash(workingDir), dest)
		// Make sure we preserve any trailing slash
		if endsInSlash {
			dest += string(os.PathSeparator)
		}
	}
	return dest, nil
}
