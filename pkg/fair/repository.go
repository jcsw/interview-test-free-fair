package fair

import (
	"fmt"

	mariadb "interview-test-free-fair/pkg/infra/mariadb"
	sys "interview-test-free-fair/pkg/sys"
)

const (
	INSERT                 = "INSERT INTO free_fair(longitude,latitude,setcens,areap,district_code,district_name,sub_pref_cod,sub_pref_name,region_05,region_O8,fair_name,fair_code,addres_street,address_number,address_distric,address_reference) VALUES (\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\")"
	SELECT                 = "SELECT id,longitude,latitude,setcens,areap,district_code,district_name,sub_pref_cod,sub_pref_name,region_05,region_O8,fair_name,fair_code,addres_street,address_number,address_distric,address_reference FROM free_fair"
	WHERE_CLAUSE_INIT      = " WHERE %v"
	WHERE_CLAUSE_AND       = "%v AND %v"
	FAIR_NAME_CLAUSE       = "fair_name = \"%v\""
	DISTRICT_NAME_CLAUSE   = "district_name = \"%v\""
	REGION_05_CLAUSE       = "region_05 = \"%v\""
	ADDRESS_DISTRIC_CLAUSE = "address_distric = \"%v\""
)

type searchParameter struct {
	FairName       string
	DistrictName   string
	Region05       string
	AddressDistric string
}

func search(searchParameter searchParameter) []FreeFair {
	sys.LogInfo("Search searchParameter [%+v]", searchParameter)

	db := mariadb.RetrieveClient()

	var whereClause string
	if len(searchParameter.FairName) > 0 {
		whereClause = addWhereClause(whereClause, fmt.Sprintf(FAIR_NAME_CLAUSE, searchParameter.FairName))
	}

	if len(searchParameter.DistrictName) > 0 {
		whereClause = addWhereClause(whereClause, fmt.Sprintf(DISTRICT_NAME_CLAUSE, searchParameter.DistrictName))
	}

	if len(searchParameter.Region05) > 0 {
		whereClause = addWhereClause(whereClause, fmt.Sprintf(REGION_05_CLAUSE, searchParameter.Region05))
	}

	if len(searchParameter.AddressDistric) > 0 {
		whereClause = addWhereClause(whereClause, fmt.Sprintf(ADDRESS_DISTRIC_CLAUSE, searchParameter.AddressDistric))
	}

	sql := SELECT + whereClause
	res, err := db.Query(sql)
	if err != nil {
		sys.LogError("Error on execute query [%v]", sql, err)
		return nil
	}
	defer res.Close()

	var result []FreeFair
	for res.Next() {
		var fair FreeFair
		err := res.Scan(
			&fair.ID,
			&fair.Long,
			&fair.Lat,
			&fair.Setcens,
			&fair.Areap,
			&fair.DistrictCode,
			&fair.DistrictName,
			&fair.SubPrefCod,
			&fair.SubPrefName,
			&fair.Region05,
			&fair.RegionO8,
			&fair.FairName,
			&fair.FairCode,
			&fair.AddresStreet,
			&fair.AddressNumber,
			&fair.AddressDistric,
			&fair.AddressReference,
		)

		if err != nil {
			sys.LogError("Error on parse query result", err)
			return nil
		}

		result = append(result, fair)
	}

	return result
}

func addWhereClause(currentWhere string, newClause string) string {
	if len(currentWhere) == 0 {
		return fmt.Sprintf(WHERE_CLAUSE_INIT, newClause)
	} else {
		return fmt.Sprintf(WHERE_CLAUSE_AND, currentWhere, newClause)
	}
}

func save(fair *FreeFair) *FreeFair {
	db := mariadb.RetrieveClient()

	sql := fmt.Sprintf(INSERT,
		fair.Long,
		fair.Lat,
		fair.Setcens,
		fair.Areap,
		fair.DistrictCode,
		fair.DistrictName,
		fair.SubPrefCod,
		fair.SubPrefName,
		fair.Region05,
		fair.RegionO8,
		fair.FairName,
		fair.FairCode,
		fair.AddresStreet,
		fair.AddressNumber,
		fair.AddressDistric,
		fair.AddressReference,
	)

	res, err := db.Exec(sql)
	if err != nil {
		sys.LogError("Error on execute query [%v]", sql, err)
		return nil
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		sys.LogError("Error on get id query", err)
		return nil
	}

	sys.LogInfo("Saved ID:%v FairName:%v FairCode:%v", lastId, fair.FairName, fair.FairCode)
	return fair
}

func saveFromCsvData(data []string) {
	db := mariadb.RetrieveClient()

	sql := fmt.Sprintf(INSERT,
		data[1],
		data[2],
		data[3],
		data[4],
		data[5],
		data[6],
		data[7],
		data[8],
		data[9],
		data[10],
		data[11],
		data[12],
		data[13],
		data[14],
		data[15],
		data[16],
	)

	res, err := db.Exec(sql)
	if err != nil {
		sys.LogError("Error on execute query [%v]", sql, err)
		return
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		sys.LogError("Error on get id query", err)
		return
	}

	sys.LogInfo("Saved ID:%v FairName:%v FairCode:%v", lastId, data[11], data[12])
}
