
## Compilação e Execução

### Executar diretamente do código-fonte

1. **Compilar o código-fonte:**

   ```bash
   go build -o loadtester ./cmd/loadtester
   ```

2. **Executar o teste de carga:**

   Substitua os valores de `--url`, `--requests` e `--concurrency` de acordo com suas necessidades.

   ```bash
   ./loadtester --url=http://google.com --requests=100 --concurrency=10
   ```

### Executar via Docker

1. **Construir a imagem Docker:**

   ```bash
   docker build -t goexpert-stress-test .
   ```

2. **Executar o teste de carga via Docker:**

   Substitua os valores de `--url`, `--requests` e `--concurrency` de acordo com suas necessidades.

   ```bash
   docker run goexpert-stress-test --url=http://google.com --requests=100 --concurrency=10
   ```

## Resultado do Teste

Após a execução, o sistema apresentará um relatório detalhado que inclui:

- Tempo total gasto na execução
- Total de requests realizados
- Quantidade de requests com status HTTP 200
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).
