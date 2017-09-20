sys.log(" skill 46 start")

--主角9技能  增加友方单体目标吸血

function SK_145_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 145	-- 技能id

	local level = 1
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,134,5)

	Battle.TargetOver(battleid)

end

sys.log(" skill 46 end")