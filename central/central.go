package central

type Command interface {
	Payload() string
}

type CommandProcessor interface {
	Start() error
}
