DROP TABLE IF EXISTS order_messages;
DROP TABLE IF EXISTS order_status_history;
DROP TABLE IF EXISTS order_item_batch_allocations;
DROP TABLE IF EXISTS order_items;
DROP TRIGGER IF EXISTS trg_orders_number ON orders;
DROP FUNCTION IF EXISTS generate_order_number();
DROP TABLE IF EXISTS orders;
DROP SEQUENCE IF EXISTS order_number_seq;

ALTER TABLE supply_items DROP CONSTRAINT IF EXISTS fk_supply_items_batch;
DROP TABLE IF EXISTS batches;
DROP TABLE IF EXISTS supply_items;
DROP TABLE IF EXISTS supplies;
DROP TABLE IF EXISTS suppliers;

DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS product_categories;

DROP TABLE IF EXISTS warehouse_employees;
DROP TABLE IF EXISTS warehouses;