package app

import (
	"database/sql"
	"fmt"

	sys "interview-test-free-fair/pkg/sys"
)

const (
	INSERT = "INSERT INTO free_fair(longitude,latitude,setcens,areap,district_code,district_name,sub_pref_cod,sub_pref_name,region_05,region_O8,fair_name,fair_code,addres_street,address_number,address_distric,address_reference) VALUES (\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\")"
	SELECT = "SELECT id,longitude,latitude,setcens,areap,district_code,district_name,sub_pref_cod,sub_pref_name,region_05,region_O8,fair_name,fair_code,addres_street,address_number,address_distric,address_reference FROM free_fair"
	DELETE = "DELETE free_fair WHERE fair_code=%v"
	UPDATE = "UPDATE free_fair SET longitude=\"%v\",latitude=\"%v\",setcens=\"%v\",areap=\"%v\",district_code=\"%v\",district_name=\"%v\",sub_pref_cod=\"%v\",sub_pref_name=\"%v\",region_05=\"%v\",region_O8=\"%v\",fair_name=\"%v\",addres_street=\"%v\",address_number=\"%v\",address_distric=\"%v\",address_reference=\"%v\" WHERE fair_code=\"%v\""

	WHERE_CLAUSE_INIT    = " WHERE %v"
	WHERE_CLAUSE_AND     = "%v AND %v"
	WHERE_CLAUSE_BY_CODE = " WHERE fair_code = \"%v\""

	FAIR_NAME_CLAUSE       = "fair_name = \"%v\""
	DISTRICT_NAME_CLAUSE   = "district_name = \"%v\""
	REGION_05_CLAUSE       = "region_05 = \"%v\""
	ADDRESS_DISTRIC_CLAUSE = "address_distric = \"%v\""
)

// FairService represents a PostgreSQL implementation of myapp.FairService
type FairServiceMariaDb struct {
	BD *sql.DB
}

func (f *FairServiceMariaDb) Fair(code string) (*Fair, error) {
	sql := SELECT + fmt.Sprintf(WHERE_CLAUSE_BY_CODE, code)

	res, err := f.BD.Query(sql)
	if err != nil {
		sys.LogError("Error on execute query [%v]", sql, err)
		return nil, err
	}
	defer res.Close()

	if !res.Next() {
		return nil, nil
	}

	var fair Fair
	err = res.Scan(
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
		return nil, err
	}

	return &fair, nil
}

func (f *FairServiceMariaDb) Fairs(searchParameter FairSearchParameter) ([]Fair, error) {
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
	res, err := f.BD.Query(sql)
	if err != nil {
		sys.LogError("Error on execute query [%v]", sql, err)
		return nil, err
	}
	defer res.Close()

	var result []Fair
	for res.Next() {
		var fair Fair
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
			return nil, err
		}

		result = append(result, fair)
	}

	return result, nil
}

func addWhereClause(currentWhere string, newClause string) string {
	if len(currentWhere) == 0 {
		return fmt.Sprintf(WHERE_CLAUSE_INIT, newClause)
	} else {
		return fmt.Sprintf(WHERE_CLAUSE_AND, currentWhere, newClause)
	}
}

func (f *FairServiceMariaDb) CreateFair(fair *Fair) error {
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

	res, err := f.BD.Exec(sql)
	if err != nil {
		sys.LogError("Error on execute query [%v]", sql, err)
		return err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		sys.LogError("Error on get id query", err)
		return err
	}

	sys.LogInfo("Saved ID:%v FairName:%v FairCode:%v", lastId, fair.FairName, fair.FairCode)
	fair.ID = lastId

	return nil
}

func (f *FairServiceMariaDb) ImportFair(data []string) error {
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

	res, err := f.BD.Exec(sql)
	if err != nil {
		sys.LogError("Error on execute query [%v]", sql, err)
		return err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		sys.LogError("Error on get id", err)
		return err
	}

	sys.LogInfo("Saved ID:%v FairName:%v FairCode:%v", lastId, data[11], data[12])

	return nil
}

func (f *FairServiceMariaDb) UpdateFair(fair *Fair) error {
	sql := fmt.Sprintf(UPDATE,
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
		fair.AddresStreet,
		fair.AddressNumber,
		fair.AddressDistric,
		fair.AddressReference,
		fair.FairCode,
	)

	_, err := f.BD.Exec(sql)
	if err != nil {
		sys.LogError("Error on execute query [%v]", sql, err)
		return err
	}

	return nil
}

func (f *FairServiceMariaDb) DeleteFair(code string) error {
	sql := fmt.Sprintf(DELETE, code)

	_, err := f.BD.Query(sql)
	if err != nil {
		sys.LogError("Error on execute query [%v]", sql, err)
		return err
	}

	return nil
}
