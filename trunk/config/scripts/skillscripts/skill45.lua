sys.log(" skill 45 start")

--主角8技能  增加友方单体目标荆棘

function SK_144_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 144	-- 技能id

	local level = 1
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,133,5)

	Battle.TargetOver(battleid)

end

sys.log(" skill 45 end")