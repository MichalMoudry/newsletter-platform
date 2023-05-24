SELECT
    n.Nltr_Name,
    n.Nltr_Author,
    s.subscriber_id
FROM
    newsletter_subscriptions as s
JOIN
    Newsletter as n on n.Nltr_ID = s.newsletter_id
WHERE
    s.newsletter_id = $1