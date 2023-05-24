SELECT
    p.Post_Title,
    p.Post_Content,
	u.user_name
FROM
	post as p
JOIN
	users as u on u.id = p.post_author
WHERE
    p.Post_ID = $1;