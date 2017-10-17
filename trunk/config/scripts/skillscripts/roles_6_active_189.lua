sys.log(" 主角 主动技能  6 开始")

--主角6技能  全体加血

function All_6_Skill(battleid, casterid,attackNum,level,skillid)
	Battle.Cure(battleid,casterid,truedamage,crit)   --jiaxue
	local  t = Player.GetFriends(battleid,casterid,attackNum)  --获取目标
	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid) -- 清空数据
		local  truedamage  = Player.GetMagicDamage(battleid,casterid,v)       --伤害 公式（）
		local per = level * 0.05
		truedamage = truedamage + truedamage * per
		local crit = Battle.GetCrit(skillid)   --是否暴击
		Battle.Cure(battleid,v,truedamage,crit)   --jiaxue
	
		Battle.TargetOver(battleid)
	end
end

function SK_175_Action(battleid, casterid)
	
	
	local skillid = 175-- 技能id

	local level = 1
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
function SK_176_Action(battleid, casterid)
	
	
	local skillid = 176-- 技能id

	local level = 2
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
function SK_177_Action(battleid, casterid)
	
	
	local skillid = 177-- 技能id

	local level = 3
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
function SK_178_Action(battleid, casterid)
	
	
	local skillid = 178-- 技能id

	local level = 4
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
function SK_179_Action(battleid, casterid)
	
	
	local skillid = 179-- 技能id

	local level = 5
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
function SK_180_Action(battleid, casterid)
	
	
	local skillid = 180-- 技能id

	local level = 6
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
function SK_181_Action(battleid, casterid)
	
	
	local skillid = 181-- 技能id

	local level = 7
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
function SK_182_Action(battleid, casterid)
	
	
	local skillid = 182-- 技能id

	local level = 8
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
function SK_183_Action(battleid, casterid)
	
	
	local skillid = 183-- 技能id

	local level = 9
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
function SK_184_Action(battleid, casterid)
	
	
	local skillid = 184-- 技能id

	local level = 10
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
function SK_185_Action(battleid, casterid)
	
	
	local skillid = 185-- 技能id

	local level = 11
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
function SK_186_Action(battleid, casterid)
	
	
	local skillid = 186-- 技能id

	local level = 12
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
function SK_187_Action(battleid, casterid)

	
	local skillid = 187-- 技能id

	local level = 13
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
function SK_188_Action(battleid, casterid)
	
	
	local skillid = 188-- 技能id

	local level = 14
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
function SK_189_Action(battleid, casterid)
	
	
	local skillid = 189-- 技能id

	local level = 15
	
	local attackNum = 0

	All_6_Skill(battleid, casterid,attackNum,level,skillid)

end
sys.log(" 主角 主动技能  6 结束")