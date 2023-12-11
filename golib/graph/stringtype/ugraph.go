package stringtype

import (
	"bufio"
	"maps"
	"os"
	"strconv"

	"aoc/golib/graph/gerrors"

	expmaps "golang.org/x/exp/maps"
)

type UnorderedGraph struct {
	vertices map[string]struct{}
	costs    map[[2]string]int
}

func NewUnorderedGraph() *UnorderedGraph {
	return &UnorderedGraph{
		vertices: make(map[string]struct{}, 8),
		costs:    make(map[[2]string]int, 8),
	}
}

func NewUnorderedGraphWithCapacity(vcap int, ecap int) *UnorderedGraph {
	return &UnorderedGraph{
		vertices: make(map[string]struct{}, vcap),
		costs:    make(map[[2]string]int, ecap),
	}
}

func (u *UnorderedGraph) Reset() {
	u.vertices = make(map[string]struct{}, 8)
	u.costs = make(map[[2]string]int, 8)
}

func (u *UnorderedGraph) VertexExists(v string) bool {
	_, ok := u.vertices[v]
	return ok
}

func (u *UnorderedGraph) VerticesExist(vertices ...string) bool {
	for _, v := range vertices {
		if _, ok := u.vertices[v]; !ok {
			return false
		}
	}
	return true
}

func (u *UnorderedGraph) GetVerticesMap() map[string]struct{} {
	return maps.Clone(u.vertices)
}

func (u *UnorderedGraph) GetVerticesUnderlyingMap() map[string]struct{} {
	return u.vertices
}

func (u *UnorderedGraph) GetVerticesSlice() []string {
	return expmaps.Keys(u.vertices)
}

func (u *UnorderedGraph) AddVertex(v string) error {
	if u.VertexExists(v) {
		return gerrors.ErrVertexAlreadyExists
	}
	u.vertices[v] = struct{}{}
	return nil
}

func (u *UnorderedGraph) MustAddVertex(v string) {
	u.vertices[v] = struct{}{}
}

func (u *UnorderedGraph) AddEdge(v1, v2 string, c int) error {
	if !u.VerticesExist(v1, v2) {
		return gerrors.ErrVertexNotExists
	}

	if v1 > v2 {
		v1, v2 = v2, v1
	}
	var edge [2]string = [2]string{v1, v2}
	if _, ok := u.costs[edge]; ok {
		return gerrors.ErrEdgeAlreadyExists
	}
	// Edge exists, no need to update neighbours
	u.costs[edge] = c
	return nil
}

func (u *UnorderedGraph) AddOrUpdateEdge(v1, v2 string, c int) error {
	if !u.VerticesExist(v1, v2) {
		return gerrors.ErrVertexNotExists
	}
	if v1 > v2 {
		v1, v2 = v2, v1
	}
	u.costs[[2]string{v1, v2}] = c
	return nil
}

func (u *UnorderedGraph) MustAddEdge(v1, v2 string, c int) {
	u.vertices[v1] = struct{}{}
	u.vertices[v2] = struct{}{}
	if v1 > v2 {
		v1, v2 = v2, v1
	}
	u.costs[[2]string{v1, v2}] = c
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

func (u *UnorderedGraph) ReadFromFile(path string) {
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
