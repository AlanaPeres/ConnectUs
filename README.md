# :globe_with_meridians: ConnectUs API: Uma Rede Social Interativa
A ConnectUs API é uma plataforma de rede social que oferece recursos poderosos para criar, interagir e compartilhar informações. A API é projetada para permitir a criação, gerenciamento e consulta de usuários e publicações, oferecendo uma experiência rica e envolvente para a sua aplicação.

## :pushpin: Entidades Principais 
### Usuários :family:
A entidade "Usuários" é o núcleo da ConnectUs API, fornecendo funcionalidades abrangentes de gerenciamento de perfis de usuário. Os recursos incluem:

<li>CRUD Básico: Crie, leia, atualize e delete perfis de usuários.</li>
<li>Seguir Usuários: Estabeleça conexões ao seguir outros usuários.</li>
<li>Parar de Seguir: Termine conexões ao deixar de seguir usuários.</li>
<li>Listar Seguindo: Obtenha a lista de usuários que um usuário específico está seguindo.</li>
<li>Listar Seguidores: Recupere a lista de usuários que seguem um determinado usuário.</li>
<li>Atualizar Senha: Permita aos usuários atualizar suas senhas com segurança.</li>

### Publicações :newspaper:
A entidade "Publicações" permite que os usuários compartilhem informações com sua rede. Os recursos incluem:

<li>CRUD Básico: Crie, leia, atualize e delete publicações.</li>
<li>Listar por Usuários Seguidos: Recupere publicações de usuários que estão sendo seguidos por um determinado usuário.</li>
<li>Curtir Publicações: Permita aos usuários curtir as publicações de outros usuários.</li>

## :wrench: Estrutura da API
A ConnectUs API é estruturada em pacotes principais e auxiliares para promover modularidade e facilitar a manutenção.

### :file_folder: Pacotes Principais

<li>Main: Configuração inicial da aplicação, inicialização do roteador HTTP e execução do servidor.</li>
<li>Router: Gerenciamento de rotas usando Gorilla Mux, permitindo o direcionamento de solicitações para os controladores adequados.</li>
<li>Controllers: Implementação das funções de manipulação de solicitações HTTP. Os controladores se comunicam com os modelos e repositórios.</li>
<li>Models: Definição das estruturas de dados para Usuários e Publicações, incluindo métodos para operações específicas.</li>
<li>Repositórios: Interação direta com o banco de dados para executar operações de leitura, escrita e consulta.</li>

### :file_folder: Pacotes Auxiliares

<li>Config: Gerenciamento de variáveis de ambiente para proteger informações sensíveis e garantir facilidade de manutenção.</li>
<li>Banco: Gerenciamento da conexão com o banco de dados MySQL.</li>
<li>Autenticação: Lida com a autenticação de usuários, incluindo a geração e validação de tokens.</li>
<li>Middleware: Camada intermediária que verifica a autenticação do usuário antes de encaminhar as solicitações para os controladores.</li>
<li>Segurança: Gerenciamento de senhas, incluindo a verificação da correspondência de senhas no banco de dados.</li>
<li>Respostas: Padronização das respostas da API para garantir consistência nas interações.</li>
