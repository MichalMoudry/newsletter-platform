INSERT INTO users (
    id,
    email,
    user_name,
    user_role,
    password_hash,
    date_added,
    date_updated,
    concurrency_stamp
) VALUES
    (:id, :email, :user_name, :user_role, :password_hash, :date_added, :date_updated, :concurrency_stamp)