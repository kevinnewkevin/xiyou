sys.log(" skill 49 start")

--主角被动3技能  增加防御力20*技能等级


function SK_148_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 148	-- 技能id

	local level = 1
	
	--local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,casterid,level,137,20)

	Battle.TargetOver(battleid)

end

sys.log(" skill 49 end")