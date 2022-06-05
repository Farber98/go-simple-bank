/* "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
); */

-- name: CreateTransfer :one
INSERT INTO transfers ( from_account_id, to_account_id,amount ) VALUES ($1, $2,$3) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers WHERE id = $1 LIMIT 1;

-- name: ListTransfersByAccount :many
SELECT * FROM transfers WHERE from_account_id = $1 OR to_account_id = $2 ORDER BY id LIMIT $3 OFFSET $4;