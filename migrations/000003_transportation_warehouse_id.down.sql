BEGIN;

ALTER TABLE products
    DROP CONSTRAINT IF EXISTS uq_products_sku;

ALTER TABLE products
    ADD COLUMN warehouse_id UUID;

UPDATE products p
SET warehouse_id = pc.warehouse_id
FROM product_categories pc
WHERE p.category_id = pc.id;

ALTER TABLE products
    ALTER COLUMN warehouse_id SET NOT NULL;

ALTER TABLE products
    ADD CONSTRAINT products_warehouse_id_fkey
        FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE;

ALTER TABLE products
    ADD CONSTRAINT uq_products_wh_sku UNIQUE (warehouse_id, sku);

CREATE INDEX idx_products_warehouse
    ON products (warehouse_id);

DROP INDEX IF EXISTS idx_product_categories_warehouse;

ALTER TABLE product_categories
    DROP CONSTRAINT IF EXISTS uq_product_categories_wh_name;

ALTER TABLE product_categories
    ADD CONSTRAINT uq_product_categories_name UNIQUE (name);

ALTER TABLE product_categories
    DROP CONSTRAINT IF EXISTS fk_product_categories_warehouse;

ALTER TABLE product_categories
    DROP COLUMN warehouse_id;

COMMIT;