package game

import (
	"fmt"
	"logic/conf"
	"logic/prpc"
	"sync/atomic"
)

type (
	ItemData struct {
		ItemId       int32
		ItemMainType int
		MaxCount     int32
		GloAction    string
		SoulVal      int32 //卡牌碎片分解的魂值
	}
)

var (
	ItemTableData        = map[int32]*ItemData{}
	rootItemInstId int64 = 1
)

func LoadItemTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		itemId := csv.GetInt32(r, "ItemID")

		if itemId <= 0 {
			fmt.Println("ItemTable Nonstandard ItemId", itemId)
			continue
		}

		if ItemTableData[itemId] != nil {
			fmt.Println("ItemTable Have The Same ID")
			continue
		}
		item := ItemData{}
		item.ItemId = itemId

		itemtype := csv.GetString(r, "ItemType")
		item.ItemMainType = prpc.ToId_ItemMainType(itemtype)

		item.MaxCount = csv.GetInt32(r, "MaxCount")
		item.GloAction = csv.GetString(r, "GloAction")
		item.SoulVal = csv.GetInt32(r, "Price")

		ItemTableData[item.ItemId] = &item
	}
	return nil
}

func GetItemTableDataById(itemid int32) *ItemData {
	return ItemTableData[itemid]
}

func GenItemInst(itemId int32, itemCount int32) []*prpc.COM_ItemInst {
	itemData := GetItemTableDataById(itemId)
	if itemData == nil {
		return nil
	}

	if itemData.MaxCount == 0 {
		fmt.Println("Item MaxCount == 0")
		return nil
	}

	items := []*prpc.COM_ItemInst{}

	for itemCount > 0 {
		item := prpc.COM_ItemInst{}
		item.ItemId = itemData.ItemId
		item.InstId = atomic.AddInt64(&rootItemInstId, 1)

		if itemData.MaxCount > itemCount {
			item.Stack = itemCount
			itemCount = 0
		} else {
			item.Stack = itemData.MaxCount
			itemCount -= itemData.MaxCount
		}

		items = append(items, &item)
	}

	return items
}
