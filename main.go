package main

import (
	"fmt"
	"net/http"
)


func sayHello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "i'm 木丁丁")
}

func main() {
	http.HandleFunc("/", sayHello)
	fmt.Print("服务启动成功")
	er := http.ListenAndServe("0.0.0.0:8890", nil)
	if er != nil {
		fmt.Println("服务启动失败，", er)
	}
}

