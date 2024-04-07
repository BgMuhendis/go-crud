package respository

import (
	"database/sql"
	"fmt"
	"go-crud/model"
)

type CityRepo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *CityRepo {
	return &CityRepo{
		db: db,
	}
}

func (repo CityRepo) Insert(city model.City) {
	stmt, err := repo.db.Prepare("insert into cities(name,code) values($1,$2)")

	r, err := stmt.Exec(city.Name, city.Code)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r.RowsAffected())
	}
}

func (repo CityRepo) List() []model.City {

	var cityList []model.City
	rows, err := repo.db.Query("select * from cities")
	if err != nil {
		fmt.Println(err)
		return cityList

	} else {

		for rows.Next() {
			var city model.City
			err := rows.Scan(&city.Name, &city.Id, &city.Code)
			if err != nil {
				fmt.Println(err)
			} else {
				cityList = append(cityList, city)
			}
		}
		rows.Close()
		return cityList
	}
}

func (repo CityRepo) GetById(id int) *model.City {
	var city model.City
	formattedSql := fmt.Sprintf("select * from cities where id= %v", id)
	err := repo.db.QueryRow(formattedSql).Scan(&city.Name, &city.Id, &city.Code)
	if err != nil {
		return nil
	} else {
		return &city
	}
}

func (repo CityRepo) DeleteById(id int) {
	stmt, err := repo.db.Prepare("delete from cities where id= $1")

	if err != nil {
		fmt.Println(err)

	} else {
		stmt.Query(id)
	}
}

func (repo CityRepo) selectWithPreparedStatement(cityName string) {
	stmt, err := repo.db.Prepare("select * from cities where id= $1")

	if err != nil {
		return
	} else {
		var city model.City
		err := stmt.QueryRow(cityName).Scan(&city.Id, &city.Name, city.Code)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(city)
		}
	}
}
