# ConnectUs

### O que é ?

É uma rede social onde é possível criar publicações de textos.

A aplicação ConnectUs possui duas entidades: Usuários e Publicações.

A entidade **usuários** possui o CRUD básico e outras funcionalidades como:

- Seguir outro usuário
- Parar de seguir
- Buscar todos os usuários que segue
- Buscar todos os usuários que são seguidos
- Atualizar senha

No banco de dados Mysql eu tenho duas tabelas para lidar com esses usuários:

- Usuários
- Seguidores

A segunda entidade são as publicações:

- CRUD básico (criar publicação, buscar, editar e excluir)
- Buscar publicações de acordo com os usuários que são seguidos por um outro usuário.
- Curtir

No banco de dados para a entidade Publicações terei apenas a tabela de publicações.

# Estrutura API ConnectUs

ConnectUs é uma aplicação web (front-end) que chama a API ConnectUs,  e ela vai ser o meio de comunicação entre a aplicação web e o banco de dados. Então o banco de dados só será manipulado pela API.

Como os programas em GO são divididos em pacotes, veremos como é a estrutura desses pacotes na ConnectUs API.

## Pacotes Principais

Esses pacotes principais são os que estão relacionados com a estrutura da API

- Main
- Router - configura o router e todas as rotas que estão em baixo dele utilizando o Gorilla mux que é um roteador  http para Golang.
- Controllers - é onde ficam todas as funções que irão lidar e executar as requisições http configuradas no pacote Router. O pacote controller vai se comunicar com os pacotes Modelos e Repositórios.
- Models - é aqui onde vão ficar guardado os modelos(structs e seus métodos) de Usuários e Publicações, que são as duas entidades da aplicação.
- Repositórios - Nesse pacote temos a interação com o banco de dados, então as interações do banco só são manipuladas através deste pacote.

## Pacotes Auxiliares

Lidam mais com as utilidades da API.

- Config - vai principalmente configurar as variáveis de ambiente para proteger as informações dos usuários(tirando as credenciais do próprio código) e para que a API tenha uma fácil manutenção se um dia for necessário.
- Banco - Abre a conexão com o banco de dados
- Autenticação - Cuida do Login e criação de Token de autenticação.
- Middleware - Camada entre a requisição e a resposta para ver se o usuário está autenticado, então ele é utilizado em conjunto com o pacote de autenticação.
- Segurança - Lida com as senhas, verifica se a senha que o usuários está informando bate com as informações do banco de dados.
- Respostas - Padronizar as respostas que a API devolve.

