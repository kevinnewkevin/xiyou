sys.log(" skill 38 start")

--主角1技能  增加自己所有属性

function SK_137_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 137	-- 技能id

	local level = 1
	
	--local  t = Player.GetTarget(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,casterid,level,126,1)

	Battle.TargetOver(battleid)

end
sys.log(" skill 38 end")