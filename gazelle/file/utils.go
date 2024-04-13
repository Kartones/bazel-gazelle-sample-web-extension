package file

import "path/filepath"

func FileBaseName(path string) string {
	base := filepath.Base(path)
	matches := fileBaseNamePattern.FindStringSubmatch(base)
	if len(matches) > 0 {
		return matches[1]
	}
	return base
}
