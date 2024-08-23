## Teste Técnico: Sistema de Reserva de Passagens

### Objetivo
Desenvolver uma API REST para um sistema de reserva de passagens de ônibus, que permita a criação, gerenciamento e busca de rotas, viagens, e reservas de assentos. O sistema deve permitir que os usuários busquem por horários de viagens, façam reservas e visualizem suas passagens.

### Requisitos Funcionais

#### Gerenciamento de Rotas
- **Endpoint**: Criar, atualizar, listar e excluir rotas.
- **Cada rota deve conter**:
  - ID da rota
  - Nome da rota (ex.: "São Paulo - Rio de Janeiro")
  - Origem (cidade de partida)
  - Destino (cidade de chegada)

#### Gerenciamento de Viagens
- **Endpoint**: Criar, atualizar, listar e excluir viagens.
- **Cada viagem deve estar associada a uma rota e deve conter**:
  - ID da viagem
  - ID da rota
  - Data e hora de partida
  - Data e hora de chegada
  - Preço da passagem
  - Número total de assentos

#### Sistema de Reserva
- **Endpoint**: Reservar um assento em uma viagem.
- **Cada reserva deve conter**:
  - ID da reserva
  - ID da viagem
  - Nome do passageiro
  - Número do assento
- **Implementar a lógica para**: Garantir que o mesmo assento não seja reservado mais de uma vez para a mesma viagem.
- **Endpoint para**: Cancelar uma reserva.

#### Busca de Viagens
- **Endpoint**: Buscar viagens com base na origem, destino e data.
- **O sistema deve retornar**: Todas as viagens disponíveis que correspondem aos critérios de busca, mostrando o número de assentos disponíveis e o preço.

#### Gerenciamento de Usuários (opcional)
- **Endpoint**: Criar e autenticar usuários.
- **Implementar um sistema de autenticação simples** (por exemplo, com JWT) para proteger as rotas de reserva e cancelamento.

### Requisitos Não Funcionais

#### Banco de Dados
- **Usar**: PostgreSQL para armazenar os dados do sistema.
- **Usar `sqlc`**: Para gerar os métodos de acesso ao banco de dados.

#### Framework Web
- **Usar**: Echo para criar as rotas e endpoints da API.

#### Documentação
- **A API deve ser documentada**: Usando Swagger ou outra ferramenta de documentação de APIs.
- **Incluir exemplos de uso**: Para cada endpoint.

#### Testes
- **Escrever testes unitários**: Para a lógica de negócio e testes de integração para os endpoints da API.
- **Testar casos de sucesso e falha**: (ex.: tentativa de reservar um assento já reservado).

#### Docker
- **Configurar o projeto para rodar**: Em containers usando Docker e Docker Compose.
- **O ambiente Docker deve incluir**: O serviço da API e um container PostgreSQL.

### Entregáveis
- Código fonte do projeto em um repositório Git.
- Arquivo `docker-compose.yml` para configurar a aplicação e o banco de dados.
- Script SQL para criar e popular as tabelas do banco de dados (ou usar migrations).
- Documentação da API com exemplos de chamadas.
- Testes unitários e de integração.

### Critérios de Avaliação
- **Qualidade do código**: Estrutura e organização do código, uso de boas práticas, e clareza.
- **Funcionalidade**: A API deve atender a todos os requisitos funcionais descritos.
- **Testes**: Cobertura e qualidade dos testes escritos.
- **Documentação**: Completude e clareza da documentação.
- **Dockerização**: Facilidade de uso e configuração do ambiente Docker.

---

Essa proposta abrange diferentes aspectos do desenvolvimento, desde o design da API até a implementação de lógica de negócio e a escrita de testes. Isso permite avaliar tanto habilidades técnicas quanto a capacidade de seguir boas práticas de engenharia de software.
