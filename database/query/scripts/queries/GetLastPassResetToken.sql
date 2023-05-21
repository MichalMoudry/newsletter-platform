SELECT
    t.id,
    t.expiration_date,
    t.date_added
FROM
    password_reset_tokens as t
WHERE
    t.email = $1
ORDER BY
    t.expiration_date DESC
LIMIT
    1