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

#### resposta

```json
{
  "Response Headers": {
    "content-type": "application/json; charset=UTF-8",
    "x-request-id": "9013ee6f-ef09-4d47-bdbc-d7aa7d52ba31",
    "date": "Sun, 21 Aug 2022 00:41:27 GMT",
    "content-length": "420"
  },
  "Response Body": {"id":1761,"long":"-46574716","lat":"-23584852","setcens":"355030893000035","are_ap":"3550308005042","district_code":"95","district_name":"JARAGUA","sub_pref_code":"29","sub_pref_name":"JARAGUA","region_05":"Norte","region_08":"Norte 1","fair_name":"JARAGUA","fair_code":"1000-1","addres_street":"RUA DOS REIS","address_number":"909.000000","address_district":"VL ZELINA","address_reference":"RUA OLIVEIRA"}
}
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

#### resposta

```json
{
    "Response Headers": {
      "content-type": "application/json; charset=UTF-8",
      "x-request-id": "1614a465-cf01-48cd-b6cf-0a4a9c1585d0",
      "date": "Sun, 21 Aug 2022 00:46:41 GMT",
      "content-length": "845"
    },
    "Response Body": [{"id":426,"long":"-46741894","lat":"-23453615","setcens":"355030842000235","are_ap":"3550308005185","district_code":"41","district_name":"JARAGUA","sub_pref_code":"2","sub_pref_name":"PIRITUBA","region_05":"Norte","region_08":"Norte 1","fair_name":"JARAGUA","fair_code":"6065-8","addres_street":"RUA LAVRINHA","address_number":"18.000000","address_district":"VL JARAGUA","address_reference":"PROXIMO A ESTACAO DO JARAGUA"},{"id":1761,"long":"-46574716","lat":"-23584852","setcens":"355030893000035","are_ap":"3550308005042","district_code":"95","district_name":"JARAGUA","sub_pref_code":"29","sub_pref_name":"JARAGUA","region_05":"Norte","region_08":"Norte 1","fair_name":"JARAGUA","fair_code":"1000-1","addres_street":"RUA JOSE DOS REIS","address_number":"909.000000","address_district":"VL ZELINA","address_reference":"RUA OLIVEIRA GOUVEIA"}]
}
```

### Excluir uma feira

```bash
curl --location --request DELETE 'http://localhost:8010/v1/fairies'
```

### Importar dados de um arquivo CSV

```bash
curl -X POST 'http://localhost:8010/v1/import_data?file=/home/file.csv'
```