package fair

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	sys "interview-test-free-fair/pkg/sys"
)

// SearchHandler function to route GET /v1/fairies
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	fairName := r.URL.Query().Get("fair_name")
	districtName := r.URL.Query().Get("district_name")
	region05 := r.URL.Query().Get("region_05")
	addressDistric := r.URL.Query().Get("address_district")

	searchParameter := searchParameter{fairName, districtName, region05, addressDistric}

	result := search(searchParameter)

	sys.HTTPResponseWithJSON(w, 200, result)
}

// CreateHandler function to route POST /v1/fairies
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var newFair FreeFair

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sys.HTTPResponseWithJSON(w, 400, err)
		return
	}

	err = json.Unmarshal(body, &newFair)
	if err != nil {
		sys.HTTPResponseWithJSON(w, 400, err)
		return
	}

	save(&newFair)

	sys.HTTPResponseWithJSON(w, 201, newFair)
}

// UpdateHandler function to route PUT /v1/fairies
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	sys.HTTPResponseWithJSON(w, 201, nil)
}

// ImportDataHandler function to handle POST /v1/import_data
func ImportDataHandler(w http.ResponseWriter, r *http.Request) {
	data := sys.LoadCsvData("data/DEINFO_AB_FEIRASLIVRES_2014.csv")

	if data == nil {
		return
	}

	for i, line := range data {
		if i == 0 {
			continue
		}

		saveFromCsvData(line)
	}
}
