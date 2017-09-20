sys.log(" skill 44 start")

--主角7技能  增加友方单体目标回复

function SK_143_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 143-- 技能id

	local level = 1
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,132,5)

	Battle.TargetOver(battleid)

end

sys.log(" skill 44 end")