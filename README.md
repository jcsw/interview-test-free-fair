# Projeto de teste para entrevista (feira livre)

## Pré-requisitos

`Go 1.18.x`
`Make`
`Docker`
`Docker Compose`

## Instalação

```bash
make install
```

## Realizando o build

```bash
make build
```

## Rodando testes únitarios

```bash
make test
```

## Rodando testes de integração

```bash
make itest
```

### Cobertura de teste

Os arquivos de cobertura teste são gerados na pasta ./build 

## Rodando o serviço

```bash
make deps-start
make run
make deps-stop
```

### Verificando se o serviço está rodando

```bash
curl -X GET 'http://localhost:8010/ping'
```

### Métricas (Prometheus)

```bash
curl -X GET 'http://localhost:8010/metrics'
```

### Logs

Os arquivos de logs são gerados na pasta ./logs no formato YYYY-MM-DD.log

### Documentação dos serviços

Na pasta ./docs estão os arquivos de documentação do serviço e postman