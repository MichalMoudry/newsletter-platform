SELECT
    id,
    email
FROM
    subscribers
WHERE
    email = $1