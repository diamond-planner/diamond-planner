SELECT users.id,
       users.email,
       CONCAT(users.first_name, ' ', users.last_name)                                       AS name,
       clubs.service_requirement                                                            AS target,
       COALESCE(SUM(serviceentries.minutes), 0)                                             AS total_minutes,
       CASE WHEN SUM(serviceentries.minutes) >= clubs.service_requirement THEN 1 ELSE 0 END AS target_met
FROM users
         LEFT JOIN serviceentries ON users.id = serviceentries.member
    AND strftime('%Y', serviceentries.service_date) = '2025'
    AND serviceentries.club = 'x2f8mayfly6zye3'
         LEFT JOIN clubs ON clubs.id = 'x2f8mayfly6zye3' -- Joined directly to get the requirement for this specific club
WHERE 1 = 1
  AND strftime('%Y', users.created) <= '2025'
  AND EXISTS (SELECT 1 FROM json_each(users.club) WHERE value = 'x2f8mayfly6zye3')
GROUP BY users.id, users.email, clubs.service_requirement;
