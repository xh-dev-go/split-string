package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/xh-dev-go/split-string/FlagCmd"
	"io"
	"os"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	delimiterParam := FlagCmd.NewStringParam("delimiter",";", "The delimiter for the spliting, should be char").
		Bind(flag.CommandLine)
	flag.Parse()

	if delimiterParam.Value()==""{
		panic("delimiter is needed")
	}
	for {
		bytes, error := reader.ReadString(delimiterParam.Value()[0])
		fmt.Println(bytes)
		if error!=nil && error==io.EOF{
			break
		} else if error!=nil {
			panic(error)
		}
	}
}
