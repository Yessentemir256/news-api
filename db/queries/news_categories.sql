-- name: AddNewsCategory :exec
INSERT INTO NewsCategories (NewsId, CategoryId) VALUES ($1, $2);

-- name: GetCategoriesByNewsId :many
SELECT CategoryId FROM NewsCategories WHERE NewsId = $1;