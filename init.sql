CREATE SCHEMA IF NOT EXISTS wisdom;

CREATE TABLE IF NOT EXISTS wisdom.vendors (
    vendor_id VARCHAR(50) PRIMARY KEY,
    name      VARCHAR(150) NOT NULL,
    contact   VARCHAR(100),
    phone     VARCHAR(20),
    email     VARCHAR(150) UNIQUE,
    address   TEXT
);

CREATE TABLE IF NOT EXISTS wisdom.products (
    product_id VARCHAR(50) PRIMARY KEY,
    name       VARCHAR(150) NOT NULL,
    price      NUMERIC(10,2) NOT NULL,
    vendor_id  VARCHAR(50) NOT NULL,
    CONSTRAINT fk_products_vendor FOREIGN KEY (vendor_id)
        REFERENCES wisdom.vendors(vendor_id)
);

CREATE TABLE IF NOT EXISTS wisdom.services (
    service_id VARCHAR(50) PRIMARY KEY,
    name       VARCHAR(150) NOT NULL,
    price      NUMERIC(10,2) NOT NULL
);

CREATE TABLE IF NOT EXISTS wisdom.customers (
    customer_id VARCHAR(50) PRIMARY KEY,
    first_name  VARCHAR(100) NOT NULL,
    last_name   VARCHAR(100) NOT NULL,
    email       VARCHAR(150) UNIQUE NOT NULL,
    phone       VARCHAR(20),
    address     TEXT
);

INSERT INTO wisdom.customers (customer_id, first_name, last_name, email, phone, address) VALUES
('CUST001', 'Alice', 'Johnson', 'alice.johnson@example.com', '+1-202-555-0147', '123 Maple Street, Springfield, IL'),
('CUST002', 'Bob', 'Williams', 'bob.williams@example.com', '+1-202-555-0198', '456 Oak Avenue, Denver, CO'),
('CUST003', 'Clara', 'Smith', 'clara.smith@example.com', '+1-202-555-0123', '789 Pine Road, Austin, TX'),
('CUST004', 'David', 'Brown', 'david.brown@example.com', '+1-202-555-0184', '321 Birch Lane, Seattle, WA'),
('CUST005', 'Ella', 'Martinez', 'ella.martinez@example.com', '+1-202-555-0172', '654 Cedar Blvd, Miami, FL');

INSERT INTO wisdom.vendors (vendor_id, name, contact, phone, email, address) VALUES
('VEND001', 'TechWorld Inc.', 'John Carter', '+1-202-555-0111', 'contact@techworld.com', '101 Silicon Ave, San Francisco, CA'),
('VEND002', 'HomeGoods Co.', 'Mary Adams', '+1-202-555-0222', 'sales@homegoods.com', '202 Market St, New York, NY'),
('VEND003', 'SportsPlus', 'Robert Lee', '+1-202-555-0333', 'support@sportsplus.com', '303 Stadium Blvd, Chicago, IL'),
('VEND004', 'FashionHub', 'Linda Kim', '+1-202-555-0444', 'info@fashionhub.com', '404 Style Rd, Los Angeles, CA');

INSERT INTO wisdom.products (product_id, name, price, vendor_id) VALUES
('PROD001', 'Laptop X200', 1200.00, 'VEND001'),
('PROD002', 'Wireless Mouse', 25.50, 'VEND001'),
('PROD003', 'Sofa Set', 899.99, 'VEND002'),
('PROD004', 'Running Shoes', 120.00, 'VEND003'),
('PROD005', 'Designer Handbag', 450.00, 'VEND004');

INSERT INTO wisdom.services (service_id, name, price) VALUES
('SERV001', 'Computer Repair', 150.00),
('SERV002', 'Home Cleaning', 80.00),
('SERV003', 'Personal Training', 60.00),
('SERV004', 'Clothing Alteration', 40.00);