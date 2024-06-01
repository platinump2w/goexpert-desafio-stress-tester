# Desafio - CLI para execução de Stress Tests
Criar um sistema CLI em Go para realizar testes de carga em um serviço web

## Descrição
Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

## Requisitos Gerais
- O CLI deve aceitar os parâmetros `--url`, `--requests` e `--concurrency`;
- Garantir que o número total de requests seja seja cumprido;
- Apresentar um relatório com os seguintes dados:
  - Tempo total gasto na execução;
  - Quantidade total de requests realizados;
  - Quantidade de requests com status HTTP 200;
  - Distribuição de outros códigos de status HTTP (como 404, 500, etc.);

## Execução

**Build da imagem Docker**

- Na pasta raíz do projeto, executar o comando:
```bash
docker build -t platinump2w/goexpert-stress-tester:latest .
```

**Execução do CLI**

- Acessar a pasta `cmd/ordersystem` e executar o comando:
```bash
docker run platinump2w/goexpert-stress-tester:latest [flags]
```

**Exemplo de Execução**
```bash
# https://httpstat.us/random retorna um status HTTP aleatório entre 200, 201, 500 e 504
docker run platinump2w/goexpert-stress-tester:latest --url=https://httpstat.us/random/200,201,500-504 --requests=18 --concurrency=5
```
