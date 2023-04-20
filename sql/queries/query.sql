-- name: FindChatByID :one
SELECT * FROM chats WHERE id = ?;

-- name: CreateChat :exec
INSERT INTO chats (id, user_id, initial_message, status, token_usage, model, model_max_tokens, temperature, stop, max_tokens, presence_penalty, frequency_penalty, created_at, updated_at) 
VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?);