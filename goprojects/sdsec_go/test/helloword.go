// vi sdsec_go/test/hello_world.go
package main
import (
	"fmt"
	log "github.com/sirupsen/logrus" //引用一个Go的扩展包，并赋一个别名
)

func IndexHttp(){
	// 使用Go的扩展包
	log.WithFields(log.Fields{
		"age": 13,
		"name": "xiaohong",
		"sex": 0,
	}).Error("小红来了")
}
func main(){
	fmt.Printf("hello world!")
	IndexHttp()
}