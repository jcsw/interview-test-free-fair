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