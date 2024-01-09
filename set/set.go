package set

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](es ...T) Set[T] {
	s := make(Set[T])
	s.AddAll(es)
	return s
}

func Union[T comparable](s1 Set[T], s2 Set[T]) Set[T] {
	s := make(Set[T])
	s.Union(s1)
	s.Union(s2)
	return s
}

func (s Set[T]) Add(e T) {
	s[e] = struct{}{}
}

func (s Set[T]) AddAll(es []T) {
	for _, e := range es {
		s.Add(e)
	}
}

func (s Set[T]) Union(_s Set[T]) {
	for k, v := range _s {
		s[k] = v
	}
}

func (s Set[T]) Contains(e T) bool {
	_, ok := s[e]
	return ok
}

func (s Set[T]) Remove(e T) {
	delete(s, e)
}

func (s Set[T]) RemoveAll(es []T) {
	for _, e := range es {
		s.Remove(e)
	}
}
