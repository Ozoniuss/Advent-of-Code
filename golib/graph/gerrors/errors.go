package gerrors

import "errors"

var (
	ErrVertexAlreadyExists = errors.New("vertex already exists")
	ErrVertexNotExists     = errors.New("vertex does not exist")
	ErrEdgeAlreadyExists   = errors.New("edge already exists")
)
