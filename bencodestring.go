package main

import (
	"fmt"
	"strconv"
	"container/list"
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
	fmt.Println("length: ", length)
	fmt.Println("start: ", end+1)
	fmt.Println(string(buffer[end+1:end+1+length]))
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
		//fmt.Println("from read_dict", key, buffer[index:])
		index, value = switcher(index, buffer)
		mydict[key.(string)] = value
		//fmt.Println("key:value", mydict)
	}
	return index, mydict
}

func read_list(index int, buffer string) (int, list.List) {
	//l4:spam4:eggse
	var mylist = list.New()
	index = index + 1
	for byte(buffer[index]) != 'e' {
		var value interface{}
		var tmp_index interface{}
		tmp_index, value = switcher(index, buffer)
		index = tmp_index.(int)
		//fmt.Println("from read_list", value)
		mylist.PushBack(value)
	}
	//fmt.Println("mylist: ", mylist)
	return index, *mylist
}

func switcher(index int, input_str string) (int, interface{}) {
	flag := byte(input_str[index])
	var result interface{}
	if flag == 'i'{
		index, result = read_int(index, input_str)
		//fmt.Printf("%d\n", result)
	}	else if flag == 'd' {
		var tmp_index interface{}
		tmp_index, result = read_dict(index, input_str)
		index = tmp_index.(int)
	} else if flag == 'l' {
		var tmp_index interface{}
		tmp_index, result = read_list(index, input_str)
		index = tmp_index.(int)
	} else { //flag is string
		index, result = read_string(index, input_str)
		//fmt.Printf("%s\n", result)
	}
	index = index + 1
	//fmt.Printf("from switcher:", result)
	return index, result
}

func my_print(content interface{}){
	switch t := content.(type) {
	case string:
		fmt.Print(content)
	case int:
		fmt.Print(content)
	case map[string]interface{}:
		for k, v := range t{
			fmt.Println(k, " : ")
			my_print(v)
		}
	case list.List:
		fmt.Println(content)
	}
}

func main(){
	var input_str string
	fmt.Scanf("%s", &input_str)
	//length := len(input_str)
	index := 0
	index, result := switcher(index, input_str)
	my_print(result)
}
