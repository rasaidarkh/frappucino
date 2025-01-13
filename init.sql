-- CREATE DATABASE frappuccino_db
--     WITH
--     OWNER = postgres
--     ENCODING = 'UTF8'
--     LOCALE_PROVIDER = 'libc'
--     CONNECTION LIMIT = -1
--     IS_TEMPLATE = False;


-- enums
CREATE TYPE sex AS ENUM ('female', 'male', 'other');
CREATE TYPE transaction_type AS ENUM
    ('incoming', 'outgoing', 'zero',
        'expired', 'returned', 'damaged',
        'adjusted', 'donated');
CREATE TYPE order_status AS ENUM ('created', 'pending', 'processing', 'completed', 'canceled');


-- create table queries
CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(50) NOT NULL UNIQUE,
    description VARCHAR(500) NOT NULL DEFAULT ''
);

CREATE TABLE menu_items
(
    menu_id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES categories(category_id) ON DELETE RESTRICT ON UPDATE CASCADE,
    menu_item_name VARCHAR(70) NOT NULL,
    description VARCHAR(500) NOT NULL DEFAULT '',
    price decimal NOT NULL
);


CREATE TABLE price_history
(
    price_history_id SERIAL PRIMARY KEY,
    menu_id INTEGER REFERENCES menu_items(menu_id) ON DELETE RESTRICT ON UPDATE CASCADE,
    old_price decimal NOT NULL,
    new_price decimal NOT NULL,
    change_date timestamptz NOT NULL DEFAULT NOW()
);



CREATE TABLE customers
(
    customer_id SERIAL PRIMARY KEY,
    customer_name VARCHAR(100) NOT NULL,
    age SMALLINT NOT NULL,
    sex sex NOT NULL,
    registration_date timestamptz NOT NULL DEFAULT NOW(),
    allergens VARCHAR(500)[]
);


CREATE TABLE inventory
(
    inventory_id SERIAL PRIMARY KEY,
    inventory_name VARCHAR(100) NOT NULL,
    quantity DECIMAL NOT NULL DEFAULT 0,
    unit VARCHAR(50) NOT NULL,
    allergens VARCHAR(500)[]
);

CREATE TABLE inventory_transactions
(
    inventory_transaction_id SERIAL PRIMARY KEY,
    inventory_id INTEGER REFERENCES inventory(inventory_id) ON DELETE RESTRICT ON UPDATE CASCADE,
    transaction_type transaction_type NOT NULL,
    quantity_changed decimal NOT NULL,
    transaction_date timestamptz NOT NULL DEFAULT NOW()
);

CREATE TABLE menu_items_ingredients
(
    menu_id INTEGER REFERENCES menu_items(menu_id) ON DELETE RESTRICT ON UPDATE CASCADE,
    inventory_id INTEGER REFERENCES inventory(inventory_id) ON DELETE RESTRICT ON UPDATE CASCADE,
    quantity decimal NOT NULL
);


CREATE TABLE orders
(
    order_id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES customers(customer_id) ON DELETE RESTRICT ON UPDATE CASCADE,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    status order_status NOT NULL DEFAULT 'created',
    total_amount decimal NOT NULL
);

CREATE TABLE order_status_history
(
    order_status_history_id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(order_id) ON DELETE RESTRICT ON UPDATE CASCADE,
    status order_status NOT NULL DEFAULT 'created',
    change_date timestamptz NOT NULL DEFAULT NOW()
);


CREATE TABLE order_items
(
    order_item_id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(order_id) ON DELETE RESTRICT ON UPDATE CASCADE,
    menu_id INTEGER REFERENCES menu_items(menu_id) ON DELETE RESTRICT ON UPDATE CASCADE,
    quantity INTEGER NOT NULL,
    price_at_order DECIMAL NOT NULL,
    customization_info jsonb NOT NULL
);


-- indices
CREATE INDEX idx_menu_items_name ON menu_items(menu_item_name);
CREATE INDEX idx_customers_name ON customers(customer_name);
CREATE INDEX idx_inventory_name ON Inventory(inventory_name);
CREATE INDEX idx_price_history_new_price ON price_history (new_price);
CREATE INDEX idx_inventory_transactions_transaction_date ON inventory_transactions (transaction_date);
CREATE INDEX idx_order_items_customization ON order_items (customization_info);
CREATE INDEX idx_order_status_history_change_date ON order_status_history (change_date);
CREATE INDEX idx_order_created_at ON orders (created_at);

-- Insert sample categories
INSERT INTO categories (category_name, description)
VALUES
    ('Beverages', 'Hot and cold drinks'),
    ('Pastries', 'Freshly baked pastries'),
    ('Sandwiches', 'Delicious sandwiches'),
    ('Salads', 'Fresh and healthy salads');

-- Insert sample menu items
INSERT INTO menu_items (category_id, menu_item_name, description, price)
VALUES
    (1, 'Caramel Frappuccino', 'A sweet, icy coffee treat with caramel drizzle.', 4.50),
    (1, 'Espresso', 'Rich and bold espresso.', 2.50),
    (2, 'Croissant', 'Buttery and flaky French pastry.', 3.00),
    (2, 'Muffin', 'Soft and flavorful muffin in assorted varieties.', 2.75),
    (3, 'Turkey Sandwich', 'Oven-roasted turkey with fresh veggies.', 5.00),
    (4, 'Caesar Salad', 'Crisp lettuce with Caesar dressing and croutons.', 6.50);

-- Insert sample price history
INSERT INTO price_history (menu_id, old_price, new_price)
VALUES
    (1, 4.00, 4.50),
    (2, 2.25, 2.50),
    (3, 2.80, 3.00);

-- Insert sample customers
INSERT INTO customers (customer_name, age, sex, allergens)
VALUES
    ('Alice Johnson', 28, 'female', ARRAY['nuts']),
    ('Bob Smith', 34, 'male', NULL),
    ('Charlie Brown', 22, 'other', ARRAY['dairy']),
    ('Diana Prince', 29, 'female', NULL);

-- Insert sample inventory
INSERT INTO inventory (inventory_name, quantity, unit, allergens)
VALUES
    ('Coffee Beans', 100, 'kg', NULL),
    ('Caramel Syrup', 50, 'L', NULL),
    ('Flour', 200, 'kg', ARRAY['gluten']),
    ('Lettuce', 20, 'kg', NULL);

-- Insert sample inventory transactions
INSERT INTO inventory_transactions (inventory_id, transaction_type, quantity_changed)
VALUES
    (1, 'incoming', 10),
    (2, 'outgoing', 5),
    (3, 'incoming', 15),
    (4, 'damaged', 2);

-- Insert sample menu_items_ingredients
INSERT INTO menu_items_ingredients (menu_id, inventory_id, quantity)
VALUES
    (1, 1, 0.1),
    (1, 2, 0.02),
    (2, 1, 0.05),
    (3, 3, 0.25);

-- Insert sample orders
INSERT INTO orders (customer_id, status, total_amount)
VALUES
    (1, 'created', 10.50),
    (2, 'pending', 12.75),
    (3, 'processing', 8.00);

-- Insert sample order status history
INSERT INTO order_status_history (order_id, status)
VALUES
    (1, 'created'),
    (1, 'pending'),
    (2, 'created'),
    (3, 'processing');

-- Insert sample order items
INSERT INTO order_items (order_id, menu_id, quantity, price_at_order, customization_info)
VALUES
    (1, 1, 2, 4.50, '{"size": "large"}'),
    (1, 2, 1, 2.50, '{"shots": 2}'),
    (2, 3, 3, 3.00, '{"toasted": true}'),
    (3, 4, 1, 6.50, '{"dressing": "on side"}');
