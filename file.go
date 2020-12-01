package helper

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type File struct {
	path string
	data []string
}

func ConstructFile(path string) File {
	return File{
		path: path,
		data: nil,
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (f *File) GetDataAsStringArray() []string {
	if f.data == nil {
		f.readData()
	}

	return f.data
}

func (f *File) GetDataAsIntArray() []int {
	if f.data == nil {
		f.readData()
	}

	intArray := make([]int, 0)
	for _, v := range f.data {
		if v != "" {
			val, err := strconv.Atoi(v)
			check(err)
			intArray = append(intArray, val)
		}
	}

	return intArray
}

func (f *File) readData() {
	data, err := ioutil.ReadFile(f.path)
	if err != nil {
		panic(err)
	}

	f.data = deleteEmpty(strings.Split(string(data), "\n"))
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
