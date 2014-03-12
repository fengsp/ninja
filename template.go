package ninja

import "strings"

type Template struct {
	source string
}

func (t *Template) Render(name string) string {
	return strings.Replace(t.source, "name", name, -1)
}

func (t *Template) Load(source string) {
	t.source = source
}

// return code object
func (t *Template) compile() {
	//source := t.parse(source)
	//source = t.generate(source)
	//code := "bytecode for Go"
	//return code
}

// parse
func (t *Template) parse(source string) {
	//return NewParser(source).parse()
}

// generate source code
func (t *Template) generate() {
}
