// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package dbrepo

import (
	"github.com/google/uuid"
	"github.com/hrz8/goatsapp/internal/repo/dto"
	"github.com/jackc/pgx/v5/pgtype"
)

type Projects struct {
	ID          uuid.UUID           `db:"id" json:"id"`
	Name        string              `db:"name" json:"name"`
	Alias       string              `db:"alias" json:"alias"`
	Description *string             `db:"description" json:"description"`
	Settings    dto.ProjectSettings `db:"settings" json:"settings"`
	CreatedAt   pgtype.Timestamptz  `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamptz  `db:"updated_at" json:"updated_at"`
}
