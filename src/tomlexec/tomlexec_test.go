package tomlexec

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExec(t *testing.T) {
	apath, _ := filepath.Abs("../")
	os.Chdir(apath)
	Exec()
}
