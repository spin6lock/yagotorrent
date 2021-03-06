package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
)

func read_int(index int, buffer string) (int, int) {
	index = index + 1 //skip beginning 'i'
	start := index
	for byte(buffer[index]) != 'e' {
		index = index + 1
	}
	end := index      //position of 'e'
	index = index + 1 //prepare read pointer for next one
	result_str := string(buffer[start:end])
	result, e := strconv.Atoi(result_str)
	if e != nil {
		fmt.Println("Position:", start)
		fmt.Println(e)
	}
	return index, result
}

func read_string(index int, buffer string) (int, string) {
	start := index //start position of len digit
	for byte(buffer[index]) != ':' {
		index = index + 1
	}
	end := index //position of ':'
	len_str := buffer[start:end]
	length, e := strconv.Atoi(len_str)
	if e != nil {
		fmt.Println("start:", start, "end:", end)
		fmt.Println(e)
	}
	string_start := end + 1             //skip ':'
	string_end := string_start + length //end of str plus 1
	return string_end, string(buffer[string_start:string_end])
}

func read_dict(index int, buffer string) (int, map[string]interface{}) {
	//d3:cow3:moo4:spam4:eggse
	mydict := make(map[string]interface{})
	index = index + 1 //skip beginning 'd'
	for byte(buffer[index]) != 'e' {
		var key interface{}
		var value interface{}
		var tmp_index interface{}
		tmp_index, key = switcher(index, buffer)
		index = tmp_index.(int)
		index, value = switcher(index, buffer)
		mydict[key.(string)] = value
	}
	index++ //skip ending 'e'
	return index, mydict
}

func read_list(index int, buffer string) (int, list.List) {
	//l4:spam4:eggse
	var mylist = list.New()
	index = index + 1 //skip beginning 'l'
	for byte(buffer[index]) != 'e' {
		var value interface{}
		var tmp_index interface{}
		tmp_index, value = switcher(index, buffer)
		index = tmp_index.(int)
		mylist.PushBack(value)
	}
	return index + 1, *mylist //skip ending 'e'
}

func switcher(index int, input_str string) (int, interface{}) {
	flag := byte(input_str[index])
	var result interface{}
	if flag == 'i' {
		index, result = read_int(index, input_str)
	} else if flag == 'd' {
		var tmp_index interface{}
		tmp_index, result = read_dict(index, input_str)
		index = tmp_index.(int)
	} else if flag == 'l' {
		var tmp_index interface{}
		tmp_index, result = read_list(index, input_str)
		index = tmp_index.(int)
	} else { //flag is string
		index, result = read_string(index, input_str)
	}
	return index, result
}

func my_print(content interface{}) {
	switch t := content.(type) {
	case string:
		fmt.Print(content)
	case int:
		fmt.Print(content)
	case map[string]interface{}:
		fmt.Println("{")
		for k, v := range t {
			my_print(k)
			fmt.Print(":")
			my_print(v)
			fmt.Println(",")
		}
		fmt.Println("}")
	case list.List:
		fmt.Println(content)
	}
}

func TestInputFromSTDIN() {
	var input_str string
	var buffer []byte = make([]byte, 4*1024)
	reader := bufio.NewReader(os.Stdin)
	input_len, e := reader.Read(buffer)
	for input_len != 0 || e == nil {
		input_str += string(buffer[0:input_len])
		input_len, e = reader.Read(buffer)
	}
	if e != io.EOF {
		fmt.Println(e)
		return
	}
	_, result := switcher(0, input_str)
	my_print(result)
}
