sys.log(" skill 39 start")

--主角2技能  减少对方单体所有属性

function SK_138_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 138	-- 技能id

	local level = 1
	
	local  t = Player.GetTarget(battleid,casterid)  --获取目标

	skill1_user_atk(battleid,casterid,t,level,127,1)

	Battle.TargetOver(battleid)

end

sys.log(" skill 39 end")