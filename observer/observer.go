package main

/**
* 观察者模式允许类型实例将事件“发布”给希望在特定事件发生时更新的其他类型实例（“观察者”）。
 */

import (
	"fmt"
	"time"
)

type Event struct {
	Data int64
}

type Observer interface {
	OnNotify(Event)
}

type Notifier interface {
	Register(Observer)
	Deregister(Observer)
	Notify(Event)
}

type EventObserver struct {
	id int
}

type EventNotifier struct {
	Observers map[Observer]struct{}
}

func (eo *EventObserver) OnNotify(e Event) {
	fmt.Printf("Observer-%d received: %d\n", eo.id, e.Data)
}

func (en *EventNotifier) Register(o Observer) {
	en.Observers[o] = struct{}{}
}

func (en *EventNotifier) Deregister(o Observer) {
	delete(en.Observers, o)
}

func (en *EventNotifier) Notify(e Event) {
	for o, _ := range en.Observers {
		o.OnNotify(e)
	}
}

func main() {
	n := EventNotifier{
		Observers: map[Observer]struct{}{},
	}
	n.Register(&EventObserver{id: 1})
	n.Register(&EventObserver{id: 2})
	n.Register(&EventObserver{id: 3})
	n.Register(&EventObserver{id: 999})

	n.Notify(Event{Data: time.Now().UnixNano()})
}
