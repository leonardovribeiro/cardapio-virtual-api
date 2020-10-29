# Derivando da imagem oficial do MariaDb
FROM mariadb:10.5.6-focal

# Adicionando os scripts SQL para serem executados na criação do banco
COPY ./sql/ /docker-entrypoint-initdb.d/
