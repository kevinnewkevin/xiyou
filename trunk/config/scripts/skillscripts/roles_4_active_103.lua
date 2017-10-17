sys.log("主角 主动技能  4 开始")

--主角4技能  增加友方shangfang

function SK_103_Action(battleid, casterid)

	
	local skillid = 103	-- 技能id

	local level = 1
	
	local attackNum = 0 

	local  t = Player.GetFriends(battleid,casterid,attackNum)  --获取目标

	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid) -- 清空数据
		Battle.Cure(battleid,v,0,0)
		Battle.AddBuff(battleid,casterid,v,112,20)   --增shangfang
		Battle.AddBuff(battleid,casterid,v,112,20)   --增shangfang
		Battle.TargetOver(battleid)  
		
		
	end

end


sys.log("主角 主动技能  4 结束")
