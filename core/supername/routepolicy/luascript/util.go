package luascript

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"

	"github.com/ironzhang/superlib/superutil/supermodel"
)

func makeMapTable(ls *lua.LState, m map[string]string) *lua.LTable {
	lt := ls.NewTable()
	for key, value := range m {
		lt.RawSetString(key, lua.LString(value))
	}
	return lt
}

func makeClusterTable(ls *lua.LState, c supermodel.Cluster) *lua.LTable {
	lt := ls.NewTable()
	lt.RawSetString("Name", lua.LString(c.Name))
	lt.RawSetString("Labels", makeMapTable(ls, c.Labels))
	lt.RawSetString("EndpointNum", lua.LNumber(len(c.Endpoints)))
	return lt
}

func makeClusterSliceTable(ls *lua.LState, clusters []supermodel.Cluster) *lua.LTable {
	lt := ls.NewTable()
	for _, cluster := range clusters {
		lt.Append(makeClusterTable(ls, cluster))
	}
	return lt
}

func makeDestinations(lt *lua.LTable) ([]supermodel.Destination, error) {
	var dests []supermodel.Destination
	for i := 1; i <= lt.Len(); i++ {
		value := lt.RawGetInt(i)
		table, ok := value.(*lua.LTable)
		if !ok {
			return nil, fmt.Errorf("invalid return value, the %dth element is not a table", i)
		}
		if table.Len() != 2 {
			return nil, fmt.Errorf("invalid return value, the %dth table's len is not 2", i)
		}

		dest := supermodel.Destination{
			Cluster: lua.LVAsString(table.RawGetInt(1)),
			Percent: float64(lua.LVAsNumber(table.RawGetInt(2))),
		}
		dests = append(dests, dest)
	}
	return dests, nil
}
