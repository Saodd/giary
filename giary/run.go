package giary

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

var ProjectDirname = "giary"
var logger = log.New(os.Stdout, "", log.Lshortfile)

func EncryptAll(client *Client) {
	check()
	for _, filepath := range RecurListMds("unlock") {
		plainText, err := ioutil.ReadFile(filepath)
		if err != nil {
			logger.Println(err)
			continue
		}
		cipherText := client.Seal(plainText)
		newPath := "locked" + filepath[6:]
		err = ioutil.WriteFile(newPath, cipherText, 0755)
		if err != nil {
			logger.Println(err)
			continue
		}
		os.Remove(filepath)
	}
}

func DecryptAll(client *Client) {
	check()
	for _, filepath := range RecurListMds("locked") {
		cipherText, err := ioutil.ReadFile(filepath)
		if err != nil {
			logger.Println(err)
			continue
		}
		plainText, err := client.Open(cipherText)
		if err != nil {
			logger.Println(err)
			continue
		}
		newPath := "unlock" + filepath[6:]
		err = ioutil.WriteFile(newPath, plainText, 0755)
		if err != nil {
			logger.Println(err)
			continue
		}
	}
}

func check() {
	workDir, err := os.Getwd()
	if err != nil {
		logger.Fatalln(err)
	}
	if path.Base(workDir) != ProjectDirname {
		logger.Fatalf("当前路径不是项目根目录(%s)！\n", ProjectDirname)
	}
	if err := os.MkdirAll("unlock", 0755); err != nil {
		logger.Fatalln(err)
	}
}

// RecurListMds 将递归遍历指定目录，返回所有 .md 文件的路径。
func RecurListMds(folder string) (mds []string) {
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			subFolder := path.Join(folder, file.Name())
			mds = append(mds, RecurListMds(subFolder)...)
		} else {
			mds = append(mds, path.Join(folder, file.Name()))
		}
	}
	return
}
