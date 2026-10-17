BEGIN;

DROP TABLE IF EXISTS supply_items CASCADE;
DROP TABLE IF EXISTS supplies CASCADE;

CREATE TABLE supplies (
    id              UUID          PRIMARY KEY DEFAULT gen_random_uuid(),
    supplier_id     UUID          NOT NULL REFERENCES suppliers(id)  ON DELETE RESTRICT,
    product_id      UUID          NOT NULL REFERENCES products(id)   ON DELETE RESTRICT,
    creator_id      BIGINT        NOT NULL REFERENCES users(id)      ON DELETE RESTRICT,
    status          TEXT          NOT NULL DEFAULT 'draft',
    comment         TEXT,
    quantity        NUMERIC(12,3) NOT NULL,
    cost_per_unit   NUMERIC(12,2) NOT NULL,
    harvest_date    DATE,
    expiry_date     DATE,
    accepted_at     TIMESTAMPTZ,
    created_at      TIMESTAMPTZ   NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_supplies_status
        CHECK (status IN ('draft', 'accepted', 'cancelled')),

    CONSTRAINT chk_supplies_qty_positive
        CHECK (quantity > 0),

    CONSTRAINT chk_supplies_cost_non_negative
        CHECK (cost_per_unit >= 0)
);

CREATE INDEX idx_supplies_supplier  ON supplies (supplier_id);
CREATE INDEX idx_supplies_creator   ON supplies (creator_id);
CREATE INDEX idx_supplies_status    ON supplies (status);
CREATE INDEX idx_supplies_product   ON supplies (product_id);

COMMIT;