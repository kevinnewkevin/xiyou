sys.log(" 主角 主动技能  6 开始")

--主角6技能  全体加血

function SK_105_Action(battleid, casterid)
	
	
	local skillid = 105-- 技能id

	local level = 1
	local attackNum  =  0
	local  t = Player.GetFriends(battleid,casterid,attackNum)  --获取目标
	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid) -- 清空数据
		local crit = Battle.GetCrit(skillid)   --是否暴击
		Battle.Cure(battleid,v,20,crit)   --jiaxue
	
		Battle.TargetOver(battleid)
	end
end

sys.log(" 主角 主动技能  6 结束")