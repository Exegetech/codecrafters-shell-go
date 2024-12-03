package main

type builtin int

const (
	echo builtin = iota
	exit
	type_
	pwd
	cd
)

func (b builtin) String() string {
	switch b {
	case echo:
		return "echo"
	case exit:
		return "exit"
	case type_:
		return "type"
	case pwd:
		return "pwd"
	case cd:
		return "cd"
	default:
		return "unknown"
	}
}

var builtins = map[string]bool{
	echo.String():  true,
	exit.String():  true,
	type_.String(): true,
	pwd.String():   true,
	cd.String():    true,
}
