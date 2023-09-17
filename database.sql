CREATE TABLE products (
	product_id serial PRIMARY KEY,
	name VARCHAR ( 200 ) UNIQUE NOT NULL,
	description TEXT,
  rating FLOAT
);

CREATE TABLE product_varieties (
	variety_id serial PRIMARY KEY,
  product_id INT REFERENCES products(product_id) ON DELETE CASCADE,
	variant VARCHAR ( 100 ),
  stock INT,
  price DECIMAL(10, 2) NOT NULL
);

INSERT INTO products (name, description, rating) VALUES (
  'Kemeja', 'Kemeja beraneka warna', '5.0'
);

INSERT INTO product_varieties (product_id, variant, stock, price) VALUES 
(1, 'Merah', 10, 20000), (1, 'Kuning', 5, 21000), (1, 'Hijau', 3, 22000);
