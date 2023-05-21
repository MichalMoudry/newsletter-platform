SELECT
    t.id,
    t.expiration_date,
    t.date_added
FROM
    password_reset_tokens as t
WHERE
    t.id = $1