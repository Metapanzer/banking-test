SELECT
    CONCAT(e.FullName, ' (ID: ', e.EmployeeID, ') has a performance rating of: ',
           CASE
               WHEN e.PerformanceScore >= 90 THEN 'Outstanding'
               WHEN e.PerformanceScore >= 75 THEN 'Exceeds Expectations'
               WHEN e.PerformanceScore >= 50 THEN 'Meets Expectations'
               ELSE 'Needs Improvement'
               END)
    FROM EMPLOYEES e
ORDER BY e.EmployeeID ASC;