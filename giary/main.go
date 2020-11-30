package giary

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func Run(rootDirName string) {
	ProjectDirname = rootDirName
	check()
	switch os.Args[len(os.Args)-1] {
	case "lock":
		fmt.Print("请输入密码:> ")
		password1 := InputPassword()
		fmt.Print("请再次输入密码:> ")
		password2 := InputPassword()
		if !bytes.Equal(password1, password2) {
			log.Fatalln("两次密码不相同")
		}
		client := NewClient(password1)
		EncryptAll(client)
	case "unlock":
		fmt.Print("请输入密码:> ")
		client := NewClient(InputPassword())
		DecryptAll(client)
	default:
		fmt.Println("请指定正确的参数。说明：\n  lock: 加密unlock文件夹中的所有.md文件；\n  unlock: 解密locked文件中的所有文件")
	}
}
