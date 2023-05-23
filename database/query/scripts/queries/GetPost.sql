SELECT
    p.Post_ID,
    p.Post_Title,
    p.Post_Content,
    n.Nltr_Description,
    u.user_name
FROM
    Post as p
JOIN
    Newsletter as n on p.Nltr_ID = n.Nltr_ID
JOIN
    users as u on u.id = n.Nltr_Author
WHERE
    p.Nltr_ID = $1;