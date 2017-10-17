sys.log(" 主角 主动技能  5 开始")
--主角5技能  自身加护盾
function All_5_Skill(battleid,casterid,level)

	sys.log("主角 5   战斗id   "..  battleid.. "释放者id  " ..casterid)
	
	local t = Player.GetFriend(battleid,casterid)  --获取目标

	sys.log("目标 "..  t)
	
	local hp = Player.GetUnitProperty(battleid, casterid, "CPT_HP")	-- 获取到攻击者的属性
	
	local per = 0.05 * level

	local damage =  hp * 0.1

	damage = damage + damage * per
	
	Battle.Cure(battleid, casterid, 0, 0)
	
	local buffid = 103
	
	Battle.AddBuff(battleid,casterid,casterid,buffid,damage)  --给zishen提供一个盾牌

end

function SK_160_Action(battleid, casterid)
	Battle.TargetOn(battleid)

	sys.log("这是160技能")
	
	local skillid = 160	-- 技能id

	local level = 1
	
	All_5_Skill(battleid, casterid,level)

	Battle.TargetOver(battleid)


end
function SK_161_Action(battleid, casterid)
	
	Battle.TargetOn(battleid)
	local skillid = 161	-- 技能id

	local level = 2
	
	All_5_Skill(battleid, casterid,level)
	Battle.TargetOver(battleid)

end
function SK_162_Action(battleid, casterid)
	
	Battle.TargetOn(battleid)
	local skillid = 162	-- 技能id

	local level = 3
	
	All_5_Skill(battleid, casterid,level)
	Battle.TargetOver(battleid)

end
function SK_163_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 163	-- 技能id

	local level = 4
	
	All_5_Skill(battleid, casterid,level)
	Battle.TargetOver(battleid)

end
function SK_164_Action(battleid, casterid)
	
	Battle.TargetOn(battleid)
	local skillid = 164	-- 技能id

	local level = 5
	
	All_5_Skill(battleid, casterid,level)
	Battle.TargetOver(battleid)

end
function SK_165_Action(battleid, casterid)
	
	Battle.TargetOn(battleid)
	local skillid = 165	-- 技能id

	local level = 6
	
	All_5_Skill(battleid, casterid,level)
	Battle.TargetOver(battleid)

end
function SK_166_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 166	-- 技能id

	local level = 7
	
	All_5_Skill(battleid, casterid,level)
	Battle.TargetOver(battleid)

end
function SK_167_Action(battleid, casterid)
	
	Battle.TargetOn(battleid)
	local skillid = 167	-- 技能id

	local level = 8
	
	All_5_Skill(battleid, casterid,level)
	Battle.TargetOver(battleid)

end
function SK_168_Action(battleid, casterid)
	
	Battle.TargetOn(battleid)
	local skillid = 168	-- 技能id

	local level = 9
	
	All_5_Skill(battleid, casterid,level)
	Battle.TargetOver(battleid)

end
function SK_169_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 169	-- 技能id

	local level = 10
	
	All_5_Skill(battleid, casterid,level)
	Battle.TargetOver(battleid)

end
function SK_170_Action(battleid, casterid)
	
	Battle.TargetOn(battleid)
	local skillid = 170	-- 技能id

	local level = 11
	
	All_5_Skill(battleid, casterid,level)
	Battle.TargetOver(battleid)

end
function SK_171_Action(battleid, casterid)
	
	Battle.TargetOn(battleid)
	local skillid = 171	-- 技能id

	local level = 12
	
	All_5_Skill(battleid, casterid,level)
	Battle.TargetOver(battleid)

end
function SK_172_Action(battleid, casterid)
	
	Battle.TargetOn(battleid)
	local skillid = 172	-- 技能id

	local level = 13
	
	All_5_Skill(battleid, casterid,level)
	Battle.TargetOver(battleid)

end
function SK_173_Action(battleid, casterid)
	
	Battle.TargetOn(battleid)
	local skillid = 173	-- 技能id

	local level = 14
	
	All_5_Skill(battleid, casterid,level)
	Battle.TargetOver(battleid)

end
function SK_174_Action(battleid, casterid)
	
	Battle.TargetOn(battleid)
	local skillid = 174	-- 技能id

	local level = 15
	
	All_5_Skill(battleid, casterid,level)
	Battle.TargetOver(battleid)

end
sys.log(" 主角 主动技能  5 结束")