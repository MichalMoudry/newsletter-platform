BEGIN;

CREATE TABLE Post (
    Post_ID VARCHAR(255) PRIMARY KEY,
    Post_Title VARCHAR(255) NOT NULL UNIQUE,
    Post_Content VARCHAR(17500) NOT NULL, --~3.5k words
    Post_Publishing_Date DATE NOT NULL,
    Post_Inserted_Datetime TIMESTAMP NOT NULL,
    Post_Updated_Datetime TIMESTAMP NOT NULL,
    Post_Author VARCHAR(255) NOT NULL,
    Nltr_ID VARCHAR(255) NOT NULL,
    concurrency_stamp VARCHAR(255) NOT NULL,
    CONSTRAINT fk_newsletter_id
        FOREIGN KEY(Nltr_ID)
            REFERENCES Newsletter(Nltr_ID)
            ON DELETE CASCADE
);

COMMIT;