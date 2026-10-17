-- +migrate Down
BEGIN;

DROP TABLE IF EXISTS supplies CASCADE;

CREATE TABLE supplies (
    id            UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    warehouse_id  UUID         NOT NULL REFERENCES warehouses(id) ON DELETE RESTRICT,
    supplier_id   UUID         NOT NULL REFERENCES suppliers(id)  ON DELETE RESTRICT,
    creator_id    BIGINT       NOT NULL REFERENCES users(id)      ON DELETE RESTRICT,
    status        TEXT         NOT NULL DEFAULT 'draft',
    accepted_at   TIMESTAMPTZ,
    comment       TEXT,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_supplies_status
        CHECK (status IN ('draft', 'accepted', 'cancelled'))
);

CREATE INDEX idx_supplies_warehouse ON supplies (warehouse_id);
CREATE INDEX idx_supplies_supplier  ON supplies (supplier_id);
CREATE INDEX idx_supplies_creator   ON supplies (creator_id);
CREATE INDEX idx_supplies_status    ON supplies (status);

CREATE TABLE supply_items (
    id            UUID          PRIMARY KEY DEFAULT gen_random_uuid(),
    supply_id     UUID          NOT NULL REFERENCES supplies(id) ON DELETE CASCADE,
    product_id    UUID          NOT NULL REFERENCES products(id) ON DELETE RESTRICT,
    batch_id      UUID,
    quantity      NUMERIC(12,3) NOT NULL,
    cost_per_unit NUMERIC(12,2) NOT NULL,
    harvest_date  DATE,
    expiry_date   DATE,

    CONSTRAINT chk_supply_items_qty_positive
        CHECK (quantity > 0),

    CONSTRAINT chk_supply_items_cost_non_negative
        CHECK (cost_per_unit >= 0)
);

CREATE INDEX idx_supply_items_supply  ON supply_items (supply_id);
CREATE INDEX idx_supply_items_product ON supply_items (product_id);
CREATE INDEX idx_supply_items_batch   ON supply_items (batch_id);

COMMIT;