package giary

import (
	"io/ioutil"
	"os"
)

const gitignore = `
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out
.idea
.vscode

unlock
`

func checkGitIgnore() {
	_, err := os.Stat(".gitignore")
	if os.IsNotExist(err) {
		ioutil.WriteFile(".gitignore", []byte(gitignore), 0755)
	}
}
