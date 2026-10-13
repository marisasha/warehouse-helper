ALTER TABLE product_categories
    ADD CONSTRAINT product_categories_type_check
    CHECK (type IN ('vegetables', 'fruits', 'berries'));