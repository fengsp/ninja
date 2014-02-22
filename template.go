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
