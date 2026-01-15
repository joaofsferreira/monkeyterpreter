package object

import (
	"fmt"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
	RETURN_OBJ  = "RETURN"
	ERROR_OBJ   = "ERROR"
)

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

type Error struct {
	Message string
	// this object would hold a stack trace if our tokens had line and col numbers
	// since we're not doing that we only give the user a message with the found error
}

func (e *Error) Inspect() string  { return "[ERROR] " + e.Message }
func (e *Error) Type() ObjectType { return ERROR_OBJ }

type Null struct{}

func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

type Return struct {
	Value Object
}

func (r *Return) Inspect() string  { return r.Value.Inspect() }
func (r *Return) Type() ObjectType { return RETURN_OBJ }
