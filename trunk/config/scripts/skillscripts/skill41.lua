sys.log(" skill 41 start")

--主角4技能  增加友方单体目标暴击伤害

function SK_140_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 140	-- 技能id

	local level = 1
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,129,5)

	Battle.TargetOver(battleid)

end

sys.log(" skill 42 end")