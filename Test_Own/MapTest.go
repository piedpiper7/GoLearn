package main

import "fmt"

func main(){
	mapInOut()
}

func mapInOut(){
	var map1 map[string] string
	map1 = make(map[string] string)
	map1 ["He"] = "ZXY"
	map1["She"] = "GJN"
	map1["Will"] = "ever"
	map1["Who"] = "dont't know"

	for sex,j:= range map1{
		fmt.Println(sex,j)
	}

	delete(map1,"Who")

	fmt.Println("After Delete")

	for sex,j:= range map1{
		fmt.Println(sex,j)
	}

	sex,ok := map1["He"]
	if(ok){
		fmt.Println("me hhh",sex)
	} else {
		fmt.Println("dont't exists ME")
	}
}
