package main

import (
	"fmt"
)

type SetMonitor interface {
	InitConfig(...interface{}) error
	NextGroup()
	String() string
	Save() error
	View() MonitorSetData
}
type GroupMonitor interface {
	InitConfig(...interface{}) error
	View() MonitorGroupData
}
type UnitMonitor interface {
	InitConfig(...interface{}) error
	GetData()
	View() MonitorUnitData
	String() string
}

type (
	MonitorSetData   map[string]MonitorGroupData
	MonitorGroupData map[string]MonitorUnitData
	MonitorUnitData  map[string]interface{}
	MonitorSet       struct {
		Name string
		Data MonitorSetData
	}
	MonitorGroup struct {
		Name string
		Data MonitorGroupData
	}
	MonitorUnit struct {
		Name string
		Data MonitorUnitData
	}
)

type RedisMonitor struct {
	Name string
	Host string
	Port string
	Data string
}

func (r *RedisMonitor) String() string {
	return fmt.Sprintf("%s: %s", r.Name, r.Data)
}
func (r *RedisMonitor) GetData() {
	r.Data += "hahah# "
}
func (r *RedisMonitor) Save() error {
	return nil
}

func RunMonitor(m SetMonitor) {
	for i := 0; i < 3; i++ {
		m.GetData()
	}
	fmt.Println(m.String())
}

type Who interface {
	GetName(...interface{}) string
}
type Wang string

func (w *Wang) GetName(args ...interface{}) string {
	s := ""
	for i, v := range args {
		vwang, ok := args[i].(Wang)
		if ok {
			s += "Wang:" + fmt.Sprint(vwang) + " # "
			continue
		}
		switch v.(type) {
		case string:
			s += "String:" + fmt.Sprint(v) + " # "
		default:

		}
	}
	return fmt.Sprint(s)
}
func Hi(w Who) {
	var wang Wang = "WNANGWNG"
	fmt.Println(w.GetName(wang, "lalallalla"))
}

func main() {
	fmt.Println("...")
	var w *Wang
	Hi(w)
	// r := &RedisMonitor{Name: "Redis", Data: ""}
	// RunMonitor(r)
}
