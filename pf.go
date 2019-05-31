package goutil

import "fmt"

func Pfv(content *string){
	fmt.Printf("%v\n",*content)
}

func PfT(content *string){
	fmt.Printf("%T\n",*content)
}
