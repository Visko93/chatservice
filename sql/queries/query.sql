-- name: FindChatByID :one
SELECT * FROM chats WHERE id = $1;

-- name: CreateChat :exec
INSERT INTO chats 
    (id, user_id, initial_message_id, status, token_usage, model, model_max_tokens,temperature, top_p, n, stop, max_tokens, presence_penalty, frequency_penalty, created_at, updated_at)
    VALUES($1 ,$2 ,$3 ,$4 ,$5 ,$6 ,$7 ,$8 ,$9 ,$10 ,$11 ,$12 ,$13 ,$14 ,$15 ,$16 );

-- name: AddMessage :exec
INSERT INTO messages (id, chat_id, role, content, tokens, model, erased, order_msg, created_at) VALUES($1 ,$2 ,$3 ,$4 ,$5 ,$6 ,$7 ,$8 ,$9);

-- name: FindMessagesByChatID :many
SELECT * FROM messages WHERE erased=FALSE and chat_id = $1 order by order_msg asc;

-- name: FindErasedMessagesByChatID :many
SELECT * FROM messages WHERE erased=TRUE and chat_id = $1 order by order_msg asc;

-- name: SaveChat :exec
UPDATE chats SET user_id = $2, initial_message_id = $3, status = $4, token_usage = $5, model = $6, model_max_tokens = $7, temperature = $8, top_p = $9, n = $10, stop = $11, max_tokens = $12, presence_penalty = $13, frequency_penalty = $14, updated_at = $15 WHERE id = $1;

-- name: DeleteChatMessages :exec
DELETE FROM messages WHERE chat_id = $1;

-- name: DeleteErasedChatMessages :exec
DELETE FROM messages WHERE erased=TRUE and chat_id = $1;