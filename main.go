package main

import (
	"bufio"
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"os"
	"strings"
)

func main() {
	//fname := "sample1.png"
	name := os.Args[1]
	//var fname string
	in := bufio.NewReader(os.Stdin)
	//fname, err := in.ReadString('\n')
	//fname = strings.ReplaceAll(fname, string('"'), "")
	fname := name
	f, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(f)
	os.WriteFile("shangus.txt", []byte(f), 777)
	if strings.Contains(string(f), "AI generated image") == true {
		startplace := strings.Index(string(f), string([]byte("tEXtDescription"))) + len("tEXtDescription") + 1
		endplace := strings.Index(string(f), string(uint8(00))+string(uint8(00))+string(uint8(00))+string(uint8(16)))
		fmt.Println("Prompt :", string(f[startplace:endplace-4]))
		startplace = strings.Index(string(f), string([]byte("tEXtComment"))) + len("tEXtComment") + 1
		endplace = strings.Index(string(f), string(uint8(00))+string(uint8(01))+string(uint8(00))+string(uint8(00)))
		//fmt.Println(string(f[startplace : endplace-4]))
		json1 := string(f[startplace : endplace-4])
		fmt.Println("Steps :", gjson.Get(json1, "steps"))
		fmt.Println("sampler :", gjson.Get(json1, "sampler"))
		fmt.Println("seed :", gjson.Get(json1, "seed"))
		fmt.Println("strength :", gjson.Get(json1, "strength"))
		fmt.Println("noise :", gjson.Get(json1, "noise"))
		fmt.Println("scale :", gjson.Get(json1, "scale"))
		fmt.Println("uc :", gjson.Get(json1, "uc"))
		fmt.Println()
	} else {
		fmt.Println("정상적인 파일이 아니거나, 값을 가져올수 없습니다.")
	}
	in.ReadString('\n')
	//exec.Command("cmd", "/C", "pause")
}
