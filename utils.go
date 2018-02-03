package main

import (
	"net/http"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
)

func GetMap(filename string, addr interface{}) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, addr)

	if err != nil {
		return err
	}

	return nil
}

func GetHttpBuffer(url string) (*[]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return &bodyBytes, nil
	}
}

func GetSourceHttp(url string, addr interface{}) error {
	pBytes, err := GetHttpBuffer(url)
	if (err != nil) {
		return (err)
	}

	err = yaml.Unmarshal(*pBytes, addr)
	if err != nil {
		return err
	}

	return nil
}

func CopyHttpToFile(url, destFile string) error {
	pBytes, err := GetHttpBuffer(url)
	if err != nil {
		return err
	}
	err = CopyBufferContents(*pBytes, destFile)
	if err != nil {
		return err
	}
	return nil
}

func Merge(in1 *map[string]interface{}, in2 *map[string]interface{}) *map[string]interface{} {
	if in1 == nil || len(*in1) == 0 {
		return in2
	}

	if in2 == nil || len(*in2) == 0 {
		return in1
	}

	new := make(map[string]interface{})
	for k, v := range *in1 {
		new[k] = v
	}

	for k, v := range *in2 {
		new[k] = v
	}

	return &new
}

// copyBufferContents copies the contents of the buffer named srcBuff to the file
// named by destFile. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source buffer.
func CopyBufferContents(srcBuff []byte, destFile string) (err error) {

	out, err := os.Create(destFile)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = out.Write(srcBuff); err != nil {
		return
	}
	err = out.Sync()
	return
}