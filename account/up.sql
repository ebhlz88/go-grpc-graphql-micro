CREATE TABLE IF NOT EXISTS accounts (
	id CHAR(27) PRIMARY KEY
	name VARCHAR(24) NOT NULL
);

CREATE TABLE IF NOT EXISTS order_products(
	order_id CHAR(27) REFERENCES orders (id) ON DELETE CASCADE,
	product_id CHAR(27),
	quantity INT NOT NULL,
	PRIMARY KEY (product_id, order_id)
);