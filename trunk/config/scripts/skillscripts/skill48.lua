sys.log(" skill 48 start")

--主角被动2技能  增加防御力20*技能等级


function SK_147_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 147	-- 技能id

	local level = 1
	
	--local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,casterid,level,136,20)

	Battle.TargetOver(battleid)

end

sys.log(" skill 48 end")