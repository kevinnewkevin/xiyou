sys.log(" skill 43 start")

--主角6技能  增加友方单体目标溅射

function SK_142_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 142-- 技能id

	local level = 1
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,131,5)

	Battle.TargetOver(battleid)

end

sys.log(" skill 43 end")