# DevBook API

DevBook API é uma aplicação que fornece uma interface para gerenciar usuários, posts e autenticação. Esta API é construída utilizando Go e segue uma arquitetura modular para facilitar a manutenção e a escalabilidade.

## Funcionalidades

- Autenticação de usuários
- Gerenciamento de usuários
- Criação, leitura, atualização e exclusão de posts

## Estrutura do Projeto

- **.vscode/**: Configurações do Visual Studio Code
- **sql/**: Scripts SQL para criação e configuração do banco de dados
- **src/**: Código-fonte do projeto
  - **authentication/**: Lógica de autenticação
  - **config/**: Configurações da aplicação
  - **controllers/**: Controladores para lidar com requisições HTTP
  - **database/**: Configuração e conexão com o banco de dados
  - **middlewares/**: Middlewares para processamento de requisições
  - **models/**: Definições de modelos de dados
  - **repositories/**: Repositórios para interação com o banco de dados
  - **responses/**: Formatação de respostas HTTP
  - **router/**: Roteamento de requisições
  - **security/**: Lógica de segurança

## Pré-requisitos

- Go 1.16 ou superior
- PostgreSQL

## Instalação

1. Clone o repositório:

    ```sh
    git clone https://github.com/seu-usuario/devbook-api.git
    cd devbook-api
    ```

2. Crie o banco de dados utilizando o script SQL fornecido:

    ```sh
    psql -U seu-usuario -d seu-banco-de-dados -f sql/sql.sql
    ```

3. Configure as variáveis de ambiente necessárias no arquivo [.env](http://_vscodecontentref_/0):

    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=seu-usuario
    DB_PASSWORD=sua-senha
    DB_NAME=seu-banco-de-dados
    JWT_SECRET=sua-chave-secreta
    ```

4. Instale as dependências:

    ```sh
    go mod tidy
    ```

## Execução

Para executar a aplicação, utilize o comando:

```sh
go run main.go
```