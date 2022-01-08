## quickstart
```bash
tar -zxvf go1.17.6.linux-amd64.tar.gz  -C ~/tools/
echo "export PATH=\$PATH:\$HOME/tools/go/bin" >> ~/.bashrc
go version
go -h
```

file `main.go` content as below shown:
```go
package main
import "fmt"

func main()  {
	fmt.Println("hello go world!!!")
}
```

```bash
go mod init github.com/leson/go_poc

go mod tidy

go run cmd/main.go
```

> Go code is grouped into packages, and packages are grouped into modules


## Reference
- [Official Web Site](https://go.dev/)
- [doc of quickstart](https://go.dev/doc/tutorial/getting-started)
- [doc of create-module](https://go.dev/doc/tutorial/create-module)