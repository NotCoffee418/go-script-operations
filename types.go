package goscriptoperations

// Operation is used to register a function with description and optional order index
type Operation struct {
	Description string
	Function    func()
	Index       float32
}
