package evaluator

import (
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"testing"
)

func TestEvalIntegerExpressinon(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 -10", 10},
		{"2 * 2 * 2", 8},
		{"4 + (2 + 3) * 5", 29},
		{"5 + 5 + 5 + 5 -10", 10},
		{"-3 + (3 + (10 * 2))", 20},
	}

	for _, tt := range tests {
		evaluated := execEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleaninon(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	}

	for _, tt := range tests {
		evaluated := execEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func execEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := execEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := execEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func TestReturnStatemewnts(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{"if(10 > 1) { if(10 > 1) { return 10; } return 1; }", 10},
	}

	for _, tt := range tests {
		evaluated := execEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mistmatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5",
			"type mistmatch: INTEGER + BOOLEAN",
		},
		{
			"-true;",
			"unknow operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknow operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknow operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknow operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
	}

	for _, tt := range tests {
		evaluated := execEval(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}

func TestLetStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	}

	for _, tt := range tests {
		evaluated := execEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2; };"

	evaluated := execEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Errorf("object is not Function. got=%T", evaluated)
	}
	if len(fn.Parameters) != 1 {
		t.Errorf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}
	if fn.Parameters[0].String() != "x" {
		t.Errorf("parameter is not 'x. got=%q", fn.Parameters[0])
	}
	expectedBody := "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Errorf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let identity = fn(x) { x; }; identity(5);", 5},
		{"let identity = fn(x) { return x; }; identity(5);", 5},
		{"let double = fn(x) { return x * 2; }; double(5);", 10},
		{"let add = fn(x, y) { return x + y; }; add(5, 5);", 10},
		{"let add = fn(x, y) { return x + y; }; add(5 + 5, add(5 , 5));", 20},
		{"fn(x) { x; }(5)", 5},
	}

	for _, tt := range tests {
		evaluated := execEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `let c = fn(x) { fn(y) { x + y } };
	          let cc = c(2);
			  cc(2) + cc(2)`
	evaluated := execEval(input)
	testIntegerObject(t, evaluated, 8)
}

func TestHigherOrderFunction(t *testing.T) {
	input := `let add = fn(a, b){ a + b };
	          let sub = fn(a, b){ a - b };
			  let apply = fn(a, b, func) { func(a,b) }
			  apply(10, 2, add) + apply(5, 3, sub)
			  `
	evaluated := execEval(input)
	testIntegerObject(t, evaluated, 14)
}
