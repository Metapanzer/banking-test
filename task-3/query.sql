WITH TariffUsage AS (
    SELECT
        r.account_id,
        t.name AS tariff_name,
        COUNT(*) AS usage_count
    FROM readings r
    JOIN tariffs t ON r.tariff_id = t.id
    GROUP BY r.account_id, t.name
)

SELECT
    a.username,
    a.email,
    (SELECT tu.tariff_name
     FROM TariffUsage tu
     WHERE tu.account_id = a.id
     ORDER BY tu.usage_count DESC
     LIMIT 1) AS most_frequent_tariff,
    SUM(r.amount) AS total_consumption,
    ROUND(AVG(r.amount*t.cost),2) AS average_cost_per_reading
FROM accounts a
JOIN readings r ON a.id = r.account_id
JOIN tariffs t ON r.tariff_id = t.id
GROUP BY a.id, a.username, a.email
ORDER BY a.username ASC;