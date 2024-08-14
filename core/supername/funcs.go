package supername

func mergeRouteContext(src map[string]string, dst map[string]string) map[string]string {
	if len(src) <= 0 {
		return dst
	}
	if len(dst) <= 0 {
		return src
	}

	m := make(map[string]string, len(src)+len(dst))
	for k, v := range src {
		m[k] = v
	}
	for k, v := range dst {
		m[k] = v
	}
	return m
}
