sys.log(" skill 47 start")

--主角被动1技能  增加生命值50*技能等级

function SK_146_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 146	-- 技能id

	local level = 1
	
	--local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,casterid,level,101,50)

	Battle.TargetOver(battleid)

end

sys.log(" skill 47 end")