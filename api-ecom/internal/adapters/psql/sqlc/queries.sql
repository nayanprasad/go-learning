-- name: CreateProduct :one
INSERT INTO products (name, price, quantity)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products
SET name = $2, price = $3, quantity = $4
WHERE id = $1
RETURNING *;

-- name: ListProducts :many
SELECT
  *
FROM
  products;

-- name: GetProductById :one
SELECT
  *
FROM
  products
WHERE id = $1
LIMIT 1;

-- name: DeleteProductById :exec
DELETE FROM products
WHERE id = $1;

-- name: CrateOrder :one
insert into orders(user_id)
values ($1)
RETURNING *;

-- name: CrateOrderItems :one
insert into order_items(order_id, product_id, price, quantity)
values ($1, $2, $3, $4)
RETURNING *;