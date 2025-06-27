// customer ddl

CREATE SEQUENCE customer_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- Table for Customer
CREATE TABLE customer (
    id BIGINT PRIMARY KEY DEFAULT nextval('customer_id_seq'),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    street VARCHAR(255) NOT NULL,
    house_number VARCHAR(50) NOT NULL,
    zip_code VARCHAR(20) NOT NULL
);


// customer order and order line
-- Sequence for Order ID
CREATE SEQUENCE customer_order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- Sequence for OrderLine ID
CREATE SEQUENCE order_line_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- Table: customer_order
CREATE TABLE customer_order (
    id BIGINT PRIMARY KEY DEFAULT nextval('customer_order_id_seq'),
    reference VARCHAR(255),
    total_amount NUMERIC(12, 2),
    payment_method VARCHAR(100),
    customer_id VARCHAR(100),  -- Assuming customer_id is a string (could be UUID or something else)
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_modified_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Table: order_line
CREATE TABLE order_line (
    id BIGINT PRIMARY KEY DEFAULT nextval('order_line_id_seq'),
    product_id BIGINT NOT NULL,
    quantity NUMERIC(10, 2),
    order_id BIGINT NOT NULL,
    CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES customer_order(id) ON DELETE CASCADE
);
....................

CREATE SEQUENCE payment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- Table: payment
CREATE TABLE payment (
    id BIGINT PRIMARY KEY DEFAULT nextval('payment_id_seq'),
    amount NUMERIC(12, 2) NOT NULL,
    payment_method VARCHAR(100) NOT NULL,
    order_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_payment_order FOREIGN KEY (order_id) REFERENCES customer_order(id) ON DELETE CASCADE
);

.............................
// for payment ddl

-- Sequence for Category ID
CREATE SEQUENCE category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- Table: category
CREATE TABLE category (
    id BIGINT PRIMARY KEY DEFAULT nextval('category_id_seq'),
    name VARCHAR(255) NOT NULL,
    description TEXT
);


// product and product category

-- Sequence for Product ID
CREATE SEQUENCE product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- Table: product
CREATE TABLE product (
    id BIGINT PRIMARY KEY DEFAULT nextval('product_id_seq'),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    available_quantity NUMERIC(12, 2) NOT NULL DEFAULT 0,
    price NUMERIC(12, 2) NOT NULL,
    category_id BIGINT NOT NULL,
    CONSTRAINT fk_product_category FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE
);




// add categoy sql


INSERT INTO category (name, description) VALUES
('Books', 'Various kinds of books'),
('Clothing', 'Men and Women fashion'),
('Home Appliances', 'Kitchen and other home appliances'),
('Electronics', 'Devices and gadgets');