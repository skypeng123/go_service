package models

import (
	"github.com/cihub/seelog"
	"go_service/helpers"
)

type BaseCharge struct {
	ParkId      int     `json:"park_id" form:"park_id"`
	ParkAreaId  int     `json:"park_area_id" form:"park_area_id"`
	CarSizeId   int     `json:"car_size_id" form:"car_size_id"`
	Maxforday   float64 `json:"maxforday" form:"maxforday"`
	Days        int     `json:"days" form:"days"`
	CreatedTime int     `json:"created_time" form:"created_time"`
	UpdatedTime int     `json:"updated_time" form:"updated_time"`
}

type Charge1 struct {
	BaseCharge
	Id         int     `json:"time_charge_id" form:"time_charge_id"`
	Freemin    int     `json:"freemin" form:"freemin"`
	Unitprice  float64 `json:"unitprice" form:"unitprice"`
	Unit       int     `json:"unit" form:"unit"`
	Firstunit  int     `json:"firstunit" form:"firstunit"`
	Firstprice float64 `json:"firstprice" form:"firstprice"`
}

func (o *Charge1) GetCharge1(park_id int, park_area_id int, car_size_id int) (charge Charge1, err error) {
	err = helpers.DB.QueryRow("SELECT time_charge_id,park_id, park_area_id, car_size_id, freemin, unitprice, unit, firstunit, firstprice, maxforday, days FROM alk_charge_1 WHERE park_id=? AND park_area_id=? AND car_size_id=? ",
		park_id, park_area_id, car_size_id).Scan(&charge.Id, &charge.ParkId, &charge.ParkAreaId, &charge.CarSizeId, &charge.Freemin, &charge.Unitprice, &charge.Unit, &charge.Firstunit, &charge.Firstprice, &charge.Maxforday, &charge.Days)
	if err != nil {
		seelog.Critical(err)
	}
	return
}
