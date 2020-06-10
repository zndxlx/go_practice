package main

import "fmt"

type Event struct {
    Info string
}

type Observer interface {
    Receive(event Event)
}

type Notifier interface {
    Register(observer Observer)
    Remove(observer Observer)
    Notify(event Event)
}

//观察者
type InvestorObserver  struct {
    Name string
}

func (inverster *InvestorObserver) Receive(event Event){
    fmt.Printf("%s, 收到时间通知 %s\n", inverster.Name, event.Info)
}

//被观察者
type ShareNotifier struct {
    Price float64
    oblist []Observer
}

func (share *ShareNotifier) Register(observer Observer){
    share.oblist = append(share.oblist, observer)
}

func (share *ShareNotifier) Remove(observer Observer){
    if len(share.oblist) == 0 {
        return
    }
    for i, ob := range share.oblist {
        if ob == observer {
            share.oblist = append(share.oblist[:i], share.oblist[i+1:]...)
        }
    }
}

func (share *ShareNotifier)Notify(event Event){
    for _, ob := range share.oblist {
        ob.Receive(event)
    }
}

func NewEvent()Event {
    return Event{Info:"价格有变动"}
}

func NewInvertorObserver(name string) *InvestorObserver {
    return &InvestorObserver{Name:name}
}

func NewShareNotifier(price float64) *ShareNotifier {
    return &ShareNotifier{Price:price}
}