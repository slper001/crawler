package persist

import (
	"database/sql"
	"github.com/crawler/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ItemSaverSql() chan interface{}{
	out := make(chan interface{})
	db ,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/danke?charset=utf8")
	if err!=nil{
		log.Printf("err:%v",err)
	}
	stmt,err := db.Prepare("insert into rent(name,price,area,number,structure,pay,orientaion,floor,region,metro,url)values(?,?,?,?,?,?,?,?,?,?,?)")
	if err!=nil {
		log.Printf("数据库预处理失败：%v",err)
	}
	//defer db.Close()
	go func() {
		itemCount := 0
		zoneCount :=0
		houseCount :=0
		for {
			item := <- out
			switch t:=item.(type) {
			case int:
				zoneCount+=1
				switch t {
				case 1:
					log.Printf("获取南山区子区域url：%d个",zoneCount)
				case 2:
					log.Printf("获取福田区子区域url：%d个",zoneCount)
				case 3:
					log.Printf("获取罗湖区子区域url：%d个",zoneCount)
				case 4:
					log.Printf("获取宝安区子区域url：%d个",zoneCount)
				case 5:
					log.Printf("获取龙岗区子区域url：%d个",zoneCount)
				case 6:
					log.Printf("获取龙华区子区域url：%d个",zoneCount)
				}
			case string:
				houseCount +=1
				log.Printf("获取子区域%v房源url",t)
				log.Printf("获取房屋url：%d个",houseCount)
			case model.Attribute:
				itemCount++
				log.Printf("解析到第%d个房源信息: #%v",itemCount,item)
				_, err :=stmt.Exec(item.(model.Attribute).Name,item.(model.Attribute).Price,item.(model.Attribute).Area,item.(model.Attribute).Number,
					item.(model.Attribute).Structure,item.(model.Attribute).Pay,item.(model.Attribute).Orientation,
					item.(model.Attribute).Floor,item.(model.Attribute).Region,item.(model.Attribute).Metro,
					item.(model.Attribute).Url)
				if err!=nil{
					log.Printf("插入数据库失败：%v",err)
				}
			}
		}
	}()
	return out
}
