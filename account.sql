-- name: createAuthor: one
INSERT INTO accounts(owner,baalnce,currency)
VALUES($1,$2,$3)
RETURNING *;


SELECT * FROM accounts
ORDER BY id LIMIT $1
OFFSET $2;


