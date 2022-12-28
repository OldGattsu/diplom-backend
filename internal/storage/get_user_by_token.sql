SELECT u.id, u.name, u.email, u.is_admin
FROM tokens AS t
         LEFT JOIN users AS u ON t.user_id = u.id
WHERE t.token = $1