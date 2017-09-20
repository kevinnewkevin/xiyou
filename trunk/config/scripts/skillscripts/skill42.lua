sys.log(" skill 42 start")

--主角5技能  增加友方单体目标连击

function SK_141_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 141	-- 技能id

	local level = 1
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,130,5)

	Battle.TargetOver(battleid)

end

sys.log(" skill 42 end")