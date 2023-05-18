UPDATE Newsletter
SET Nltr_ID = $1,
    Nltr_Name = $1,
    Nltr_Description = $3,
    Nltr_Inserted_Datetime = $4,
    Nltr_Updated_Datetime = $5,
    Nltr_Author = $6,
    concurrency_stamp = $7,
WHERE Nltr_ID = $8,
AND   Nltr_Author = $9