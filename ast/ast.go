package ast

import (
	"bytes"
	"strings"

	"github.com/salleaffaire/ynt/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Value interface {
	Node
	valueNode()
}

type Document struct {
	Values []Value
}

func (p *Document) String() string {
	var out bytes.Buffer
	for _, s := range p.Values {
		if s != nil {
			out.WriteString(s.String())
			out.WriteString("\n")
		} else {
			out.WriteString("nil")
		}
	}
	return out.String()
}

func (p *Document) TokenLiteral() string {
	if len(p.Values) > 0 {
		return p.Values[0].TokenLiteral()
	} else {
		return ""
	}
}

type NumberValue struct {
	Token token.Token
	Value float64
}

func (nv *NumberValue) valueNode()           {}
func (nv *NumberValue) TokenLiteral() string { return nv.Token.Literal }
func (nv *NumberValue) String() string       { return nv.Token.Literal }

type StringValue struct {
	Token token.Token
	Value string
}

func (sv *StringValue) valueNode()           {}
func (sv *StringValue) TokenLiteral() string { return sv.Token.Literal }
func (sv *StringValue) String() string {
	var out bytes.Buffer

	out.WriteString("\"")
	out.WriteString(sv.Token.Literal)
	out.WriteString("\"")

	return out.String()
}

type BooleanValue struct {
	Token token.Token
	Value bool
}

func (bv *BooleanValue) valueNode()           {}
func (bv *BooleanValue) TokenLiteral() string { return bv.Token.Literal }
func (bv *BooleanValue) String() string       { return bv.Token.Literal }

type ArrayValue struct {
	Token  token.Token
	Values []Value
}

func (av *ArrayValue) valueNode()           {}
func (av *ArrayValue) TokenLiteral() string { return av.Token.Literal }
func (av *ArrayValue) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range av.Values {
		if e != nil {
			elements = append(elements, e.String())
		} else {
			elements = append(elements, "nil")
		}
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

type Attribute struct {
	Key string
	V   Value
}

func (a *Attribute) String() string {
	var out bytes.Buffer

	out.WriteString("\"")
	out.WriteString(a.Key)
	out.WriteString("\"")

	out.WriteString(":")
	out.WriteString(a.V.String())

	return out.String()
}

type ObjectValue struct {
	Token      token.Token
	Attributes []Attribute
}

func (ov *ObjectValue) valueNode()           {}
func (ov *ObjectValue) TokenLiteral() string { return ov.Token.Literal }
func (ov *ObjectValue) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ov.Attributes {
		elements = append(elements, e.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("}")
	return out.String()
}
