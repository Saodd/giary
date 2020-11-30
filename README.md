# Giary

一款简单的本地加密/解密应用。

把你的数据加密后，传到Github（或者其他的Git平台）来进行版本管理吧！

## 使用方法

```shell
$ go mod init 你的项目根目录名字
$ go get github.com/Saodd/giary
```

```go
package main

import "github.com/Saodd/giary/giary"

func main() {
	giary.Run("你的项目根目录名字")
}
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

> 试试输入中文？

### 上传到Git

把你加密后的数据传到Git来进行分布式的数据管理！

你可以建立一个私有的 Github Repo ，也可以使用任意你喜欢的私有的（甚至公开的！）Git服务。

只要别忘了你的密码：）
