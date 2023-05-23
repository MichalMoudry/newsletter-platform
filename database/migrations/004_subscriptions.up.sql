BEGIN;

CREATE TABLE subscribers (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    date_added TIMESTAMP NOT NULL
);

CREATE TABLE newsletter_subscriptions (
    id UUID PRIMARY KEY,
    newsletter_id UUID,
    subscriber_id VARCHAR(255),
    CONSTRAINT fk_subscriber
        FOREIGN KEY(subscriber_id)
            REFERENCES subscribers(email)
                ON DELETE CASCADE,
    CONSTRAINT fk_newsletter
        FOREIGN KEY(newsletter_id)
            REFERENCES Newsletter(Nltr_ID)
                ON DELETE CASCADE
);

COMMIT;