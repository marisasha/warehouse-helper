-- WAREHOUSES
CREATE TABLE warehouses (
    id          UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    name        TEXT         NOT NULL,
    address     TEXT,
    country     TEXT,
    region      TEXT,
    district    TEXT,
    city        TEXT,
    timezone    TEXT         NOT NULL DEFAULT 'Europe/Moscow',
    latitude    NUMERIC(9,6),
    longitude   NUMERIC(9,6),
    is_active   BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TRIGGER trg_warehouses_updated_at
    BEFORE UPDATE ON warehouses
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();


-- WAREHOUSE_EMPLOYEES
CREATE TABLE warehouse_employees (
    id            UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    warehouse_id  UUID         NOT NULL REFERENCES warehouses(id) ON DELETE CASCADE,
    user_id       BIGINT       NOT NULL REFERENCES users(id)      ON DELETE CASCADE,
    role          TEXT         NOT NULL,
    hired_at      TIMESTAMPTZ,
    fired_at      TIMESTAMPTZ,
    is_active     BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_warehouse_employees_role
        CHECK (role IN ('owner', 'admin', 'worker')),

    CONSTRAINT uq_warehouse_employees_wh_user
        UNIQUE (warehouse_id, user_id)
);

CREATE INDEX idx_warehouse_employees_warehouse ON warehouse_employees (warehouse_id);
CREATE INDEX idx_warehouse_employees_user      ON warehouse_employees (user_id);
CREATE INDEX idx_warehouse_employees_role      ON warehouse_employees (role);

CREATE TRIGGER trg_warehouse_employees_updated_at
    BEFORE UPDATE ON warehouse_employees
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();


-- PRODUCT_CATEGORIES
CREATE TABLE product_categories (
    id    UUID  PRIMARY KEY DEFAULT gen_random_uuid(),
    name  TEXT  NOT NULL,
    type  TEXT  NOT NULL,

    CONSTRAINT uq_product_categories_name UNIQUE (name)
);


-- PRODUCTS
CREATE TABLE products (
    id              UUID          PRIMARY KEY DEFAULT gen_random_uuid(),
    category_id     UUID          NOT NULL REFERENCES product_categories(id) ON DELETE RESTRICT,
    warehouse_id    UUID          NOT NULL REFERENCES warehouses(id)         ON DELETE CASCADE,
    sku             TEXT          NOT NULL,
    name            TEXT          NOT NULL,
    unit            TEXT          NOT NULL,
    price_per_unit  NUMERIC(12,2) NOT NULL,
    description     TEXT,
    min_stock       NUMERIC(12,3) NOT NULL DEFAULT 0,
    is_active       BOOLEAN       NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ   NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_products_unit
        CHECK (unit IN ('kg', 'pcs', 'box', 'liter')),

    CONSTRAINT chk_products_price_non_negative
        CHECK (price_per_unit >= 0),

    CONSTRAINT chk_products_min_stock_non_negative
        CHECK (min_stock >= 0),

    CONSTRAINT uq_products_wh_sku UNIQUE (warehouse_id, sku)
);

CREATE INDEX idx_products_warehouse ON products (warehouse_id);
CREATE INDEX idx_products_category  ON products (category_id);
CREATE INDEX idx_products_is_active ON products (is_active) WHERE is_active = TRUE;

CREATE TRIGGER trg_products_updated_at
    BEFORE UPDATE ON products
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();


-- SUPPLIERS
CREATE TABLE suppliers (
    id          UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    name        TEXT         NOT NULL,
    contact     TEXT,
    phone       TEXT,
    address     TEXT,
    is_active   BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TRIGGER trg_suppliers_updated_at
    BEFORE UPDATE ON suppliers
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();


-- SUPPLIES (поставки)
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


-- SUPPLY_ITEMS (позиции поставки)
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


-- BATCHES (партии)
CREATE TABLE batches (
    id              UUID          PRIMARY KEY DEFAULT gen_random_uuid(),
    warehouse_id    UUID          NOT NULL REFERENCES warehouses(id)   ON DELETE RESTRICT,
    product_id      UUID          NOT NULL REFERENCES products(id)     ON DELETE RESTRICT,
    supply_item_id  UUID          REFERENCES supply_items(id)          ON DELETE SET NULL,
    quantity        NUMERIC(12,3) NOT NULL,
    remaining       NUMERIC(12,3) NOT NULL,
    cost_per_unit   NUMERIC(12,2) NOT NULL,
    harvest_date    DATE,
    expiry_date     DATE,
    received_at     TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    status          TEXT          NOT NULL DEFAULT 'active',

    CONSTRAINT chk_batches_status
        CHECK (status IN ('active', 'expired', 'written_off')),

    CONSTRAINT chk_batches_qty_non_negative
        CHECK (quantity >= 0),

    CONSTRAINT chk_batches_remaining_non_negative
        CHECK (remaining >= 0),

    CONSTRAINT chk_batches_remaining_le_quantity
        CHECK (remaining <= quantity),

    CONSTRAINT chk_batches_cost_non_negative
        CHECK (cost_per_unit >= 0)
);

CREATE INDEX idx_batches_warehouse         ON batches (warehouse_id);
CREATE INDEX idx_batches_status            ON batches (status);
-- FEFO: ищем активные партии по товару, отсортированные по сроку годности.
CREATE INDEX idx_batches_product_expiry    ON batches (product_id, expiry_date ASC NULLS LAST)
    WHERE status = 'active';

-- После создания batches можно связать supply_items.batch_id.
ALTER TABLE supply_items
    ADD CONSTRAINT fk_supply_items_batch
    FOREIGN KEY (batch_id) REFERENCES batches(id) ON DELETE SET NULL;


-- ORDERS (заявки)
CREATE SEQUENCE order_number_seq START 1;

CREATE TABLE orders (
    id                  UUID          PRIMARY KEY DEFAULT gen_random_uuid(),
    warehouse_id        UUID          NOT NULL REFERENCES warehouses(id) ON DELETE RESTRICT,
    customer_id         BIGINT        NOT NULL REFERENCES users(id)      ON DELETE RESTRICT,
    assignee_id         BIGINT        REFERENCES users(id)               ON DELETE SET NULL,
    order_number        TEXT          NOT NULL UNIQUE,
    status              TEXT          NOT NULL DEFAULT 'pending',
    requested_ship_date DATE,
    ship_window_start   TIMESTAMPTZ,
    ship_window_end     TIMESTAMPTZ,
    customer_comment    TEXT,
    staff_comment       TEXT,
    rejection_reason    TEXT,
    total_amount        NUMERIC(12,2) NOT NULL DEFAULT 0,
    confirmed_at        TIMESTAMPTZ,
    shipped_at          TIMESTAMPTZ,
    completed_at        TIMESTAMPTZ,
    cancelled_at        TIMESTAMPTZ,
    created_at          TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ   NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_orders_status
        CHECK (status IN ('pending', 'processing', 'ready', 'completed', 'rejected', 'cancelled')),

    CONSTRAINT chk_orders_total_non_negative
        CHECK (total_amount >= 0),

    CONSTRAINT chk_orders_ship_window
        CHECK (
            ship_window_start IS NULL
            OR ship_window_end IS NULL
            OR ship_window_start <= ship_window_end
        )
);

CREATE INDEX idx_orders_warehouse   ON orders (warehouse_id);
CREATE INDEX idx_orders_customer    ON orders (customer_id);
CREATE INDEX idx_orders_assignee    ON orders (assignee_id);
CREATE INDEX idx_orders_status      ON orders (status);
CREATE INDEX idx_orders_created_at  ON orders (created_at DESC);

CREATE TRIGGER trg_orders_updated_at
    BEFORE UPDATE ON orders
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- Автогенерация order_number вида ORD-2026-000123.
CREATE OR REPLACE FUNCTION generate_order_number()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.order_number IS NULL OR NEW.order_number = '' THEN
        NEW.order_number :=
            'ORD-' || TO_CHAR(NOW(), 'YYYY') || '-' ||
            LPAD(nextval('order_number_seq')::TEXT, 6, '0');
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_orders_number
    BEFORE INSERT ON orders
    FOR EACH ROW EXECUTE FUNCTION generate_order_number();


-- ORDER_ITEMS
CREATE TABLE order_items (
    id             UUID          PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id       UUID          NOT NULL REFERENCES orders(id)   ON DELETE CASCADE,
    product_id     UUID          NOT NULL REFERENCES products(id) ON DELETE RESTRICT,
    quantity       NUMERIC(12,3) NOT NULL,
    price_per_unit NUMERIC(12,2) NOT NULL,
    total          NUMERIC(12,2) NOT NULL,
    comment        TEXT,

    CONSTRAINT chk_order_items_qty_positive
        CHECK (quantity > 0),

    CONSTRAINT chk_order_items_price_non_negative
        CHECK (price_per_unit >= 0),

    CONSTRAINT chk_order_items_total_non_negative
        CHECK (total >= 0)
);

CREATE INDEX idx_order_items_order   ON order_items (order_id);
CREATE INDEX idx_order_items_product ON order_items (product_id);


-- ORDER_ITEM_BATCH_ALLOCATIONS (FEFO-распределение)
CREATE TABLE order_item_batch_allocations (
    id            UUID          PRIMARY KEY DEFAULT gen_random_uuid(),
    order_item_id UUID          NOT NULL REFERENCES order_items(id) ON DELETE CASCADE,
    batch_id      UUID          NOT NULL REFERENCES batches(id)     ON DELETE RESTRICT,
    quantity      NUMERIC(12,3) NOT NULL,

    CONSTRAINT chk_oi_alloc_qty_positive
        CHECK (quantity > 0)
);

CREATE INDEX idx_oi_alloc_order_item ON order_item_batch_allocations (order_item_id);
CREATE INDEX idx_oi_alloc_batch      ON order_item_batch_allocations (batch_id);


-- ORDER_STATUS_HISTORY
CREATE TABLE order_status_history (
    id          UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id    UUID         NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    old_status  TEXT,
    new_status  TEXT         NOT NULL,
    changed_by  BIGINT       REFERENCES users(id) ON DELETE SET NULL,
    comment     TEXT,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_order_status_history_order ON order_status_history (order_id);
CREATE INDEX idx_order_status_history_time  ON order_status_history (created_at DESC);


-- ORDER_MESSAGES (чат по заявке)
CREATE TABLE order_messages (
    id          UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id    UUID         NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    sender_id   BIGINT       NOT NULL REFERENCES users(id)  ON DELETE RESTRICT,
    body        TEXT         NOT NULL,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_order_messages_order   ON order_messages (order_id);
CREATE INDEX idx_order_messages_sender  ON order_messages (sender_id);
CREATE INDEX idx_order_messages_time    ON order_messages (created_at DESC);