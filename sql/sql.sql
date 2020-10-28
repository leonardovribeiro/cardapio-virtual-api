DROP DATABASE IF EXISTS cardapiovirtual;

CREATE DATABASE cardapiovirtual;

USE cardapiovirtual;

CREATE TABLE cliente(
  	id INT NOT NULL AUTO_INCREMENT,
	documento VARCHAR(11) NOT NULL,
  	mesa int NOT NULL,
  	update_at TIMESTAMP,
  	PRIMARY KEY(id)
);

CREATE TABLE estabelecimento(
	id INT NOT NULL AUTO_INCREMENT,
  	usuario VARCHAR(20) NOT NULL,
  	senha VARCHAR(20) NOT NULL,
  	PRIMARY KEY(id)
);

CREATE TABLE pedido(
	id int NOT NULL AUTO_INCREMENT,
  	cliente_id INT NOT NULL,
  	estabelecimento_id INT NOT NULL,
  	total DECIMAL(8,2) NOT NULL,
  	observacoes VARCHAR(140),
  	status VARCHAR(20),
  	data DATE,
  	PRIMARY KEY (id),
  	FOREIGN KEY (cliente_id) REFERENCES cliente(id),
  	FOREIGN KEY (estabelecimento_id) REFERENCES estabelecimento(id)
);

CREATE TABLE produto(
	id INT NOT NULL AUTO_INCREMENT,
  	nome VARCHAR(100) NOT NULL,
  	descricao VARCHAR(140) NOT NULL,
  	preco DECIMAL(8,2) NOT NULL,
  	PRIMARY KEY(id)
);

CREATE TABLE detalhes_pedido(
	id_produto INT NOT NULL,
  	id_pedido INT NOT NULL,
  	preco_unitario DECIMAL(8,2) NOT NULL,
  	quantidade INT NOT NULL,
  	FOREIGN KEY (id_produto) REFERENCES produto(id),
  	FOREIGN KEY (id_pedido) REFERENCES pedido(id),
  	PRIMARY KEY (id_produto, id_pedido)
);

CREATE TABLE categoria(
	id INT NOT NULL AUTO_INCREMENT,
  	nome VARCHAR(100) NOT NULL,
  	PRIMARY KEY(id)
);

CREATE TABLE categoria_produto(
	id_produto INT NOT NULL,
  	id_categoria INT NOT NULL,
  	FOREIGN KEY (id_produto) REFERENCES produto(id),
  	FOREIGN KEY (id_categoria) REFERENCES categoria(id),
  	PRIMARY KEY(id_produto, id_categoria)
);

CREATE TABLE pagamento(
	id INT NOT NULL PRIMARY KEY,
  	pedido_id INT NOT NULL,
  	tipo VARCHAR(100) NOT NULL,
  	data DATE NOT NULL
);


