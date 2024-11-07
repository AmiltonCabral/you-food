TRUNCATE users, stores, delivery_man, products, orders CASCADE;

INSERT INTO users (id, name, password, order_code, address) VALUES
('user1', 'João Silva', 'pass123', 1234, 'Rua A, 123'),
('user2', 'Maria Santos', 'pass456', 2345, 'Av B, 456'),
('user3', 'Pedro Lima', 'pass789', 3456, 'Rua C, 789'),
('user4', 'Ana Oliveira', 'pass321', 4567, 'Av D, 321'),
('user5', 'Carlos Souza', 'pass654', 5678, 'Rua E, 654');

INSERT INTO stores (id, name, password, address) VALUES
('store1', 'Pizzaria Bella', 'store123', 'Rua das Pizzas, 100'),
('store2', 'Hamburgueria Top', 'store456', 'Av dos Hambúrgueres, 200'),
('store3', 'Comida Japonesa Express', 'store789', 'Rua do Sushi, 300'),
('store4', 'Padaria Delícia', 'store321', 'Av dos Pães, 400'),
('store5', 'Restaurante Mineiro', 'store654', 'Rua da Comida Caseira, 500');

INSERT INTO delivery_man (id, name, password) VALUES
('delivery1', 'José Entregador', 'delivery123'),
('delivery2', 'Paulo Motoboy', 'delivery456'),
('delivery3', 'Roberto Entregas', 'delivery789'),
('delivery4', 'Marcos Delivery', 'delivery321'),
('delivery5', 'Lucas Express', 'delivery654');

-- INSERT INTO products (store_id, name, description, price, ammount) VALUES
-- ('store1', 'Pizza Margherita', 'Pizza tradicional italiana com molho de tomate, muçarela e manjericão', 45.90, 20),
-- ('store1', 'Pizza Pepperoni', 'Pizza com pepperoni e muçarela', 49.90, 15),
-- ('store2', 'Hambúrguer Clássico', 'Hambúrguer com queijo, alface e tomate', 32.90, 30),
-- ('store2', 'Hambúrguer Especial', 'Hambúrguer duplo com bacon e cheddar', 39.90, 25),
-- ('store3', 'Combo Sushi', '20 peças variadas de sushi', 89.90, 10),
-- ('store3', 'Hot Roll', '10 peças de hot roll', 45.90, 15),
-- ('store4', 'Pão Francês', 'Pão francês fresquinho', 0.75, 100),
-- ('store4', 'Bolo de Chocolate', 'Bolo caseiro de chocolate', 35.90, 8),
-- ('store5', 'Feijoada Completa', 'Feijoada com acompanhamentos', 55.90, 12),
-- ('store5', 'Frango com Quiabo', 'Prato típico mineiro', 45.90, 10),
-- ('store5', 'Frango de Batata', 'Prato com batata doce', 55.90, 15);



-- INSERT INTO orders (user_id, product_id, quantity, total_price, status) VALUES
-- ('user1', 1, 1, 45.90, 'entregue'),
-- ('user2', 3, 2, 65.80, 'em preparo'),
-- ('user3', 5, 1, 89.90, 'em entrega'),
-- ('user4', 7, 10, 7.50, 'confirmado'),
-- ('user5', 9, 1, 55.90, 'pendente'),
-- ('user1', 2, 1, 49.90, 'confirmado'),
-- ('user2', 4, 1, 39.90, 'em preparo'),
-- ('user3', 6, 2, 91.80, 'pendente'),
-- ('user4', 8, 1, 35.90, 'entregue'),
-- ('user5', 10, 2, 91.80, 'em entrega');
