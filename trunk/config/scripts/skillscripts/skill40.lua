sys.log(" skill 40 start")

--主角3技能  增加友方单体目标暴击

function SK_139_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 139	-- 技能id

	local level = 1
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end

sys.log(" skill 40 end")