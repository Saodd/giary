package giary

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func stty(start bool) error {
	r := "raw"
	if !start {
		r = "-raw"
	}

	rawMode := exec.Command("stty", r)
	rawMode.Stdin = os.Stdin
	err := rawMode.Run()
	if err != nil {
		return err
	}

	return rawMode.Wait()
}
func InputPassword() []byte {
	stty(true)
	defer stty(false)

	var password []rune
	inp := bufio.NewReader(os.Stdin)
	for {
		r, _, err := inp.ReadRune()
		if err != nil {
			logger.Fatalln(err)
		}

		if r == '\x03' { // ctrl+c
			os.Exit(0)
		} else if r == '\r' { // enter
			fmt.Print("\r\n")
			break
		} else if r == '\u007f' { // backspace
			if len(password) > 0 {
				password = password[:len(password)-1]
			}
			continue
		}

		password = append(password, r)
	}

	return []byte(string(password))
}
