-- name: StorePrice :exec
INSERT INTO prices (code, price, timestamp, currency) VALUES (?, ?, ?, ?);

-- name: LoadPrices :many
SELECT code, price, timestamp, currency FROM prices ORDER BY timestamp DESC;