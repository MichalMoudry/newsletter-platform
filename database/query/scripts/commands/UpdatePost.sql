UPDATE Post
SET Post_ID = $1,
    Post_Title = $2,
    Post_Content = $3,
    Post_Publishing_Date = $4,
    Post_Inserted_Datetime = $5,
    Post_Updated_Datetime = $6,
    Post_Author = $7,
    Nltr_ID = $8,
    concurrency_stamp = $9
WHERE Post_ID = $10
AND   Post_Author = $11