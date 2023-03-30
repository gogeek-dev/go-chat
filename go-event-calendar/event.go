package event

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type TblCalendarDetail struct {
	Id          int    `gorm:"primaryKey;autoIncrement"`
	Name        string `form:"titlename"`
	Description string `form:"desctext"`
	Date        string
}

var DB *gorm.DB

// viewcalendar
func ViewCalendar(c *gin.Context) {

	c.HTML(200, "calendar.html", nil)
}

// Insert EventDesc
func AddTask(c *gin.Context) {

	var tbladdevent TblCalendarDetail

	c.ShouldBindWith(&tbladdevent, binding.Form)

	createerr := DB.Create(&TblCalendarDetail{
		Name:        tbladdevent.Name,
		Description: tbladdevent.Description,
		Date:        c.Request.PostFormValue("datestr"),
	})

	if createerr != nil {
		log.Println(createerr)
	}

	c.Redirect(301, "/viewcalendar")
}

// ListEvent
func ListEvent(c *gin.Context) {

	var Listcaldet1 []TblCalendarDetail

	DB.Model(TblCalendarDetail{}).Select("date").Group("date").Find(&Listcaldet1)

	var Array []map[string]interface{}

	for _, val := range Listcaldet1 {

		var Listcaldet []TblCalendarDetail

		DB.Model(TblCalendarDetail{}).Where("date=?", val.Date).Limit(5).Find(&Listcaldet)

		for _, value := range Listcaldet {
			Array = append(Array, map[string]interface{}{
				"start": value.Date,
				"title": value.Name,
			})
		}
	}

	json.NewEncoder(c.Writer).Encode(Array)
}

// individual date list
func IndivDateEventandcount(c *gin.Context) {

	var Listcaldet []TblCalendarDetail

	if err := DB.Model(TblCalendarDetail{}).Where("date=?", c.Query("date")).Find(&Listcaldet).Error; err != nil {
		log.Println(err)
	}

	var count int64

	DB.Model(Listcaldet).Where("date=?", c.Query("date")).Count(&count)

	c.HTML(200, "calendarlist.html", gin.H{"list": Listcaldet, "count": count, "date": c.Query("date")})

}

func AddRouteandMigrate(r *gin.Engine, Db *gorm.DB) error {

	err := Db.AutoMigrate(
		&TblCalendarDetail{},
	)

	DB = Db

	if err != nil {
		return err
	}

	r.POST("/insert", AddTask)

	r.GET("/getdata", ListEvent)

	r.GET("/datelist", IndivDateEventandcount)

	r.GET("/viewcalendar", ViewCalendar)

	return nil

}
