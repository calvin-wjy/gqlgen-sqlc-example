-- name: ListStudents :many
SELECT * FROM students
ORDER BY name;

-- name: CreateStudent :one
INSERT INTO students (name, nim)
VALUES ($1, $2)
RETURNING *;