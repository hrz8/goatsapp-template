-- name: GetProjects :many
SELECT * FROM projects WHERE 1 = 1 ORDER BY created_at ASC;

-- name: CreateNewProjects :copyfrom
INSERT INTO
    projects (
        name, alias, description, settings
    )
VALUES ($1, $2, $3, $4);