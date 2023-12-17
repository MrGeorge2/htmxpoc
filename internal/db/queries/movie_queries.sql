-- name: InsertMovie :one
INSERT INTO movies ( movie_name ) VALUES($1) RETURNING id;


-- name: UpdateMovie :exec
UPDATE movies SET movie_name = $1 AND id = $2;


-- name: GetMovies :many
SELECT * FROM movies;


-- name: DeleteMovie :exec
DELETE FROM movies WHERE id = $1;