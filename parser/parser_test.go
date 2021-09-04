package parser

import (
	"fmt"
	"testing"

	"github.com/salleaffaire/ynt/ast"
	"github.com/salleaffaire/ynt/lexer"
)

func TestNumberValue(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"5", 5},
		{"5.25", 5.25},
		{"0.763", 0.763},
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		document := p.ParseDocument()

		if len(document.Values) != 1 {
			t.Fatalf("document.Values does not contain 1 value. got=%d",
				len(document.Values))
		}

		obj := document.Values[0]
		numberValue, ok := obj.(*ast.NumberValue)
		if !ok {
			t.Fatalf("stmt not *ast.NumberValue. got=%T", obj)
		}
		if testLiteralValue(t, numberValue, tt.expectedValue) {
			return
		}
	}
}

func TestStringValue(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"\"Luc\"", "Luc"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		document := p.ParseDocument()

		if len(document.Values) != 1 {
			t.Fatalf("document.Values does not contain 1 value. got=%d",
				len(document.Values))
		}

		obj := document.Values[0]
		stringValue, ok := obj.(*ast.StringValue)
		if !ok {
			t.Fatalf("stmt not *ast.StringValue. got=%T", obj)
		}
		if testLiteralValue(t, stringValue, tt.expectedValue) {
			return
		}
	}
}

func TestArrayValue(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue string
	}{
		{"[\"Luc\", 0, 1, true]", "[Luc, 0, 1, true]"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		document := p.ParseDocument()

		if len(document.Values) != 1 {
			t.Fatalf("document.Values does not contain 1 value. got=%d",
				len(document.Values))
		}

		obj := document.Values[0]
		arrayValue, ok := obj.(*ast.ArrayValue)
		if !ok {
			t.Fatalf("stmt not *ast.ArrayValue. got=%T", obj)
		}
		actual := arrayValue.String()
		if actual != tt.expectedValue {
			t.Errorf("expected=%q, got=%q", tt.expectedValue, actual)
		}
	}
}

func TestObjectValue(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue string
	}{
		{"{\"Key\" : 0}", "{\"Key\":0}"},
		{"{\"key\":0, \"Array\":[1,2,\"Luc\",true]}", "{\"key\":0, \"Array\":[1, 2, \"Luc\", true]}"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		document := p.ParseDocument()

		if len(document.Values) != 1 {
			t.Fatalf("document.Values does not contain 1 value. got=%d",
				len(document.Values))
		}

		obj := document.Values[0]
		objectValue, ok := obj.(*ast.ObjectValue)
		if !ok {
			t.Fatalf("stmt not *ast.ObjectValue. got=%T", obj)
		}
		actual := objectValue.String()

		if actual != tt.expectedValue {
			t.Errorf("expected=%q, got=%q", tt.expectedValue, actual)
		}
	}
}

func testNumberValue(t *testing.T, no ast.Value, value float64) bool {
	num, ok := no.(*ast.NumberValue)
	if !ok {
		t.Errorf("il not *ast.NumberValuet. got=%T", no)
		return false
	}
	if num.Value != value {
		t.Errorf("num.Value not %f. got=%f", value, num.Value)
		return false
	}
	// if num.TokenLiteral() != fmt.Sprintf("%f", value) {
	// 	t.Errorf("integ.TokenLiteral not %f. got=%s", value,
	// 		num.TokenLiteral())
	// 	return false
	// }
	return true
}

func testStringValue(t *testing.T, so ast.Value, value string) bool {
	num, ok := so.(*ast.StringValue)
	if !ok {
		t.Errorf("il not *ast.StringValue. got=%T", so)
		return false
	}
	if num.Value != value {
		t.Errorf("so.Value not %s. got=%s", value, num.Value)
		return false
	}
	// if num.TokenLiteral() != fmt.Sprintf("%f", value) {
	// 	t.Errorf("integ.TokenLiteral not %f. got=%s", value,
	// 		num.TokenLiteral())
	// 	return false
	// }
	return true
}

func testBooleanValue(t *testing.T, so ast.Value, value bool) bool {
	num, ok := so.(*ast.BooleanValue)
	if !ok {
		t.Errorf("il not *ast.BooleanValue. got=%T", so)
		return false
	}
	if num.Value != value {
		t.Errorf("so.Value not %v. got=%v", value, num.Value)
		return false
	}
	// if num.TokenLiteral() != fmt.Sprintf("%f", value) {
	// 	t.Errorf("integ.TokenLiteral not %f. got=%s", value,
	// 		num.TokenLiteral())
	// 	return false
	// }
	return true
}

func testLiteralValue(
	t *testing.T,
	obj ast.Value,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int:
		return testNumberValue(t, obj, float64(v))
	case int64:
		return testNumberValue(t, obj, float64(v))
	case float64:
		return testNumberValue(t, obj, v)
	case string:
		return testStringValue(t, obj, v)
	case bool:
		return testBooleanValue(t, obj, v)
	default:
		fmt.Printf("%T", v)
	}
	t.Errorf("type of exp not handled. got=%T", obj)
	return false
}
