BEGIN;

-- product_categories: добавляем warehouse_id
ALTER TABLE product_categories
    ADD COLUMN warehouse_id UUID;

UPDATE product_categories pc
SET warehouse_id = sub.warehouse_id
FROM (
    SELECT DISTINCT ON (category_id) category_id, warehouse_id
    FROM products
    ORDER BY category_id, created_at
) AS sub
WHERE pc.id = sub.category_id;

ALTER TABLE product_categories
    ALTER COLUMN warehouse_id SET NOT NULL;

ALTER TABLE product_categories
    ADD CONSTRAINT fk_product_categories_warehouse
        FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE;

ALTER TABLE product_categories
    DROP CONSTRAINT uq_product_categories_name;

ALTER TABLE product_categories
    ADD CONSTRAINT uq_product_categories_wh_name UNIQUE (warehouse_id, name);

CREATE INDEX idx_product_categories_warehouse
    ON product_categories (warehouse_id);

-- products: удаляем warehouse_id

DROP INDEX IF EXISTS idx_products_warehouse;

ALTER TABLE products
    DROP CONSTRAINT IF EXISTS uq_products_wh_sku;

ALTER TABLE products
    DROP CONSTRAINT IF EXISTS products_warehouse_id_fkey;

ALTER TABLE products
    DROP COLUMN warehouse_id;

ALTER TABLE products
    ADD CONSTRAINT uq_products_sku UNIQUE (sku);

COMMIT;