package delaunay

type triangleSet struct {
	triangles map[Triangle]struct{}
}

func (ts *triangleSet) Add(t Triangle) (set bool) {
	if _, ok := ts.triangles[t]; ok {
		return false
	}
	ts.triangles[t] = struct{}{}
	return true
}

func (ts *triangleSet) Remove(t Triangle) (deleted bool) {
	if _, ok := ts.triangles[t]; ok {
		delete(ts.triangles, t)
		return true
	}
	return false
}

func (ts *triangleSet) ToSlice() []Triangle {
	var res = make([]Triangle, 0, len(ts.triangles))
	for t := range ts.triangles {
		res = append(res, t)
	}
	return res
}
