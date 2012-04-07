package main

import (
	"fmt"
	"strconv"
//	"strings"
)

func read_int(index int, buffer string) (int, int) {
	index = index + 1
	start := index
	for byte(buffer[index]) != 'e' {
		index = index + 1
	}
	end := index
	index = index + 1
	result_str := string(buffer[start:end])
	result, _ := strconv.Atoi(result_str)
	return index, result
}

func read_string(index int, buffer string) (int, string) {
	start := index
	for byte(buffer[index]) != ':' {
		index = index + 1
	}
	end := index 
	len_str := buffer[start:end]
	length, _ := strconv.Atoi(len_str)
	return index, string(buffer[end+1:end+1+length])
}

func main(){
	var input_str string
	fmt.Scanf("%s", &input_str)
	length := len(input_str)
	index := 0
	for index < length {
		flag := byte(input_str[index])
		if flag == 'i'{
			var result int
			index, result = read_int(index, input_str)
			fmt.Printf("%d\n", result)
		}	else if flag == 'd' {
			//var result 
		} else if flag == 'l' {
		} else { //flag is string
			var bencode_str string
			index, bencode_str = read_string(index, input_str)
			fmt.Printf("%s\n", bencode_str)
		}
		index = index + 1
	}
}
