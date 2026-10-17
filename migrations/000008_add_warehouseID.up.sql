ALTER TABLE suppliers
    ADD COLUMN warehouse_id UUID;

ALTER TABLE suppliers
    ADD CONSTRAINT uq_supplier_wi_name UNIQUE (warehouse_id,name);