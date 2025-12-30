SELECT
    users.id,
    users.email,
    concat(users.first_name, ' ', users.last_name) AS name,
    clubs.service_requirement AS target,
    SUM(serviceentries.minutes) AS total_minutes,
    CASE WHEN SUM(serviceentries.minutes) >= clubs.service_requirement THEN 1 ELSE 0 END AS target_met
FROM serviceentries
         INNER JOIN users ON users.id = serviceentries.member
         INNER JOIN clubs ON clubs.id = serviceentries.club
WHERE 1 = 1
  AND serviceentries.club = 'x2f8mayfly6zye3'
  AND strftime('%Y', serviceentries.service_date) = '2025'
GROUP BY users.id, users.email
;
