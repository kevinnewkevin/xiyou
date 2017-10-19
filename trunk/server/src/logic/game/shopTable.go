package game

import (
	"suzuki/conf"
	"logic/prpc"
	"fmt"
	"strings"
	"strconv"
	"time"
	"math/rand"
)

type (
	ShopData struct {
		ShopId				int32
		ShopType			int
		CardPondId			int32
		CurrenciesKind		int32
		Price				int32
		Copper				int32
		ShopItemId			int32
	}
	CardPondData struct {
		PondId				int32
		OrangePond			[]int32				//产出池
		OrangeNum			int32				//产出数量
		PurplePond			[]int32
		PurpleNum			int32
		BluePond			[]int32
		BlueNum				int32
		GreenPond			[]int32
		GreenNum			int32
	}
)

var (
	ShopTableData  		= map[int32]*ShopData{}
	CardPondTableData  	= map[int32]*CardPondData{}
	BlackMarket			[]int32
	BlackMarketSize	int = 6
)

func LoadShopTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		c := ShopData{}
		id := csv.GetInt32(r,"ID")
		if ShopTableData[id] != nil {
			return fmt.Errorf("ShopTableData SameId=%d",id)
		}
		c.ShopId 				= id
		shoptype 				:= csv.GetString(r,"ShopType")
		c.ShopType				= prpc.ToId_ShopType(shoptype)
		c.CardPondId			= csv.GetInt32(r,"CardcloseID")
		kind 					:= csv.GetString(r,"ShopPayType")
		c.CurrenciesKind		= int32(prpc.ToId_IPropertyType(kind))
		c.Price					= csv.GetInt32(r,"Price")
		c.Copper				= csv.GetInt32(r,"COPPER")
		c.ShopItemId			= csv.GetInt32(r,"ItemID")

		if c.ShopType == prpc.SHT_BlackMarket {
			if GetItemTableDataById(c.ShopItemId) == nil {
				return fmt.Errorf("LoadShopTable Can not find itemId=%d",c.ShopItemId)
			}
			BlackMarket = append(BlackMarket,c.ShopId)
		}

		ShopTableData[c.ShopId] = &c
	}
	return nil
}

func GetShopDataById(id int32) *ShopData {
	return ShopTableData[id]
}

func GetBlackMarketShopItems() []int32 {
	var temp []int32
	var index int  = 0
	var rr = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<BlackMarketSize; i++{
		index = rr.Intn(len(BlackMarket))
		temp = append(temp,BlackMarket[index])
	}

	return temp
}

///////////////////////////////////////////////////////////////////////////////////////////

func LoadCardPondTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		c := CardPondData{}
		id := csv.GetInt32(r,"ID")
		if CardPondTableData[id] != nil {
			return fmt.Errorf("CardPondTableData SameId=%d",id)
		}
		c.PondId 				= id
		strTmp := strings.Split(csv.GetString(r,"OrangeitemID"),";")
		for i:=0;i<len(strTmp);i++{
			itemid,_ := strconv.Atoi(strTmp[i])

			if GetItemTableDataById(int32(itemid)) == nil {
				return fmt.Errorf("LoadCardPondTable Can not find itemId=%d",itemid)
			}

			c.OrangePond = append(c.OrangePond,int32(itemid))
		}
		c.OrangeNum = csv.GetInt32(r,"Orangenum")

		strTmp1 := strings.Split(csv.GetString(r,"PurpleitemID"),";")
		for i:=0;i<len(strTmp1);i++{
			itemid,_ := strconv.Atoi(strTmp1[i])

			if GetItemTableDataById(int32(itemid)) == nil {
				return fmt.Errorf("LoadCardPondTable Can not find itemId=%d",itemid)
			}

			c.PurplePond = append(c.PurplePond,int32(itemid))
		}
		c.PurpleNum = csv.GetInt32(r,"Purplenum")

		strTmp2 := strings.Split(csv.GetString(r,"BlueitemID"),";")
		for i:=0;i<len(strTmp2);i++{
			itemid,_ := strconv.Atoi(strTmp2[i])

			if GetItemTableDataById(int32(itemid)) == nil {
				return fmt.Errorf("LoadCardPondTable Can not find itemId=%d",itemid)
			}

			c.BluePond = append(c.BluePond,int32(itemid))
		}
		c.BlueNum = csv.GetInt32(r,"Bluenum")

		strTmp3 := strings.Split(csv.GetString(r,"GreenitemID"),";")
		for i:=0;i<len(strTmp3);i++{
			itemid,_ := strconv.Atoi(strTmp3[i])

			if GetItemTableDataById(int32(itemid)) == nil {
				return fmt.Errorf("LoadCardPondTable Can not find itemId=%d",itemid)
			}

			c.GreenPond = append(c.GreenPond,int32(itemid))
		}
		c.GreenNum = csv.GetInt32(r,"Greennum")

		CardPondTableData[c.PondId] = &c
	}
	return nil
}

func GetCardPondTableDataById(pondId int32) *CardPondData {
	return CardPondTableData[pondId]
}

func (this *CardPondData)GetOrangeCardItems() []int32 {
	var temp []int32
	var index int  = 0
	var rr = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<int(this.OrangeNum); i++{
		index = rr.Intn(len(this.OrangePond))
		temp = append(temp,this.OrangePond[index])
	}

	return temp
}

func (this *CardPondData)GetPurpleCardItems() []int32 {
	var temp []int32
	var index int  = 0
	var rr = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<int(this.PurpleNum); i++{
		index = rr.Intn(len(this.PurplePond))
		temp = append(temp,this.PurplePond[index])
	}

	return temp
}

func (this *CardPondData)GetBlueCardItems() []int32 {
	var temp []int32
	var index int  = 0
	var rr = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<int(this.BlueNum); i++{
		index = rr.Intn(len(this.BluePond))
		temp = append(temp,this.BluePond[index])
	}

	return temp
}

func (this *CardPondData)GetGreenCardItems() []int32 {
	var temp []int32
	var index int  = 0
	var rr = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<int(this.GreenNum); i++{
		index = rr.Intn(len(this.GreenPond))
		temp = append(temp,this.GreenPond[index])
	}

	return temp
}