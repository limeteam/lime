package book

import (
	"lime/pkg/api/model"
	"io/ioutil"
	"github.com/go-yaml/yaml"
)

//写入文件
func WriteYamlFile(filename string,chapter *model.Store) error {
	b, err := yaml.Marshal(chapter)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename, b, 0775); err != nil {
		return err
	}
	return nil
}