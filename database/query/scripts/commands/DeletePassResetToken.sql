DELETE FROM
    password_reset_tokens
WHERE
    id = :token_id