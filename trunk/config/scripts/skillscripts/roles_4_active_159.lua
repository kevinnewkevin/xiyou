sys.log("主角 主动技能  4 开始")

--主角4技能  增加友方shangfang

function All_4_Skill(battleid, casterid,attackNum,level)

	local  t = Player.GetFriends(battleid,casterid,attackNum)  --获取目标

	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid) -- 清空数据
		local  atkdef  = Player.GetCalcDef(battleid,v)       --伤害 公式（）
		local  magdaef  = Player.GetCalcMagicDef(battleid,v)       --伤害 公式（）
		local per = 0.05 * level
		atkdef = atkdef + atkdef * per
		magdaef = magdaef + magdaef * per
		Battle.Cure(battleid,v,0,0)
		Battle.AddBuff(battleid,casterid,v,118,atkdef)   --增wulifang
		Battle.AddBuff(battleid,casterid,v,114,magdaef)   --增fashufang
		Battle.TargetOver(battleid)  
	end

end

function SK_145_Action(battleid, casterid)
	
	
	local skillid = 145	-- 技能id

	local level = 1

	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)

end

function SK_146_Action(battleid, casterid)
	
	
	local skillid = 146	-- 技能id

	local level = 2
	
	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)

end
function SK_147_Action(battleid, casterid)
	
	
	local skillid = 147	-- 技能id

	local level = 3
	
	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)

end
function SK_148_Action(battleid, casterid)
	
	
	local skillid = 148	-- 技能id

	local level = 4
	
	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)

end
function SK_149_Action(battleid, casterid)
	
	
	local skillid = 149	-- 技能id

	local level = 5
	
	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)

end
function SK_150_Action(battleid, casterid)
	
	
	local skillid = 150	-- 技能id

	local level = 6
	
	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)

end
function SK_151_Action(battleid, casterid)
	
	
	local skillid = 151	-- 技能id

	local level = 7
	
	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)
end
function SK_152_Action(battleid, casterid)
	
	
	local skillid = 152	-- 技能id

	local level = 8
	
	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)

end
function SK_153_Action(battleid, casterid)
	
	
	local skillid = 153	-- 技能id

	local level = 9
	
	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)

end
function SK_154_Action(battleid, casterid)
	
	
	local skillid = 154	-- 技能id

	local level = 10
	
	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)

end
function SK_155_Action(battleid, casterid)
	
	
	local skillid = 155	-- 技能id

	local level = 11
	
	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)

end
function SK_156_Action(battleid, casterid)
	
	
	local skillid = 156	-- 技能id

	local level = 12
	
	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)

end
function SK_157_Action(battleid, casterid)
	
	
	local skillid = 157	-- 技能id

	local level = 13
	
	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)

end
function SK_158_Action(battleid, casterid)
	
	
	local skillid = 158	-- 技能id

	local level = 14
	
	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)

end
function SK_159_Action(battleid, casterid)
	
	
	local skillid = 159	-- 技能id

	local level = 15
	
	local  attackNum = 0
	
	All_4_Skill(battleid, casterid,attackNum,level)

end
sys.log("主角 主动技能  4 结束")
