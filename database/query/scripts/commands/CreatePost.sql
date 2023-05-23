INSERT INTO Post (    
    Post_ID,
    Post_Title,
    Post_Content,
    Post_Publishing_Date,
    Post_Inserted_Datetime,
    Post_Updated_Datetime,
    Post_Author,
    Nltr_ID,
    concurrency_stamp
    ) VALUES
    (:Post_ID,
     :Post_Title,
     :Post_Content,
     :Post_Publishing_Date,
     :Post_Inserted_Datetime,
     :Post_Updated_Datetime,
     :Nltr_ID,
     :Post_Author,
     :concurrency_stamp)
RETURNING
    Post_ID