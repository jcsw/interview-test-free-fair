## Serviços

### Cadastrar uma nova feira

```bash
curl --location --request PUT 'http://localhost:8010/v1/fairies?fair_code=4041-0' \
--header 'Content-Type: application/json' \
--data-raw '{
        "long": "-46574716",
        "lat": "-23584852",
        "setcens": "355030893000035",
        "are_ap": "3550308005042",
        "district_code": "95",
        "district_name": "JARAGUA",
        "sub_pref_code": "29",
        "sub_pref_name": "JARAGUA",
        "region_05": "Norte",
        "region_08": "Norte 1",
        "fair_name": "JARAGUA",
        "fair_code": "1000-1",
        "addres_street": "RUA JOSE DOS REIS",
        "address_number": "909.000000",
        "address_district": "VL ZELINA",
        "address_reference": "RUA OLIVEIRA GOUVEIA"
    }'
```

### Atualizar uma feira

```bash
curl --location --request POST 'http://localhost:8010/v1/fairies' \
--header 'Content-Type: application/json' \
--data-raw '{
        "long": "-46574716",
        "lat": "-23584852",
        "setcens": "355030893000035",
        "are_ap": "3550308005042",
        "district_code": "95",
        "district_name": "JARAGUA",
        "sub_pref_code": "29",
        "sub_pref_name": "JARAGUA",
        "region_05": "Leste",
        "region_08": "Leste 1",
        "fair_name": "JARAGUA",
        "fair_code": "1000-1",
        "addres_street": "RUA JOSE DOS REIS",
        "address_number": "909.000000",
        "address_district": "JARAGUA",
        "address_reference": "RUA OLIVEIRA GOUVEIA"
    }'
```

### Consultar uma feira

```bash
curl --location --request GET 'http://localhost:8010/v1/fairies?fair_code=1000-1'
```

### Buscar feiras

```bash
curl --location --request GET 'http://localhost:8010/v1/fairies/search?fair_name=JARAGUA'
```

```bash
curl --location --request GET 'http://localhost:8010/v1/fairies/search?district_name=VILA PRUDENTE'
```

```bash
curl --location --request GET 'http://localhost:8010/v1/fairies/search?district_name=VILA PRUDENTE'
```

```bash
curl --location --request GET 'http://localhost:8010/v1/fairies/search?address_district=VL ZELINA'
```

### Excluir uma feira

```bash
curl --location --request DELETE 'http://localhost:8010/v1/fairies'
```

### Importar dados de um arquivo CSV

```bash
curl -X POST 'http://localhost:8010/v1/import_data?file=/home/file.csv'
```