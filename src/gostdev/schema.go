package main

type Field struct {
	Description string
	Name        string
	Type        string
	Length      int
	Minval      float64
	Maxval      float64
}

func NewField() *Field {
	return &Field{}
}

func (f *Field) SetDescription(description string) *Field {
	f.Description = description
	return f
}

func (f *Field) SetName(name string) *Field {
	f.Name = name
	return f
}

func (f *Field) SetType(fieldType string) *Field {
	f.Type = fieldType
	return f
}

func (f *Field) SetLength(length int) *Field {
	f.Length = length
	return f
}

func (f *Field) SetMinval(minval float64) *Field {
	f.Minval = minval
	return f
}

func (f *Field) SetMaxval(maxval float64) *Field {
	f.Maxval = maxval
	return f
}

type Function struct {
	Description string
	Name        string
	Method      string
	Uri         string
	Args        []*Field
	Required    []string
}

func NewFunction() *Function {
	return &Function{}
}

func (f *Function) SetDescription(description string) *Function {
	f.Description = description
	return f
}

func (f *Function) SetName(name string) *Function {
	f.Name = name
	return f
}

func (f *Function) SetMethod(method string) *Function {
	f.Method = method
	return f
}

func (f *Function) SetUri(uri string) *Function {
	f.Uri = uri
	return f
}

func (f *Function) AddArg(arg *Field) *Function {
	f.Args = append(f.Args, arg)
	return f
}

type Entity struct {
	Description string
	Name        string
	Fields      []*Field
}

func NewEntity() *Entity {
	return &Entity{}
}

func (e *Entity) SetDescription(description string) *Entity {
	e.Description = description
	return e
}

func (e *Entity) SetName(name string) *Entity {
	e.Name = name
	return e
}

func (e *Entity) AddField(field *Field) *Entity {
	e.Fields = append(e.Fields, field)
	return e
}

type Schema struct {
	Entities  []*Entity
	Functions []*Function
}

func NewSchema() *Schema {
	return &Schema{}
}

func (s *Schema) AddEntity(entity *Entity) *Schema {
	s.Entities = append(s.Entities, entity)
	return s
}

func (s *Schema) AddFunction(function *Function) *Schema {
	s.Functions = append(s.Functions, function)
	return s
}
