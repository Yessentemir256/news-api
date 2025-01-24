-- name: CreateNews :exec
INSERT INTO News (Title, Content) VALUES ($1, $2);

-- name: GetNewsById :one
SELECT * FROM News WHERE Id = $1;