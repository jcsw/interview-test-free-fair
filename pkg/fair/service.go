package fair

import (
	sys "interview-test-free-fair/pkg/sys"
)

func createNewFair(fair FreeFair) FreeFair {
	return fair
}

func updateFair(fair FreeFair) FreeFair {
	return fair
}

func importData() {

	data := sys.LoadCsvData("data/DEINFO_AB_FEIRASLIVRES_2014.csv")

	for i, line := range data {
		if i == 0 {
			continue
		}

		saveFromCsvData(line)
	}

}
