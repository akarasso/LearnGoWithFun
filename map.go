package main

import "fmt"

func main()  {
	var stringMap map[string]int
	var intMap map[int]int

	stringMap = make(map[string]int)
	intMap = make(map[int]int)
	stringMap["key1"] = 10
	stringMap["key2"] = 10
	stringMap["key3"] = 10
	fmt.Println(stringMap)

	intMap[1] = 10
	intMap[2] = 10
	intMap[3] = 10
	fmt.Println(intMap)
	delete(intMap, 1)
	fmt.Println(intMap)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	nameLi, okLi := elements["Li"]
	fmt.Println(nameLi, okLi)
	fmt.Println(elements["Un"])

	nameUn, okUn := elements["Un"]
	fmt.Println(nameUn, okUn)
}
