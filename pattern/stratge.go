package main

type Operator interface {
    Apple(int, int) int
}

type Operation struct {
    operator Operator
}

func (op *Operation)operate(left, right int)int {
    return op.operator.Apple(left, right)
}

type Addtion struct {}

func (add *Addtion) Apply(left, right int) int {
    return left + right
}

type Multiplication struct {}

func (mu *Multiplication)Apply(left, right int) int {
    return left*right
}

func CreateOpration(op Operator) Operation{
    return Operation{operator:op}
}
