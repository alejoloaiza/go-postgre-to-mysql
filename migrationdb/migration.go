package migrationdb

import (
	"go-postgre-to-mysql/extra"
	"math/rand"
	"time"
)

func SelectAndInsert() {
	myQuery := `
	SELECT "Business","Code","Type","Location","City","Area","Price","Numrooms","Numbaths","Parking" FROM parallel.webscrapingresults
	`
	rows, err := extra.Postgredb.Query(myQuery)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var pBusiness string
		var pCode string
		var pType string
		var pLocation string
		var pCity string
		var pArea string
		var pPrice string
		var pNumrooms string
		var pNumbaths string
		var pParking string
		err = rows.Scan(&pBusiness, &pCode, &pType, &pLocation, &pCity, &pArea, &pPrice, &pNumrooms, &pNumbaths, &pParking)

		//fmt.Printf("%s | %s | %s | %s | %s | %s | %s | %s | %s | %s | \n", pBusiness, pCode, pType, pLocation, pCity, pArea, pPrice, pNumrooms, pNumbaths, pParking)

		Id := pCode
		Type := pType
		Business := pBusiness
		Sector := pLocation
		City := pCity
		Value := pPrice
		Address := ""
		Area := pArea
		Bedrooms := pNumrooms
		Bathrooms := pNumbaths
		Parking := pParking
		StorageRoom := GetRandom(0, 2)
		OpenKitchen := GetRandom(0, 2)
		Pool := GetRandom(0, 2)
		Gym := GetRandom(0, 2)
		SocialRoom := GetRandom(0, 2)
		CommonAreas := GetRandom(0, 2)
		PetFriendly := GetRandom(0, 2)
		EcoRoutes := GetRandom(0, 2)
		SportFields := GetRandom(0, 2)
		WetArea := GetRandom(0, 2)
		FrontDeskSecurity := GetRandom(0, 2)
		ChildrenPark := GetRandom(0, 2)
		ConstructionAge := GetRandom(0, 10)
		NearMetro := GetRandom(0, 2)
		NearMalls := GetRandom(0, 2)
		NearParks := GetRandom(0, 2)
		NearMainStreets := GetRandom(0, 2)
		NearHospitals := GetRandom(0, 2)
		NearRestaurants := GetRandom(0, 2)
		_, err := extra.MySqldb.Exec("INSERT INTO data.properties VALUES(?,?,?,?,?,?,?,?,?,? ,?,?,?,?,?,?,?,?,?,? ,?,?,?,?,?,?,?,?,?,?);", Id, Type, Business, Sector, Value, City, Address, Area, Bedrooms, Bathrooms, Parking, StorageRoom, OpenKitchen, Pool, Gym, SocialRoom, CommonAreas, PetFriendly, EcoRoutes, SportFields, WetArea, FrontDeskSecurity, ChildrenPark, ConstructionAge, NearMetro, NearMalls, NearParks, NearMainStreets, NearHospitals, NearRestaurants)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

}
func GetRandom(min, max int) int {
	return rand.Intn(max-min) + min
}

func InsertRandom(max int) {
	rand.Seed(time.Now().UTC().UnixNano())
	aType := []string{"Apartamento", "Casa"}
	aBusiness := []string{"Venta", "Arriendo"}
	aSector := []string{"Provenza", "Patio bonito"}
	aCity := []string{"Medellin", "Bello", "Envigado", "Itagui", "Sabaneta"}
	aAddress := []string{"CALLE 44 # 22 41", "CR 27 # 33 11", "CALLE 88 # 22", "CALLE 1 # 2 -3", "CR 33 # 23 - 11"}
	//fmt.Println(myrand)

	var i int
	for i < max {
		i++
		Id := GetRandom(0, 1000000000)
		Type := aType[GetRandom(0, len(aType))]
		Business := aBusiness[GetRandom(0, len(aBusiness))]
		Sector := aSector[GetRandom(0, len(aSector))]
		City := aCity[GetRandom(0, len(aCity))]
		Value := GetRandom(10000000, 1000000000)
		Address := aAddress[GetRandom(0, len(aAddress))]
		Area := GetRandom(10, 250)
		Bedrooms := GetRandom(1, 6)
		Bathrooms := GetRandom(1, 4)
		Parking := GetRandom(0, 2)
		StorageRoom := GetRandom(0, 2)
		OpenKitchen := GetRandom(0, 2)
		Pool := GetRandom(0, 2)
		Gym := GetRandom(0, 2)
		SocialRoom := GetRandom(0, 2)
		CommonAreas := GetRandom(0, 2)
		PetFriendly := GetRandom(0, 2)
		EcoRoutes := GetRandom(0, 2)
		SportFields := GetRandom(0, 2)
		WetArea := GetRandom(0, 2)
		FrontDeskSecurity := GetRandom(0, 2)
		ChildrenPark := GetRandom(0, 2)
		ConstructionAge := GetRandom(0, 10)
		NearMetro := GetRandom(0, 2)
		NearMalls := GetRandom(0, 2)
		NearParks := GetRandom(0, 2)
		NearMainStreets := GetRandom(0, 2)
		NearHospitals := GetRandom(0, 2)
		NearRestaurants := GetRandom(0, 2)
		/*
			fmt.Printf("%s %s %s %s %s %s %s", strconv.Itoa(Id), Business, Type, Sector, City, Value, Address)
			fmt.Printf("%s %s %s %s %s %s %s", strconv.Itoa(Id), Business, Type, Sector, City, Value, Address)
		*/
		//execute statement
		_, err := extra.MySqldb.Exec("INSERT data.properties VALUES(?,?,?,?,?,?,?,?,?,? ,?,?,?,?,?,?,?,?,?,? ,?,?,?,?,?,?,?,?,?,?);", Id, Type, Business, Sector, Value, City, Address, Area, Bedrooms, Bathrooms, Parking, StorageRoom, OpenKitchen, Pool, Gym, SocialRoom, CommonAreas, PetFriendly, EcoRoutes, SportFields, WetArea, FrontDeskSecurity, ChildrenPark, ConstructionAge, NearMetro, NearMalls, NearParks, NearMainStreets, NearHospitals, NearRestaurants)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
}
