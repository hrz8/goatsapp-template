// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: copyfrom.go

package dbrepo

import (
	"context"
)

// iteratorForCreateNewProjects implements pgx.CopyFromSource.
type iteratorForCreateNewProjects struct {
	rows                 []*CreateNewProjectsParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateNewProjects) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateNewProjects) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].Name,
		r.rows[0].Alias,
		r.rows[0].Description,
		r.rows[0].Settings,
	}, nil
}

func (r iteratorForCreateNewProjects) Err() error {
	return nil
}

func (q *Queries) CreateNewProjects(ctx context.Context, arg []*CreateNewProjectsParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"projects"}, []string{"name", "alias", "description", "settings"}, &iteratorForCreateNewProjects{rows: arg})
}
