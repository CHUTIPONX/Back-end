SELECT 
    c.customers_id,
    c.customers_firstname,
    c.customers_lastname,
    o.menu_id,
    o.quantity,
    o.total
FROM tbl_customers c
JOIN tbl_orders o ON c.customers_id = o.customer_id
WHERE o.total < 500;
