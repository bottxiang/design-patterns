package strategy
/**
* 策略行为设计模式允许在运行时选择算法的行为。 它定义算法，封装它们，并交替使用它们。
*/

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(l, r int) int {
	return o.Operator.Apply(l, r)
}

type Addition struct {}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

type Multiplication struct {}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}