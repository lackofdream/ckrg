package ckrg

type Result struct {
	Name       string
	Ok         bool
	Region     string
	Restricted bool
}

type Plugin struct {
	ID    string
	Name  string
	Check func() (Result, error)
}

var (
	Plugins map[string]Plugin
)

func RegisterPlugin(id string, name string, check func() (Result, error)) {
	if id == "" {
		panic("plugin id missing")
	}
	if check == nil {
		panic("missing Check()")
	}
	Plugins[id] = Plugin{
		ID:    id,
		Name:  name,
		Check: check,
	}
}

func init() {
	Plugins = make(map[string]Plugin, 0)
}
