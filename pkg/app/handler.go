package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	sys "interview-test-free-fair/pkg/sys"
)

type FairHandler struct {
	Service FairService
}

// Search function to route GET /v1/fairies
func (h *FairHandler) Search(w http.ResponseWriter, r *http.Request) {
	fairName := r.URL.Query().Get("fair_name")
	districtName := r.URL.Query().Get("district_name")
	region05 := r.URL.Query().Get("region_05")
	addressDistric := r.URL.Query().Get("address_district")

	searchParameter := FairSearchParameter{fairName, districtName, region05, addressDistric}

	result, err := h.Service.Fairs(searchParameter)
	if err != nil {
		sys.HTTPResponseWithJSON(w, 500, err)
		return
	}

	sys.HTTPResponseWithJSON(w, 200, result)
}

// Find function to route GET /v1/fairies
func (h *FairHandler) Find(w http.ResponseWriter, r *http.Request) {
	fairCode := r.URL.Query().Get("fair_code")
	if len(fairCode) == 0 {
		sys.HTTPResponseWithJSON(w, 400, "parameter 'fair_code' is required")
		return
	}

	fair, err := h.Service.Fair(fairCode)
	if err != nil {
		sys.HTTPResponseWithJSON(w, 500, err)
		return
	}

	if fair == nil {
		sys.HTTPResponseWithJSON(w, 404, "fair not found")
		return
	}

	sys.HTTPResponseWithJSON(w, 200, fair)
}

// Create function to route POST /v1/fairies
func (h *FairHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newFair Fair

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

	err = h.Service.CreateFair(&newFair)
	if err != nil {
		sys.HTTPResponseWithJSON(w, 500, err)
		return
	}

	sys.HTTPResponseWithJSON(w, 201, newFair)
}

// Update function to route PUT /v1/fairies
func (h *FairHandler) Update(w http.ResponseWriter, r *http.Request) {
	fairCode := r.URL.Query().Get("fair_code")
	if len(fairCode) == 0 {
		sys.HTTPResponseWithJSON(w, 400, "parameter 'fair_code' is required")
		return
	}

	currentFair, err := h.Service.Fair(fairCode)
	if err != nil {
		sys.HTTPResponseWithJSON(w, 500, err)
		return
	}

	if currentFair == nil {
		sys.HTTPResponseWithJSON(w, 400, "fair not found")
		return
	}

	var fair Fair
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sys.HTTPResponseWithJSON(w, 400, err)
		return
	}

	err = json.Unmarshal(body, &fair)
	if err != nil {
		sys.HTTPResponseWithJSON(w, 400, err)
		return
	}

	fair.ID = currentFair.ID
	err = h.Service.UpdateFair(&fair)
	if err != nil {
		sys.HTTPResponseWithJSON(w, 500, err)
		return
	}

	sys.HTTPResponseWithJSON(w, 200, fair)
}

// Delete function to route DELETE /v1/fairies
func (h *FairHandler) Delete(w http.ResponseWriter, r *http.Request) {
	fairCode := r.URL.Query().Get("fair_code")
	if len(fairCode) == 0 {
		sys.HTTPResponseWithJSON(w, 400, "parameter 'fair_code' is required")
		return
	}

	currentFair, err := h.Service.Fair(fairCode)
	if err != nil {
		sys.HTTPResponseWithJSON(w, 500, err)
		return
	}

	if currentFair == nil {
		sys.HTTPResponseWithJSON(w, 400, "fair not found")
		return
	}

	err = h.Service.DeleteFair(fairCode)
	if err != nil {
		sys.HTTPResponseWithJSON(w, 500, err)
		return
	}

	sys.HTTPResponseWithJSON(w, 200, "fair deleted")
}

// ImportData function to handle POST /v1/import_data
func (h *FairHandler) ImportData(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	if len(file) == 0 {
		sys.HTTPResponseWithJSON(w, 400, "parameter 'file' is required")
		return
	}

	data := sys.LoadCsvData(file)
	if data == nil {
		return
	}

	for i, line := range data {
		if i == 0 {
			continue
		}
		h.Service.ImportFair(line)
	}
}
