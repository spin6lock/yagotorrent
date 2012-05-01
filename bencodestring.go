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

func read_dict(index int, buffer string) (int, map[string]interface{}) {
	//d3:cow3:moo4:spam4:eggse
	mydict := make(map[string]interface{})
	index = index + 1
	start := index
	for byte(buffer[index]) != 'e' {
		var key interface{}
		var value interface{}
		var tmp_index interface{}
		tmp_index, key = switcher(index, buffer)
		index = tmp_index.(int)
		index, value = switcher(index, buffer)
		mydict[key.(string)] = value
		index = index + 1
	}
	fmt.Println(mydict)
	return index, mydict
}

func read_list(index int, buffer string) (int, string) {
}

func switcher(index int, input_str string) (int, interface{}) {
	flag := byte(input_str[index])
	var result interface{}
	if flag == 'i'{
		index, result = read_int(index, input_str)
		fmt.Printf("%d\n", result)
	}	else if flag == 'd' {
		var result, tmp_index interface{}
		tmp_index, result = read_dict(index, input_str)
		index = tmp_index.(int)
		fmt.Printf(result.(string))
	} else if flag == 'l' {
		//flag is list
	} else { //flag is string
		var result string
		index, result = read_string(index, input_str)
		fmt.Printf("%s\n", result)
	}
	index = index + 1
	return index, result
}

func main(){
	var input_str string
	fmt.Scanf("%s", &input_str)
	length := len(input_str)
	index := 0
	for index < length {
		switcher(index, input_str)
	}
}
