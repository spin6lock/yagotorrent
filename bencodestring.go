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
	return index + length, string(buffer[end+1:end+1+length])
}

func read_dict(index int, buffer string) (int, map[string]interface{}) {
	//d3:cow3:moo4:spam4:eggse
	mydict := make(map[string]interface{})
	index = index + 1
	for byte(buffer[index]) != 'e' {
		var key interface{}
		var value interface{}
		var tmp_index interface{}
		tmp_index, key = switcher(index, buffer)
		index = tmp_index.(int)
		fmt.Println("from read_dict", key, buffer[index:])
		index, value = switcher(index, buffer)
		mydict[key.(string)] = value
		fmt.Println("key:value", mydict)
	}
	return index, mydict
}

func read_list(index int, buffer string) (int, string) {
	return 1, "hello"
}

func switcher(index int, input_str string) (int, interface{}) {
	flag := byte(input_str[index])
	var result interface{}
	if flag == 'i'{
		index, result = read_int(index, input_str)
		fmt.Printf("%d\n", result)
	}	else if flag == 'd' {
		var tmp_index interface{}
		tmp_index, result = read_dict(index, input_str)
		index = tmp_index.(int)
	} else if flag == 'l' {
		//flag is list
	} else { //flag is string
		index, result = read_string(index, input_str)
		fmt.Printf("%s\n", result)
	}
	index = index + 1
	fmt.Printf("from switcher:", result)
	return index, result
}

func main(){
	var input_str string
	fmt.Scanf("%s", &input_str)
	length := len(input_str)
	index := 0
	for index < length {
		index, _ = switcher(index, input_str)
	}
}
