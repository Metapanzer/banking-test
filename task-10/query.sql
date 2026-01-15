WITH BalanceSums AS (
    SELECT customer_id, SUM(balance) AS total_balance
    FROM balances
    GROUP BY customer_id
),
     TransactionSums AS (
         SELECT customer_id, SUM(transaction_amount) AS total_spent, COUNT(*) AS transaction_count
         FROM transactions
         GROUP BY customer_id
     ),
    CustomerTotals AS (
        SELECT
            c.iban,
            (COALESCE(b.total_balance, 0) - COALESCE(t.total_spent, 0)) AS amount,
            COALESCE(t.transaction_count, 0) AS transaction_count
        FROM customers c
        LEFT JOIN BalanceSums b ON c.id = b.customer_id
        LEFT JOIN TransactionSums t ON c.id = t.customer_id
        )
SELECT * FROM CustomerTotals
WHERE amount < 0
ORDER BY amount ASC;