package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	var isDir bool
	flag.BoolVar(&isDir, "d", false, "boolean of folder")
	flag.Parse()

	re, err := regexp.Compile(`(\\\W)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	if isDir {
		dirName := flag.Arg(0)
		files, err := ioutil.ReadDir(dirName)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, v := range files {
			fileName := dirName + "/" + v.Name()
			fmt.Println(fileName)
			if err := convert(fileName, re); err != nil {
				fmt.Println(err)
				return
			}
		}
		return
	}
	fileName := flag.Arg(0)
	if err := convert(fileName, re); err != nil {
		fmt.Println(err)
		return
	}

}

func convert(fileName string, re *regexp.Regexp) error {
	if md, err := os.ReadFile(fileName); err != nil {
		return err
	} else {
		res := re.ReplaceAll(md, []byte(`\$1`))
		return os.WriteFile(fileName, res, 0644)
	}
}
