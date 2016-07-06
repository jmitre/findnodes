package main
//import "fmt"

import(
	"bufio"
	"io"
	//"io/ioutil"
	"os"
	"strings"
	"fmt"
	"strconv"
)

func check(e error){
	if e != nil && e != io.EOF{
		panic(e)
	}
}


func main() {
	var s = new(region)

	readNodes(s);
	readLinks(s);

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
		seperated := strings.Split(line, " ")
		if line != "" {
			fmt.Println(seperated[1])
			cost, _ := strconv.Atoi(seperated[1])
			r.addnode(node{seperated[0], false, cost, []*node{}})
		}
	}
}

func readLinks(r *region){
	f, err := os.Open("NodeLinks.txt")
	check(err)

	reader := bufio.NewReader(f)
	var line string


	for err != io.EOF {
		line , err = reader.ReadString('\n')
		check(err)
		line = strings.TrimSpace(line)

		if line != "" {

			seperated := strings.Split(line, ":")
			linksToAdd := strings.Split(seperated[1], ",")

			for _,node := range linksToAdd {
				r.connect(seperated[0], node)
			}
		}
	}
}
