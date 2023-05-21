INSERT INTO password_reset_tokens (
    id,
    email,
    expiration_date,
    date_added
) VALUES
    (:id, :email, :expiration_date, :date_added)