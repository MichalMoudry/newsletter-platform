DELETE FROM users
WHERE
    email = $1
    AND
    concurrency_stamp = $2