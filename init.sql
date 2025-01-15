CREATE DATABASE frappuccino_db
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LOCALE_PROVIDER = 'libc'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

ALTER USER postgres WITH PASSWORD 'latte';

-- enums
CREATE TYPE sex AS ENUM ('female', 'male', 'other');

CREATE TYPE transaction_type AS ENUM
    ('subtract', 'add', 'zero');

CREATE TYPE order_status AS ENUM ('created', 'pending', 'processing', 'completed', 'canceled', 'rejected');

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
    allergens VARCHAR(500)[],
    is_active BOOL DEFAULT TRUE
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
-- CREATE INDEX idx_inventory_transactions_transaction_date ON inventory_transactions (transaction_date);
CREATE INDEX idx_order_items_customization ON order_items (customization_info);
CREATE INDEX idx_order_status_history_change_date ON order_status_history (change_date);
CREATE INDEX idx_order_created_at ON orders (created_at);



-- Mock data for categories
INSERT INTO categories (category_name, description) VALUES
('Coffee', 'Freshly brewed coffee beverages'),
('Tea', 'Variety of teas from around the world'),
('Snacks', 'Delicious snacks to go with your drink'),
('Desserts', 'Sweet treats to indulge in'),
('Cold Beverages', 'Chilled drinks for hot days');

-- Mock data for menu_items
INSERT INTO menu_items
    (category_id, menu_item_name, description, price) VALUES
(1, 'Espresso', 'Classic espresso shot', 2.5),
(1, 'Latte', 'Espresso with steamed milk', 3.5),
(1, 'Cappuccino', 'Espresso with frothed milk', 3.7),
(2, 'Green Tea', 'Organic green tea', 2.0),
(2, 'Chai Latte', 'Spiced tea with milk', 3.0),
(3, 'Blueberry Muffin', 'Freshly baked muffin', 2.8),
(3, 'Chocolate Cookie', 'Large chocolate chip cookie', 1.5),
(4, 'Cheesecake', 'Rich and creamy cheesecake', 4.0),
(5, 'Iced Coffee', 'Chilled coffee beverage', 3.2),
(5, 'Lemonade', 'Refreshing lemon drink', 2.5);

-- Mock data for price_history
INSERT INTO price_history
    (menu_id, old_price, new_price, change_date) VALUES
(1, 2.3, 2.5, '2024-01-01'),
(2, 3.0, 3.5, '2024-02-01'),
(3, 3.5, 3.7, '2024-03-01'),
(4, 1.8, 2.0, '2024-01-15'),
(5, 2.8, 3.0, '2024-02-10'),
(6, 2.5, 2.8, '2024-03-20'),
(7, 1.3, 1.5, '2024-02-05'),
(8, 3.8, 4.0, '2024-03-10'),
(9, 3.0, 3.2, '2024-04-01'),
(10, 2.3, 2.5, '2024-01-25');

-- Mock data for customers
INSERT INTO customers
    (customer_name, age, sex, allergens) VALUES
('Alice Johnson', 28, 'female', '{"nuts"}'),
('Bob Smith', 35, 'male', '{}'),
('Charlie Brown', 22, 'other', '{"gluten"}'),
('Diana Prince', 30, 'female', '{}'),
('Edward Clark', 40, 'male', '{}'),
('Fiona Adams', 27, 'female', '{"dairy"}'),
('George Hill', 33, 'male', '{}'),
('Hannah Moore', 25, 'female', '{"soy"}'),
('Ian Scott', 29, 'male', '{}'),
('Jane Doe', 31, 'female', '{}');

-- Mock data for inventory
INSERT INTO inventory
    (inventory_name, quantity, unit, allergens) VALUES
('Coffee Beans', 100, 'kg', '{}'),
('Milk', 200, 'liters', '{"dairy"}'),
('Sugar', 50, 'kg', '{}'),
('Flour', 80, 'kg', '{"gluten"}'),
('Eggs', 300, 'units', '{}'),
('Butter', 100, 'kg', '{"dairy"}'),
('Tea Leaves', 40, 'kg', '{}'),
('Chocolate Chips', 60, 'kg', '{}'),
('Lemons', 30, 'kg', '{}'),
('Blueberries', 20, 'kg', '{}'),
('Vanilla Syrup', 25, 'liters', '{}'),
('Honey', 15, 'kg', '{}'),
('Cinnamon', 10, 'kg', '{}'),
('Baking Powder', 12, 'kg', '{}'),
('Ice Cubes', 500, 'kg', '{}'),
('Straws', 1000, 'units', '{}'),
('Cups', 2000, 'units', '{}'),
('Napkins', 1500, 'units', '{}'),
('Trays', 500, 'units', '{}'),
('Plates', 300, 'units', '{}');

-- Mock data for inventory_transactions
-- INSERT INTO inventory_transactions
--     (inventory_id, transaction_type, quantity_changed, transaction_date) VALUES
-- (1, 'add', 20, '2024-01-10'),
-- (2, 'subtract', 50, '2024-01-15'),
-- (3, 'add', 10, '2024-02-01'),
-- (4, 'subtract', 20, '2024-02-15'),
-- (5, 'add', 30, '2024-03-01'),
-- (6, 'subtract', 15, '2024-03-20'),
-- (7, 'add', 5, '2024-04-01'),
-- (8, 'subtract', 10, '2024-04-10'),
-- (9, 'add', 15, '2024-05-01'),
-- (10, 'subtract', 5, '2024-05-10');

-- Mock data for orders
INSERT INTO orders
    (customer_id, created_at, status, total_amount) VALUES
(1, '2024-01-10', 'completed', 15.0),
(2, '2024-01-15', 'canceled', 20.0),
(3, '2024-02-01', 'completed', 12.0),
(4, '2024-02-15', 'processing', 18.0),
(5, '2024-03-01', 'pending', 22.0),
(6, '2024-03-20', 'completed', 30.0),
(7, '2024-04-01', 'rejected', 25.0),
(8, '2024-04-10', 'created', 28.0),
(9, '2024-05-01', 'completed', 10.0),
(10, '2024-05-10', 'processing', 40.0);

-- Mock data for order_status_history
INSERT INTO order_status_history
    (order_id, status, change_date) VALUES
(1, 'created', '2024-01-10'),
(1, 'completed', '2024-01-11'),
(2, 'created', '2024-01-15'),
(2, 'canceled', '2024-01-16'),
(3, 'created', '2024-02-01'),
(3, 'completed', '2024-02-02'),
(4, 'created', '2024-02-15'),
(4, 'processing', '2024-02-16'),
(5, 'created', '2024-03-01'),
(5, 'pending', '2024-03-02');