package pokemon

type Pokemon struct {
	Name string
	Value string
}

func (p Pokemon) GetName() string {
	return p.Name
}

func (p Pokemon) GetValue() string {
	return p.Value
}
