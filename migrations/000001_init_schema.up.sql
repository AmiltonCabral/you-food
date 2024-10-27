CREATE TABLE users (
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  order_code INTEGER NOT NULL,
  password VARCHAR(32) NOT NULL,
  address VARCHAR(255) NOT NULL
);

CREATE TABLE stores (
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  password VARCHAR(32) NOT NULL,
  address VARCHAR(255) NOT NULL
);

CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  store_id VARCHAR(32),
  name VARCHAR(100) NOT NULL,
  description VARCHAR(255),
  price DECIMAL(10,2) NOT NULL,
  ammount INT DEFAULT 0,
  FOREIGN KEY (store_id) REFERENCES stores(id) ON DELETE CASCADE
);

CREATE TABLE delivery_man (
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  password VARCHAR(32) NOT NULL
);

CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  user_id VARCHAR(32) NOT NULL,
  product_id INT NOT NULL,
  quantity INT NOT NULL,
  total_price DECIMAL(10, 2) NOT NULL,
  status VARCHAR(255) DEFAULT 'pending',
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);
