package bytetype

import (
	"bufio"
	"os"
	"strconv"

	"aoc/golib/graph/gerrors"
)

type CappedString interface {
	[2]byte | [4]byte | [8]byte | [16]byte | [32]byte
}

func toByteArr[T CappedString](s string) T {
	var data T
	for i := 0; i < len(s); i++ {
		data[i] = s[i]
	}
	return data
}

type UnorderedGraph[T CappedString] struct {
	vertices map[T]struct{}
	costs    map[[2]T]int
}

func NewUnorderedGraph[T CappedString]() *UnorderedGraph[T] {
	return &UnorderedGraph[T]{
		vertices: make(map[T]struct{}, 8),
		costs:    make(map[[2]T]int, 8),
	}
}

func NewUnorderedGraphWithCapacity[T CappedString](vcap int, ecap int) *UnorderedGraph[T] {
	return &UnorderedGraph[T]{
		vertices: make(map[T]struct{}, vcap),
		costs:    make(map[[2]T]int, ecap),
	}
}

func (u *UnorderedGraph[T]) Reset() {
	u.vertices = make(map[T]struct{}, 8)
	u.costs = make(map[[2]T]int, 8)
}

func (u *UnorderedGraph[T]) VertexExists(v string) bool {
	_, ok := u.vertices[toByteArr[T](v)]
	return ok
}

func (u *UnorderedGraph[T]) VerticesExist(vertices ...string) bool {
	for _, v := range vertices {
		if _, ok := u.vertices[toByteArr[T](v)]; !ok {
			return false
		}
	}
	return true
}

// func (u *UnorderedGraph[T]) GetVerticesMap() map[T]struct{} {
// 	return maps.Clone(u.vertices)
// }

// func (u *UnorderedGraph[T]) GetVerticesUnderlyingMap() map[T]struct{} {
// 	return u.vertices
// }

// func (u *UnorderedGraph) GetVerticesSlice() []string {
// 	return expmaps.Keys(u.vertices)
// }

func (u *UnorderedGraph[T]) AddVertex(v string) error {
	if u.VertexExists(v) {
		return gerrors.ErrVertexAlreadyExists
	}
	u.vertices[toByteArr[T](v)] = struct{}{}
	return nil
}

func (u *UnorderedGraph[T]) MustAddVertex(v string) {
	u.vertices[toByteArr[T](v)] = struct{}{}
}

// func (u *UnorderedGraph[T]) AddEdge(v1, v2 string, c int) error {
// 	if !u.VerticesExist(v1, v2) {
// 		return ErrVertexNotExists
// 	}

// 	if v1 > v2 {
// 		v1, v2 = v2, v1
// 	}
// 	var edge [2]string = [2]string{v1, v2}
// 	if _, ok := u.costs[edge]; ok {
// 		return ErrEdgeAlreadyExists
// 	}
// 	// Edge exists, no need to update neighbours
// 	u.costs[edge] = c
// 	return nil
// }

// func (u *UnorderedGraph[T]) AddOrUpdateEdge(v1, v2 string, c int) error {
// 	if !u.VerticesExist(v1, v2) {
// 		return ErrVertexNotExists
// 	}
// 	if v1 > v2 {
// 		v1, v2 = v2, v1
// 	}
// 	u.costs[[2]string{v1, v2}] = c
// 	return nil
// }

func (u *UnorderedGraph[T]) MustAddEdge(v1, v2 string, c int) {
	if v1 > v2 {
		v1, v2 = v2, v1
	}
	v1arr := toByteArr[T](v1)
	v2arr := toByteArr[T](v2)
	u.vertices[v1arr] = struct{}{}
	u.vertices[v2arr] = struct{}{}
	u.costs[[2]T{v1arr, v2arr}] = c
}

// parse returns the ints from a line with no extra allocationss
func parse(line string) (string, string, int) {
	sep1 := 0
	sep2 := 0
	var v1, v2 string
	var c int

	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			if sep1 == 0 {
				sep1 = i
			} else {
				sep2 = i
				break
			}
		}
	}
	v1 = line[0:sep1]
	v2 = line[sep1+1 : sep2]
	c, _ = strconv.Atoi(line[sep2+1:])
	return v1, v2, c
}

func (u *UnorderedGraph[T]) ReadFromFile(path string) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	s.Scan() // read header
	for s.Scan() {
		line := s.Text()
		var v1, v2 string
		var cost int
		v1, v2, cost = parse(line)
		u.MustAddEdge(v1, v2, cost)
	}
}
