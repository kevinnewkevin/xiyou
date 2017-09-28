sys.log(" skill 3 start")

--主角3技能  增加友方单体目标暴击

function SK_130_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 130	-- 技能id

	local level = 1
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标
	
	sys.log("好友是"..t)

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end
function SK_131_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 131	-- 技能id

	local level = 2
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end
function SK_132_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 132	-- 技能id

	local level = 3
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end
function SK_133_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 133	-- 技能id

	local level = 4
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end
function SK_134_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 134	-- 技能id

	local level = 5
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end
function SK_135_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 135	-- 技能id

	local level = 6
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end
function SK_136_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 136	-- 技能id

	local level = 7
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end
function SK_137_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 137	-- 技能id

	local level = 8
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end
function SK_138_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 138	-- 技能id

	local level = 9
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end
function SK_139_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 139	-- 技能id

	local level = 10
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end
function SK_140_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 140	-- 技能id

	local level = 11
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end
function SK_141_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 141	-- 技能id

	local level = 12
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end
function SK_142_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 142	-- 技能id

	local level = 13
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end
function SK_143_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 143	-- 技能id

	local level = 14
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end
function SK_144_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 144	-- 技能id

	local level = 15
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标

	skill_user_atk(battleid,casterid,t,level,128,5)

	Battle.TargetOver(battleid)

end

sys.log(" skill 3 end")