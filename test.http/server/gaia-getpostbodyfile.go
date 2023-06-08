package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func saybodyHandler(w http.ResponseWriter, r *http.Request) {
	s, _ := ioutil.ReadAll(r.Body) // 把body 内容读入字符串 s
	// getJson, _ := json.Marshal(s)
	fmt.Println(string(s))
	fmt.Println("==============")

	// write to file
	fileName := "./posted.json"
	writeFile(fileName, string(s))
}

func main() {

	http.HandleFunc("/", saybodyHandler)
	fmt.Println("server on :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func writeFile(fileName, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, _ = file.WriteString(content)

	_ = file.Close()
}
