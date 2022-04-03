package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

const (
	package_     = "package "
	_test        = "_test.go"
	func_        = "func "
	len_func_    = 5
	len_package_ = 8
)

var (
	golang_files      []string
	golang_test_files []string
)

func read_folders(folder_name string) {
	files, err := ioutil.ReadDir(folder_name)
	err_panic(err)
	for _, f := range files {
		dir := folder_name + "/" + f.Name()
		switch {
		case f.IsDir():
			read_folders(dir)
		case filepath.Ext(f.Name()) == ".go":
			if strings.Contains(f.Name(), _test) {
				golang_test_files = append(golang_test_files, dir)
			} else {
				golang_files = append(golang_files, dir)
			}
		}
	}
}

func file_lines(file_name string) []string {
	content, err := ioutil.ReadFile(file_name)
	err_panic(err)
	return strings.Split(string(content), "\n")
}

func err_panic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	read_folders(".")
	var all_func_name []string
	for _, a := range golang_files {
		all_func_name = append(all_func_name, return_all_func_name(a)...)
	}
	var all_file_lines []string
	for _, a := range golang_files {
		all_file_lines = append(all_file_lines, file_lines(a)...)
	}
big_loop:
	for _, a := range all_func_name {
		var number_of_time_used int
		for _, b := range all_file_lines {
			if strings.Contains(b, a+"(") || strings.Contains(b, a+"[") {
				number_of_time_used++
			}
			if number_of_time_used == 2 {
				continue big_loop
			}
		}
		fmt.Println("not used:", a)
	}
}

func return_all_func_name(path string) []string {
	var all_func_name []string
	for _, b := range file_lines(path) {
		var func_name string
		if len(b) > len_func_ && b[:len_func_] == func_ {
			for _, c := range b[len_func_:] {
				if c == '(' || c == '[' {
					all_func_name = append(all_func_name, func_name)
					break
				}
				func_name += string(c)
			}
		}
	}
	return all_func_name
}
