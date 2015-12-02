package interpreter

import "strings"
import "strconv"
import "container/heap"

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
  expression_stack := heap.Init(Expression)
  var Left Expression
  var Right Expression

  statement_arr := strings.Split(statement, " ")
  for idx, elem := range statement_arr {
    if inArray(this.Symbol_arr, elem) {
      Left = heap.Pop(expression_stack)
      value, _ := strconv.Atoi(statement_arr[idx+1])
      Right = NewValueExpression(value)
      if elem == "+" {
        heap.Push(expression_stack, NewPlusExpression(Left, Right))
      } else if elem == "-" {
        heap.Push(expression_stack, NewMinusExpression(Left, Right))
      } else if elem == "*" {
        heap.Push(expression_stack, NewMulExpression(Left, Right))
      } else if elem == "/" {
        heap.Push(expression_stack, NewDivExpression(Left, Right))
      } else {
        heap.Push(expression_stack, NewModExpression(Left, Right))
      }
    } else {
      value, _ := strconv.Atoi(elem)
      heap.Push(expression_stack, NewValueExpression(value))
    }
  }
  this.expression = heap.Pop(expression_stack)
}

func (this *Calculator) Compute() int {
  return this.expression.Calculate()
}
