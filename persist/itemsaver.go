package persist

import (
	"github.com/crawler/model"
	"log"
)

func ItemSaver() chan interface{}{
	out := make(chan interface{})
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
			}
		}
	}()
	return out
}
