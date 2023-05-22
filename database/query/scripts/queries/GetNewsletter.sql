SELECT
    n.Nltr_ID,
    n.Nltr_Name,
    n.Nltr_Description,
    u.user_name
FROM
    Newsletter as n
JOIN
    users as u on u.id = n.Nltr_Author
WHERE
    Nltr_ID = $1