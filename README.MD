# Cardápio Virtual API
Respositório da API responsável por servir o frontend do projeto Cardápio Virtual.

## Recursos Utilizados
- **[Visual Studio Code](https://code.visualstudio.com/)**
- **[Golang](https://golang.org/)**
- **[Diagrams](https://www.diagrams.net/)**
- **[Docker](https://docs.docker.com/get-docker/)**

## Executando a Aplicação Localmente
Primeiramente deve-se ter o banco de dados MariaDB instalado e configurado na máquina, abaixo você confere uma maneira de rodar o banco de dados usando o Docker. 

1. Faça o build da imagem usando o comando abaixo dentro da pasta mariadb:
    ```
    $ docker build -t cardapiovirtual-db .
    ```
2. Execute o container usando o seguinte comando:
    ```
    $ docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=senharoot -e MYSQL_DATABASE=Cardapiovirtual -e MYSQL_USER=usuario -e MYSQL_PASSWORD=senhamysql cardapiovirtual-db
    ```
3. Conecte-se ao banco de dados usando um client ([Beekeeper Studio](https://www.beekeeperstudio.io/)) ou usando o comando:
    ```
    $ docker exec -it container_id bash
    root@container_id:/# mysql -u root -p
    Enter password: senharoot
    ```
4. Execute o comando abaixo dentro da pasta api:
    ```
    go run main.go
    ```
## Links e Recursos Adicionais
- **[Aprenda Go](https://www.youtube.com/watch?v=WiGU_ZB-u0w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg)**
- **[Aprenda Golang do Zero](https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/)**
