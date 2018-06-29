package services

import (
	"fmt"
	"github.com/cihub/seelog"
	"go_service/helpers"
	"go_service/models"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type chargeRes struct {
	ParkingFee float64 `json:"parking_fee"`
	Days       int     `json:"days"`
	Maxforday  float64 `json:"maxforday"`
}

func Charge(c *gin.Context) {
	//get request params
	park_id := helpers.Intval(c.PostForm("park_id"))
	charge_type := helpers.Intval(c.PostForm("charge_type"))
	car_size_id := helpers.Intval(c.PostForm("car_size_id"))
	park_area_id := helpers.Intval(c.PostForm("park_area_id"))
	start_time := c.PostForm("start_time")
	end_time := c.PostForm("end_time")

	seelog.Debugf("请求参数 park_id:%d,charge_type:%d,car_size_id:%d,park_area_id:%d,start_time:%s,end_time:%s", park_id, charge_type, car_size_id, park_area_id, start_time, end_time)

	if park_id == 0 || charge_type == 0 || car_size_id == 0 || start_time == "" || end_time == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 20001, "msg": "参数错误"})
		return
	}

	if charge_type != 1 && charge_type != 2 && charge_type != 3 && charge_type != 4 && charge_type != 5 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 20011, "msg": "收费类型不存在"})
		return
	}

	var parking_fee float64
	if charge_type == 1 {
		model := new(models.Charge1)
		chargeInfo, err := model.GetCharge1(park_id, park_area_id, car_size_id)
		fmt.Println(chargeInfo.Days)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 20012, "msg": "停车场未设置收费规则"})
			return
		}

		time_slice := timeslice(chargeInfo.Days, start_time, end_time)
		fmt.Println(time_slice)

		var fee float64
		for _, min := range time_slice {
			fee += charge1(min, int64(chargeInfo.Freemin), int64(chargeInfo.Firstunit), chargeInfo.Firstprice, int64(chargeInfo.Unit), chargeInfo.Unitprice, chargeInfo.Maxforday)
		}

		parking_fee = helpers.FormatParkingFee(fee)

		c.JSON(http.StatusOK, gin.H{"code": 200, "data": chargeRes{parking_fee, chargeInfo.Days, chargeInfo.Maxforday}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": chargeRes{}})
	return

}

func charge1(parkingmin int64, freemin int64, firstunit int64, firstprice float64, unit int64, unitprice float64, maxforday float64) float64 {
	var parking_fee float64
	if parkingmin <= freemin {
		//免费时段
		parking_fee = 0.00
	} else if parkingmin <= firstunit {
		//首计时计算
		parking_fee = firstprice
	} else {
		//超出首计时计算
		parking_fee = firstprice + float64(parkingmin-firstunit)/float64(unit)*unitprice
	}

	//最高收费检查
	if parking_fee > maxforday {
		parking_fee = maxforday
	}

	return helpers.FormatParkingFee(parking_fee)
}

func timeslice(day_type int, start_time string, end_time string) (time_slice []int64) {
	start_ts := helpers.TimeUnix(start_time)
	end_ts := helpers.TimeUnix(end_time)

	total_ts := end_ts - start_ts
	total_min := helpers.CeilMin(total_ts)

	if day_type == 1 {
		/*自然日*/
		tomorrow_ts := helpers.TomorrowUnix(start_ts)
		if end_ts < tomorrow_ts {
			//未跨天
			time_slice = append(time_slice, total_min)
		} else if end_ts >= tomorrow_ts && end_ts-tomorrow_ts < 86400 {
			//跨1天
			time1 := helpers.CeilMin(tomorrow_ts - start_ts)
			time2 := helpers.CeilMin(end_ts - tomorrow_ts)
			time_slice = append(time_slice, time1, time2)
		} else {
			//跨多天
			//第一天时间
			time1 := helpers.CeilMin(tomorrow_ts - start_ts)
			time_slice = append(time_slice, time1)

			//中间天时间
			middays := (end_ts - tomorrow_ts) / 86400
			for i := 0; i < int(middays); i++ {
				time_slice = append(time_slice, 1440)
			}

			//最后一天时间
			time2 := helpers.CeilMin(total_ts - (tomorrow_ts - start_ts) - 86400*middays)
			time_slice = append(time_slice, time2)
		}
	} else if day_type == 2 {
		/*24小时*/
		if total_ts < 86400 {
			//未跨天
			time_slice = append(time_slice, total_min)
		} else {
			//跨天
			daynum := total_ts / 86400
			for i := 0; i < int(daynum); i++ {
				time_slice = append(time_slice, 1440)
			}
			time2 := helpers.CeilMin(total_ts - daynum*86400)
			time_slice = append(time_slice, time2)
		}
	}
	return
}
