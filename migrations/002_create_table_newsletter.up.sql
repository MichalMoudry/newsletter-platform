BEGIN;

CREATE TABLE Newsletter (
    Nltr_ID VARCHAR(255) PRIMARY KEY,
    Nltr_Name VARCHAR(255) NOT NULL UNIQUE,
    Nltr_Description VARCHAR(5000) NOT NULL,
    Nltr_Inserted_Datetime TIMESTAMP NOT NULL,
    Nltr_Updated_Datetime TIMESTAMP NOT NULL,
    Nltr_Author uuid NOT NULL,
    concurrency_stamp VARCHAR(255) NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(Nltr_Author)
            REFERENCES users(id)
                ON DELETE CASCADE
);

COMMIT;