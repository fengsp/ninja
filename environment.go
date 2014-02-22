package ninja

type Environment struct {
    loader *FileSystemLoader
    cache map[string]*Template
}

func (e *Environment) loadTemplate(name string) *Template {
    template := e.cache[name]
    if  template != nil {
        return template
    }
    template = e.loader.Load(name)
    e.cache[name] = template
    return template
}

func (e *Environment) GetTemplate(name string) *Template {
    return e.loadTemplate(name)
}

func NewEnvironment(loader *FileSystemLoader) *Environment {
    cache := make(map[string]*Template)
    environment := &Environment{loader: loader, cache: cache}
    return environment
}
