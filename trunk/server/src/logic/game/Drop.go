package game

import (
	"fmt"
	"math/rand"
	"time"
	"suzuki/conf"
)

//////////////////权重随机////////////////////////////////////////
func RandWeight(pair map[int32]int32) int32 {
	if len(pair) <= 0 {
		return -1
	}
	var(
		rr				= rand.New(rand.NewSource(time.Now().UnixNano()))
		first 	int32  = 0
		sum		int32  = 0
	)

	//for k,v := range pair{
	//	sum += v
	//	fmt.Println(k,v)
	//}

	for _,v := range pair{
		sum += v
	}

	randNum := rr.Int31n(sum)

	for k,v := range pair{
		if randNum < v {
			break
		}
		randNum -= v
		first = k
	}
	//fmt.Println(first)
	return first
}

/////////////////////////////////产出///////////////////////////////////////\
type (
	DropItem struct {
		ItemId		int32
		ItemNum		int32
		ProbType	int32
		Prob		int32
	}
	Drop struct {
		DropID		int32
		Exp			int32
		Money		int32
		Items		[]DropItem
	}
)

var (
	DropTableData  = map[int32]*Drop{}
)

func LoadDropTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		DropId := csv.GetInt32(r,"DropID")

		if DropId <= 0 {
			fmt.Println("DropTable Nonstandard DropId",DropId)
			continue
		}

		if DropTableData[DropId] != nil {
			fmt.Println("DropTable Have The Same ID",DropId)
			continue
		}
		item := Drop{}
		item.DropID			= DropId
		item.Exp			= csv.GetInt32(r,"exp")
		item.Money			= csv.GetInt32(r,"money")

		item1 := DropItem{}
		item1.ItemId		= csv.GetInt32(r,"item-1");
		item1.ItemNum		= csv.GetInt32(r,"item-num-1");
		item1.ProbType		= csv.GetInt32(r,"pro-type-1");
		item1.Prob			= csv.GetInt32(r,"probability-1");

		item2 := DropItem{}
		item2.ItemId		= csv.GetInt32(r,"item-2");
		item2.ItemNum		= csv.GetInt32(r,"item-num-2");
		item2.ProbType		= csv.GetInt32(r,"pro-type-2");
		item2.Prob			= csv.GetInt32(r,"probability-2");

		item3 := DropItem{}
		item3.ItemId		= csv.GetInt32(r,"item-3");
		item3.ItemNum		= csv.GetInt32(r,"item-num-3");
		item3.ProbType		= csv.GetInt32(r,"pro-type-3");
		item3.Prob		= csv.GetInt32(r,"probability-3");

		item4 := DropItem{}
		item4.ItemId		= csv.GetInt32(r,"item-4");
		item4.ItemNum		= csv.GetInt32(r,"item-num-4");
		item4.ProbType		= csv.GetInt32(r,"pro-type-4");
		item4.Prob			= csv.GetInt32(r,"probability-4");

		item5 := DropItem{}
		item5.ItemId		= csv.GetInt32(r,"item-5");
		item5.ItemNum		= csv.GetInt32(r,"item-num-5");
		item5.ProbType		= csv.GetInt32(r,"pro-type-5");
		item5.Prob			= csv.GetInt32(r,"probability-5");

		item6 := DropItem{}
		item6.ItemId		= csv.GetInt32(r,"item-6");
		item6.ItemNum		= csv.GetInt32(r,"item-num-6");
		item6.ProbType		= csv.GetInt32(r,"pro-type-6");
		item6.Prob			= csv.GetInt32(r,"probability-6");

		item.Items = append(item.Items,item1)
		item.Items = append(item.Items,item2)
		item.Items = append(item.Items,item3)
		item.Items = append(item.Items,item4)
		item.Items = append(item.Items,item5)
		item.Items = append(item.Items,item6)

		DropTableData[DropId] = &item
	}
	return nil
}

func GetDropBaseById(dropid int32) *Drop {
	return DropTableData[dropid]
}

func GetDropById(dropid int32) *Drop {
	drop := GetDropBaseById(dropid)
	if drop==nil {
		fmt.Println("Can Not Find Drop In DropTable DropId=",dropid)
		return nil
	}

	var rr = rand.New(rand.NewSource(time.Now().UnixNano()))

	newDrop := Drop{}
	newDrop.DropID 	= drop.DropID
	newDrop.Exp		= drop.Exp
	newDrop.Money	= drop.Money

	if len(drop.Items) == 0 {
		return &newDrop
	}

	for _,item := range drop.Items{
		if item.ItemId == 0 {
			continue
		}
		if item.ProbType == 1 {
			var ratio int32 = rr.Int31n(10000)
			if ratio <= item.Prob {
				newDrop.Items = append(newDrop.Items,item)
			}
		}else if item.ProbType == 2 {
			map1 := make(map[int32]int32)
			for i:=0;i<len(drop.Items) ;i++  {
				map1[int32(i)] = drop.Items[i].Prob
			}
			index := RandWeight(map1)
			newDrop.Items = append(newDrop.Items,drop.Items[index])
			break
		}else{
			fmt.Println("Drop ProbType Error DropId=",dropid,"item=",item.ItemId)
		}
	}

	return &newDrop
}