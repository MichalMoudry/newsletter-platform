UPDATE users
SET
    email = :email,
    user_name = :user_name,
    date_updated = :date_updated,
    concurrency_stamp = :new_concurrency_stamp
WHERE
    id = :user_id
    AND concurrency_stamp = :old_concurrency_stamp