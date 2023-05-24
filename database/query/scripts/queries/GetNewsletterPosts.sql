SELECT
	p.post_title,
	(SELECT user_name FROM users WHERE id = p.post_author) as user_name,
	p.post_content
FROM
    newsletter as n
JOIN
	post as p on p.nltr_id = n.nltr_id
WHERE
    n.nltr_id = $1