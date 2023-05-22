UPDATE Newsletter
SET
    Nltr_Name = :Nltr_Name,
    Nltr_Description = :Nltr_Description,
    Nltr_Updated_Datetime = :Nltr_Updated_Datetime,
    concurrency_stamp = :new_concurrency_stamp,
WHERE
    Nltr_ID = :Nltr_ID,
AND
    Nltr_Author = :Nltr_Author