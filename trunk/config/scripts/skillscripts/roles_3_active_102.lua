sys.log("主角 主动技能 3 开始")

--主角3技能  增加友方quan体所有攻击

function SK_102_Action(battleid, casterid)
	
	
	local skillid = 102	-- 技能id

	local level = 1

	local attackNum  =  0
	
	local  t = Player.GetFriends(battleid,casterid,attackNum)  --获取目标

	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid) -- 清空数据
		Battle.Cure(battleid,v,0,0)
		Battle.AddBuff(battleid,casterid,v,105,20)   --增法攻
		Battle.AddBuff(battleid,casterid,v,102,20)   --增物理攻击
	
		Battle.TargetOver(battleid)  
		
		
	end
	
end

sys.log("主角 主动技能 3 结束")
