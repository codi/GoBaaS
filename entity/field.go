package entity

type Field struct {
	Id      string
	Name    string
	Type    FieldType
	Comment string
}

type FieldType string

const (
	FieldTypeString = "string"
)
