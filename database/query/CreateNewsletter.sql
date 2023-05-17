INSERT INTO Newsletter (    
    Nltr_ID,
    Nltr_Name,
    Nltr_Description,
    Nltr_Inserted_Datetime,
    Nltr_Updated_Datetime,
    Nltr_Author,
    concurrency_stamp
    ) VALUES
    (:Nltr_ID,
     :Nltr_Name,
     :Nltr_Description,
     :Nltr_Inserted_Datetime,
     :Nltr_Updated_Datetime,
     :Nltr_Author,
     :concurrency_stamp)