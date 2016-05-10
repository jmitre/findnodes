package main
//import "fmt"

import(
	"fmt"
	"bufio"
	//"io"
	//"io/ioutil"
	"os"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}
	

func main() {
	f, err := os.Open("NodeList.txt")
	check(err)
	
	reader := bufio.NewReader(f)
	
	line , err := reader.ReadString('\n')//Dosen't stop when it hits a line break
	
	fmt.Print(line)


	var s = new(region)
	s.addnode(node{"Velia", true})
	s.addnode(node{"Heidel", true})
	//s.Display()
}


