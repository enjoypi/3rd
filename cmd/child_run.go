package cmd

// config for sub
type cfgChild struct {
	Config config
	Child  child
}

type child struct {
	Bool bool
	Str  string
}

type config struct {
	File string
}
