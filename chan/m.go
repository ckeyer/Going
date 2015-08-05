package main

import (
	"fmt"
	"sync"
	"time"
)

type Set struct {
	Name  string
	State chan int
	Data  interface{}
	sync.Mutex
}
type Group struct {
	Name  string
	State chan int
	Data  interface{}
	sync.Mutex
}
type Unit struct {
	Name  string
	State chan int
	Data  interface{}
	sync.Mutex
}

func (s *Set) GetName() string                         { return s.Name }
func (s *Set) GetGroups() []Grouper                    { return nil }
func (s *Set) GetData()                                {}
func (s *Set) AddGroup(name string, group interface{}) {}

func (g *Group) GetName() string                       { return g.Name }
func (g *Group) GetUnits() []Uniter                    { return nil }
func (g *Group) GetData()                              {}
func (g *Group) AddUnit(name string, unit interface{}) {}

func (u *Unit) GetName() string                       { return u.Name }
func (u *Unit) GetItem()                              {}
func (u *Unit) GetData()                              {}
func (u *Unit) AddItme(name string, item interface{}) {}

type Seter interface {
	GetName() string
	GetGroups() []Grouper
	GetData()
	AddGroup(string, interface{})
}
type Grouper interface {
	GetName() string
	GetUnits() []Uniter
	GetData()
	AddUnit(string, interface{})
}
type Uniter interface {
	GetName() string
	GetItem()
	GetData()
	AddItme(string, interface{})
}
type Monitor struct {
	Sets map[string]Set
}

func (m *Monitor) Run() {
	for _, s := range m.Sets {
		if v, ok := s.(Seter); ok {
			for _, g := range v.GetGroups() {
				for _, u := range g.GetUnits() {
					a := make(chan int64, 1)
					go func(chan<- int64) {
						start := time.Now().Unix()
						g.AddUnit(u.GetName(), u)
						runtime := time.Now().Unix() - start
						a <- runtime
					}(a)
					<-a
				}
				v.AddGroup(g.GetName(), g)
			}
		}
	}
}

func main() {
	fmt.Println("run")
}
