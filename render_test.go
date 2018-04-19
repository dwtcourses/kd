package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"testing"
)

var emptymap map[string]string

func readfile(filepath string) string {

	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return string(dat)
}

func TestRender(t *testing.T) {
	test_data := make(map[string]string)
	test_data["MY_LIST"] = "one,two,three"
	test_data["FILE_PATH"] = "test/complex-file.pem"

	cases := []struct {
		name      string
		inputdata string
		inputvars map[string]string
		want      string
	}{
		{
			name:      "Check plain file is rendered",
			inputdata: readfile("test/deployment.yaml"),
			inputvars: emptymap,
			want:      readfile("test/deployment.yaml"),
		},
		{
			name:      "Check list variables are rendered",
			inputdata: readfile("test/list-prerendered.yaml"),
			inputvars: test_data,
			want:      readfile("test/list-rendered.yaml"),
		},
		{
			name:      "Check file function is rendered",
			inputdata: readfile("test/file-prerendered.yaml"),
			inputvars: test_data,
			want:      readfile("test/file-rendered.yaml"),
		},
		{
			name:      "Check contains function works as expected",
			inputdata: readfile("test/contains-prerendered.yaml"),
			inputvars: test_data,
			want:      readfile("test/contains-rendered.yaml"),
		},
		{
			name:      "Check hasPrefix function works as expected",
			inputdata: readfile("test/hasPrefix-prerendered.yaml"),
			inputvars: test_data,
			want:      readfile("test/hasPrefix-rendered.yaml"),
		},
		{
			name:      "Check hasSuffix function works as expected",
			inputdata: readfile("test/hasSuffix-prerendered.yaml"),
			inputvars: test_data,
			want:      readfile("test/hasSuffix-rendered.yaml"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Render(c.inputdata, c.inputvars)
			if err != nil {
				fmt.Println("Testing if folder doesnt exist")
			}
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got: %#v\nwant: %#v\n", got, c.want)
			}
		})
	}
}
