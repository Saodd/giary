# Giary

一款简单的本地加密/解密应用。

把你的数据加密后，传到Github（或者其他的Git平台）来进行版本管理吧！

## 使用方法

```shell
$ go mod init 你的项目根目录名字
$ go get -u github.com/Saodd/giary
```

```go
package main

import "github.com/Saodd/giary/giary"

func main() {
	giary.Run("你的项目根目录名字")
}
```

```shell
$ git init && git add . && git commit -m "Init"
```

### 解密

```shell
$ go run main.go unlock
  请输入密码:> 
```

如果是clone的本项目代码，请输入密码`123`，即可将`./locked/password-123.md`文件解密到`./unlock/password-123.md`中。

### 加密

```shell
$ go run main.go lock  
请输入密码:> 
请再次输入密码:> 
```

`./unlock`中的所有文件都会被加密写入到`./lock`文件夹中，并且清空所有文件。

> 试试输入中文密码？

### 上传到Git

把你加密后的数据传到Git来进行分布式的数据管理！

你可以建立一个私有的 Github Repo ，也可以使用任意你喜欢的私有的Git服务。

只要别忘了你的密码：）

## 友情提醒

数据珍贵。请重视个人隐私安全。

- 本地加密是有可能被暴力破解的。
- 请将本工具视为一道额外的保险，附加在你原先已经觉得比较可靠的加密流程中，而不是代替。
