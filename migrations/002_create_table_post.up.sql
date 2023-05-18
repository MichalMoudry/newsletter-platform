BEGIN;

CREATE TABLE Post (
    Post_ID integer PRIMARY KEY,
    Post_Title VARCHAR(255) NOT NULL UNIQUE,
    Post_Content VARCHAR() NOT NULL,
    Post_Publishing_Date DATE NOT NULL,
    Post_Inserted_Datetime TIMESTAMP NOT NULL,
    Post_Updated_Datetime TIMESTAMP NOT NULL,
    Post_Author VARCHAR(255) NOT NULL,
    Nltr_ID integer FOREIGN KEY NOT NULL,
    concurrency_stamp VARCHAR(255) NOT NULL
);

COMMIT;