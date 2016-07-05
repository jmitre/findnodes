package main
//import "fmt"

import(
	"bufio"
	"io"
	//"io/ioutil"
	"os"
	"strings"
	"fmt"
)

func check(e error){
	if e != nil && e != io.EOF{
		panic(e)
	}
}


func main() {
	var s = new(region)

	readNodes(s);

	fmt.Println(s)
}

func readNodes(r *region){
	f, err := os.Open("NodeList.txt")
	check(err)

	reader := bufio.NewReader(f)
	var line string

	for err != io.EOF {
		line , err = reader.ReadString('\n')
		check(err)

		line = strings.TrimSpace(line)

		if line != "" {
			r.addnode(node{line, false})
		}
	}
}
