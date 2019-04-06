package main

import (
	"strings"
	"fmt"
	"os"
	"bufio"
)

// FindReplaceFile takes 
// src as source file
// old as old pattern
// new as new pattern
// and retun :
// occ as number of occurences
// lines as slice of line nb where old was found
// err as error 
func FindReplaceFile(src, old, new string) (occ int, lines []int, err error){
	linenb := 1
	file, err := os.Open(src)
	if err != nil{
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line:= scanner.Text()
		found, res, lineocc :=  ProcessLine(line, old, new)
		if found {
			lines = append(lines, linenb)
			occ += lineocc
		}
		fmt.Println(res)
		linenb++
	}
	return
}

// ProcessLine takes 
// line as line to be processed
// old as old pattern
// new as new pattern
// and return :
// found too indicate if old was found in line at least one time
// res as result of replacement
// occ as number of occurences
func ProcessLine(line, old, new string)(found bool, res string, occ int){
	found = strings.Contains(line, old)
	res = strings.Replace(line, old, new, -1)
	occ = strings.Count(line, old)
	return
}

func main(){
	occ, lines, err := FindReplaceFile("input.tx", "test", "toto")
	if err != nil {
			fmt.Println(err)
			return
	}
	fmt.Print(occ, lines)
	return
}
