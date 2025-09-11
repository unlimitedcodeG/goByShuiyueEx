-- name: GetUser :one
SELECT id, username, email, password_hash, full_name, created_at, updated_at
FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, username, email, password_hash, full_name, created_at, updated_at
FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT id, username, email, password_hash, full_name, created_at, updated_at
FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT id, username, email, full_name, created_at, updated_at
FROM users
ORDER BY created_at DESC;

-- name: CreateUser :one
INSERT INTO users (
    username, email, password_hash, full_name
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, username, email, password_hash, full_name, created_at, updated_at;

-- name: UpdateUser :one
UPDATE users
SET 
    username = $2,
    email = $3,
    full_name = $4
WHERE id = $1
RETURNING id, username, email, password_hash, full_name, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetPost :one
SELECT p.id, p.title, p.content, p.author_id, p.status, p.created_at, p.updated_at,
       u.username as author_username
FROM posts p
JOIN users u ON p.author_id = u.id
WHERE p.id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT p.id, p.title, p.content, p.author_id, p.status, p.created_at, p.updated_at,
       u.username as author_username
FROM posts p
JOIN users u ON p.author_id = u.id
WHERE p.status = $1
ORDER BY p.created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListPostsByAuthor :many
SELECT id, title, content, author_id, status, created_at, updated_at
FROM posts
WHERE author_id = $1
ORDER BY created_at DESC;

-- name: CreatePost :one
INSERT INTO posts (
    title, content, author_id, status
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, title, content, author_id, status, created_at, updated_at;

-- name: UpdatePost :one
UPDATE posts
SET 
    title = $3,
    content = $4,
    status = $5
WHERE id = $1 AND author_id = $2
RETURNING id, title, content, author_id, status, created_at, updated_at;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1 AND author_id = $2;

-- name: GetTag :one
SELECT id, name, created_at
FROM tags
WHERE id = $1 LIMIT 1;

-- name: GetTagByName :one
SELECT id, name, created_at
FROM tags
WHERE name = $1 LIMIT 1;

-- name: ListTags :many
SELECT id, name, created_at
FROM tags
ORDER BY name;

-- name: CreateTag :one
INSERT INTO tags (name) VALUES ($1)
RETURNING id, name, created_at;

-- name: DeleteTag :exec
DELETE FROM tags
WHERE id = $1;

-- name: AddPostTag :exec
INSERT INTO post_tags (post_id, tag_id) VALUES ($1, $2);

-- name: RemovePostTag :exec
DELETE FROM post_tags
WHERE post_id = $1 AND tag_id = $2;

-- name: GetPostTags :many
SELECT t.id, t.name, t.created_at
FROM tags t
JOIN post_tags pt ON t.id = pt.tag_id
WHERE pt.post_id = $1
ORDER BY t.name;

-- name: GetPostsByTag :many
SELECT p.id, p.title, p.content, p.author_id, p.status, p.created_at, p.updated_at
FROM posts p
JOIN post_tags pt ON p.id = pt.post_id
WHERE pt.tag_id = $1 AND p.status = 'published'
ORDER BY p.created_at DESC;
