SELECT u.id, u.name, u.email
FROM tokens AS t
         LEFT JOIN users AS u ON t.user_id = u.id
WHERE t.token = $1