package modifyfile

import (
	"io/ioutil"
	"fmt"
	"strings"
)

func ReadFileParameter(filename string, parameter string) []byte {
	var p []byte

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	//location means the end of "address 1"
	l := FindParameterLocation(b, parameter)
	if l == -1 {
		p = append(p, ' ')
		return p
	}

	for i := l; b[i] != '\n' ; i++ {
		p = append(p, b[i])
	}

	return p
}

func ChangeFileParameter(filename string, parameter string, value string) {

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	//location means the end of "address 1"
	location := FindParameterLocation(b, parameter)

	len := FindParameterLength(b, location)

	b1 := append([]byte{}, b[:location]...)
	b2 := []byte(value)
	b3 := append([]byte{}, b[location + len + 1 :]...)//this b[l] can lead to bug !

	d := append(b1, b2...)
	d = append(d, b3...)

	ioutil.WriteFile(filename, d, 0777)
}

func FindParameterLocation(b []byte, parameter string) int {
	str := string(b)
	if strings.Index(str, parameter) == -1 {
		return -1
	}
	return strings.Index(str, parameter) + len(parameter) + 1
}

func FindParameterLength(b []byte, start int) int {
	i := start
	for b[i] != '\n' {
		i++
	}
	return i - start - 1
}
