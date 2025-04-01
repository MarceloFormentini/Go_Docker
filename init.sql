CREATE TABLE IF NOT EXISTS produto (
	id SERIAL PRIMARY KEY,
	nome VARCHAR(255) NOT NULL,
	descricao TEXT NOT NULL,
	preco NUMERIC(10, 2) NOT NULL
);

INSERT INTO produto (nome, descricao, preco) VALUES
('Produto 1', 'Descrição do Produto 1', 10.50),
('Produto 2', 'Descrição do Produto 2', 20.00),
('Produto 3', 'Descrição do Produto 3', 30.75),
('Produto 4', 'Descrição do Produto 4', 40.99),
('Produto 5', 'Descrição do Produto 5', 50.25);