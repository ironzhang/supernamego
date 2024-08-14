package selection

import "github.com/ironzhang/superlib/superutil/supermodel"

func getValues(s Store, t supermodel.Token) []string {
	switch t.Type {
	case supermodel.Table:
		return s.GetValues(t.Table, t.Key)
	case supermodel.Const:
		return t.Consts
	default:
		return nil
	}
}

func equals(lefts, rights []string) bool {
	if len(lefts) != len(rights) {
		return false
	}
	if len(lefts) == 1 {
		return lefts[0] == rights[0]
	}

	m := make(map[string]struct{}, len(lefts))
	for _, v := range lefts {
		m[v] = struct{}{}
	}
	for _, v := range rights {
		if _, ok := m[v]; !ok {
			return false
		}
	}
	return true
}

func belongs(lefts, rights []string) bool {
	if len(lefts) == 1 && len(rights) == 1 {
		return lefts[0] == rights[0]
	}

	m := make(map[string]struct{}, len(rights))
	for _, v := range rights {
		m[v] = struct{}{}
	}
	for _, v := range lefts {
		if _, ok := m[v]; !ok {
			return false
		}
	}
	return true
}

func matches(s Store, r supermodel.Requirement) bool {
	lefts := getValues(s, r.Left)
	rights := getValues(s, r.Right)

	var result bool
	switch r.Operator {
	case supermodel.Equals:
		result = equals(lefts, rights)
	case supermodel.Belongs:
		result = belongs(lefts, rights)
	}
	if r.Not {
		result = !result
	}
	return result
}
