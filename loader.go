package ninja

import "io/ioutil"

type FileSystemLoader struct {
    path string
}

func (loader *FileSystemLoader) Load(name string) *Template {
    filePath := loader.path + name
    bytes, err := ioutil.ReadFile(filePath)
    source := string(bytes)
    if err != nil {
        // Log
    }
    template := &Template{source}
    return template
}

func NewFileSystemLoader(path string) *FileSystemLoader {
    loader := &FileSystemLoader{path: path}
    return loader
}
