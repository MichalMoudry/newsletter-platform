DELETE FROM
    newsletter_subscriptions
WHERE
    subscriber_id = $1
AND
    newsletter_id = $2