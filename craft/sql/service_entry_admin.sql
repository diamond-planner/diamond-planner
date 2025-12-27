SELECT users.id, users.email, clubs.service_requirement AS target, SUM(entries.minutes) AS totalMinutes
FROM serviceentries AS entries
         INNER JOIN users ON users.id = entries.member
         INNER JOIN clubs ON clubs.id = entries.club
WHERE 1 = 1
  AND entries.club = 'x2f8mayfly6zye3'
  AND entries.service_date BETWEEN '2025-01-01' AND '2025-12-31'
GROUP BY users.id, users.email
;
