package main
/**
* 建造者模式将复杂对象的构造与其表示分离，以便相同的构造过程可以创建不同的表示。
*/
import "fmt"

type Speed float64
const (
	KPH Speed = 1
	MPH Speed = 1.60934
)

type Color string
const (
	BlueColor Color = "blue"
	GreenColor Color = "green"
	RedColor Color = "red"
)

type Wheels string
const (
	SportsWheels Wheels = "sports"
	SteelWheels ="steel"
)

type Builder interface {
	Paint(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}
type CarBuilder struct {
	color Color
	wheels Wheels
	topspeed Speed
}

func (cb *CarBuilder) Paint(c Color) Builder{
	cb.color = c
	return cb
}

func (cb *CarBuilder) Wheels(w Wheels) Builder{
	cb.wheels = w
	return cb
}

func (cb *CarBuilder) TopSpeed(ts Speed) Builder{
	cb.topspeed = ts
	return cb
}

func (cb *CarBuilder) Build() Interface {
	co := CarOperation{
		color: cb.color,
		wheels: cb.wheels,
		topspeed: cb.topspeed,
	}
	return &co
}

type Interface interface {
	Drive() error
	Stop() error
}

type CarOperation struct {
	color Color
	wheels Wheels
	topspeed Speed
}

func (co *CarOperation) Drive() error {
	fmt.Printf("A [%s] [%s] car driven by %d KMP\n", co.color, co.wheels, int64(co.topspeed))
	return nil
}

func (co *CarOperation) Stop() error {
	fmt.Printf("A [%s] [%s] car stopped\n", co.color, co.wheels)
	return nil
}

func NewBuilder() Builder {
	return &CarBuilder{}
}

func main() {
	co := NewBuilder().Paint(RedColor).Wheels(SportsWheels).TopSpeed(Speed(25)).Build()
	co.Drive()
	co.Stop()
}