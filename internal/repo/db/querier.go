// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package dbrepo

import (
	"context"
)

type Querier interface {
	GetProjects(ctx context.Context) ([]*Projects, error)
}

var _ Querier = (*Queries)(nil)
