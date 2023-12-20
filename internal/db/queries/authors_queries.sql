-- name: InsertAuthor :one
INSERT INTO
    authors (first_name, last_name)
VALUES
    ($1, $2) RETURNING *;

-- name: UpdateAuthor :exec
UPDATE
    authors
SET
    first_name = $1,
    last_name = $2
WHERE
    id = $3;

-- name: GetAuthors :many
SELECT
    *
from
    authors;

-- name: DeleteDeleteAuthor :exec
DELETE FROM
    authors
WHERE
    id = $1;