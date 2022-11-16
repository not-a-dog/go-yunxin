package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

const headerFormat = `package %s

import "context"

//go:generate go run gen/main.go
`

var paramRegex = regexp.MustCompile(`type (\w+)Param struct {`)

const templateParamName = `UserUpdate`
const templateBody = `
func (UserUpdateParam) GetPath() string {
	return PathUserUpdate
}

func (c *Client) UserUpdate(ctx context.Context, param *UserUpdateParam) (*UserUpdateResponse, error) {
	return Request[UserUpdateResponse](c, ctx, param)
}
`

func dumpFile(gofile string, gopackage string, params []string) error {
	file, err := os.Create(gofile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, _ = file.WriteString(fmt.Sprintf(headerFormat, gopackage))
	sort.Strings(params)

	for _, param := range params {
		body := strings.ReplaceAll(templateBody, templateParamName, param)
		_, _ = file.WriteString(body)
	}

	return nil
}

func readParam(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	groups := paramRegex.FindAllStringSubmatch(string(data), -1)
	var params []string
	for _, g := range groups {
		params = append(params, g[1])
	}

	return params, nil
}

func readParams(cwd string, gofile string) ([]string, error) {
	es, err := os.ReadDir(cwd)
	if err != nil {
		return nil, err
	}

	var params []string
	for _, e := range es {
		if e.Type().IsRegular() {
			filename := e.Name()
			if filename == gofile || !strings.HasSuffix(filename, ".go") {
				continue
			}
			ps, err := readParam(e.Name())
			if err != nil {
				return nil, err
			}
			params = append(params, ps...)
		}
	}
	// params = []string{"AddUser", "ChatRoomCreate"}
	return params, nil
}

func main() {
	gofile := os.Getenv("GOFILE")
	gopackage := os.Getenv("GOPACKAGE")

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	params, err := readParams(cwd, gofile)
	if err != nil {
		panic(err)
	}

	err = dumpFile(gofile, gopackage, params)
	if err != nil {
		panic(err)
	}
}
