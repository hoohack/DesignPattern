package interpreter

import "strings"
import "strconv"

type Stack struct {
	top *Element
	size int
}

type Element struct {
	value Expression // All types satisfy the empty interface, so we can store anything here.
	next *Element
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value Expression) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value Expression) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

type Calculator struct {
  statement string
  expression  Expression
  Symbol_arr [5]string
}

func inArray(arr [5]string, elem string) bool {
  for _, data := range arr {
    if elem == data {
      return true
    }
  }

  return false
}

func (this *Calculator) Init(statement string) {
  expression_stack := new(Stack)
  var Left Expression
  var Right Expression

  statement_arr := strings.Split(statement, " ")
  for idx := 0; idx < len(statement_arr); idx++ {
    elem := statement_arr[idx]
    if inArray(this.Symbol_arr, elem) {
      Left = expression_stack.Pop()
      value, _ := strconv.Atoi(statement_arr[idx+1])
      idx = idx + 1
      Right = NewValueExpression(value)
      if elem == "+" {
        expression_stack.Push(NewPlusExpression(Left, Right))
      } else if elem == "-" {
        expression_stack.Push(NewMinusExpression(Left, Right))
      } else if elem == "*" {
        expression_stack.Push(NewMulExpression(Left, Right))
      } else if elem == "/" {
        expression_stack.Push(NewDivExpression(Left, Right))
      } else {
        expression_stack.Push(NewModExpression(Left, Right))
      }
    } else {
      value2, _ := strconv.Atoi(elem)
      expression_stack.Push(NewValueExpression(value2))
    }
  }
  this.expression = expression_stack.Pop()
}

func (this *Calculator) Compute() int {
  return this.expression.Calculate()
}
