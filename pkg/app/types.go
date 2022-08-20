package app

type Fair struct {
	ID               int64  `json:"id"`
	Long             string `json:"long"`
	Lat              string `json:"lat"`
	Setcens          string `json:"setcens"`
	Areap            string `json:"are_ap"`
	DistrictCode     string `json:"district_code"`
	DistrictName     string `json:"district_name"`
	SubPrefCod       string `json:"sub_pref_code"`
	SubPrefName      string `json:"sub_pref_name"`
	Region05         string `json:"region_05"`
	RegionO8         string `json:"region_08"`
	FairName         string `json:"fair_name"`
	FairCode         string `json:"fair_code"`
	AddresStreet     string `json:"addres_street"`
	AddressNumber    string `json:"address_number"`
	AddressDistric   string `json:"address_district"`
	AddressReference string `json:"address_reference"`
}

type FairSearchParameter struct {
	FairName       string
	DistrictName   string
	Region05       string
	AddressDistric string
}

type FairService interface {
	Fair(code string) (*Fair, error)
	Fairs(parameters FairSearchParameter) ([]Fair, error)
	CreateFair(fair *Fair) error
	ImportFair(data []string) error
	UpdateFair(fair *Fair) error
	DeleteFair(code string) error
}
