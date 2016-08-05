package template

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Parser interface {
	parse(content string) string
}

type symbolStruct struct {
	strVars map[string]string
}

var symbolTable symbolStruct

func init() {
	symbolTable.strVars = make(map[string]string)
}

func rootPath() string {
	file, _ := exec.LookPath(os.Args[0])
	fullpath, _ := filepath.Abs(file)
	root := strings.Replace(fullpath, "/bin/goweb", "/src/yaoguais/goweb", 1)
	return root
}

func fileContent(filepath string) (string, error) {
	fp, err := os.Open(filepath)
	if err != nil {
		log.Println("file ", filepath, " not found")
		return "", err
	}
	defer fp.Close()
	fd, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Println("read file ", filepath, " content failed")
		return "", err
	}
	return string(fd), nil
}

func strVarsParse(content string) string {
	for k, v := range symbolTable.strVars {
		content = strings.Replace(content, "{{"+k+"}}", v, -1)
	}
	return content
}

func Fetch(filename string) string {
	root := rootPath()
	file := root + "/template/" + filename + ".tpl"
	content, err := fileContent(file)
	if err != nil {
		return ""
	}
	content = strVarsParse(content)
	return content
}

func Assign(key string, value string) {
	symbolTable.strVars[key] = value
}
