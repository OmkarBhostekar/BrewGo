DROP INDEX IF EXISTS "idx_counter_orders_user_id";
DROP INDEX IF EXISTS "idx_order_date";
DROP TABLE IF EXISTS "counter_orders";

DROP INDEX IF EXISTS "idx_counter_order_items_order_id";
DROP INDEX IF EXISTS "idx_counter_order_items_product_id";
DROP TABLE IF EXISTS "counter_order_items";