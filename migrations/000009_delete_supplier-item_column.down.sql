ALTER TABLE batches
    ADD COLUMN warehouse_id UUID;
    
ALTER TABLE batches
    ADD COLUMN supply_item_id UUID;