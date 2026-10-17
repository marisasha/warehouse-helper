ALTER TABLE products
    DROP CONSTRAINT IF EXISTS uq_products_cat_sku

ALTER TABLE products
    ADD CONSTRAINT uq_products_sku UNIQUE (sku);

