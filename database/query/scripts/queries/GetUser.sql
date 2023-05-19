SELECT
    e.email,
    e.user_name,
    e.user_role,
    e.concurrency_stamp
FROM
    users AS e
WHERE
    e.email = $1