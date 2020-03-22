package directory

import (
	"os"
	"path/filepath"
)

func RootDir() string {
	b, _ := os.Getwd()
	//d := path.Join(path.Dir(b))
	return filepath.Dir(b)
}
