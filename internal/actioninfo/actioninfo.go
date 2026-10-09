package actioninfo

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
}
