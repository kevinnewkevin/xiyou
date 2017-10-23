sys.log("主角 主动技能 3 开始")

--主角3技能  增加法  物  攻击  每升一级加伤0.05

function All_3_Skill(battleid,casterid,attackNum,level)
	local  t = Player.GetFriends(battleid,casterid,attackNum)  --获取目标

	for i,v in ipairs(t) do
		Battle.TargetOn(battleid) -- 清空数据
		local  atkdamage  = Player.GetUnitDamage(battleid,casterid,v)       --伤害 公式（）
		local  magdamage  = Player.GetMagicDamage(battleid,casterid,v)       --伤害 公式（）
		local per = 0.05 * level
		atkdamage = atkdamage + atkdamage * per
		magdamage = magdamage + magdamage * per
		Battle.Cure(battleid,v,0,0)
		Battle.AddBuff(battleid,casterid,v,105,magdamage)   --增法攻
		Battle.AddBuff(battleid,casterid,v,102,atkdamage)   --增物理攻击
		Battle.TargetOver(battleid)  
		
		
	end
end

function SK_130_Action(battleid, casterid)
	
	
	local skillid = 130	-- 技能id

	local level = 1
	
	local attackNum = 0 

	All_3_Skill(battleid,casterid,attackNum,level)

end
function SK_131_Action(battleid, casterid)
	
	
	local skillid = 131	-- 技能id

	local level = 2
	
	local attackNum = 0 
	
	All_3_Skill(battleid,casterid,attackNum,level)

end
function SK_132_Action(battleid, casterid)
	
	
	local skillid = 132	-- 技能id

	local level = 3
	
	local attackNum = 0 
	
	All_3_Skill(battleid,casterid,attackNum,level)

end
function SK_133_Action(battleid, casterid)
	
	
	local skillid = 133	-- 技能id

	local level = 4
	
	local attackNum = 0 
	
	All_3_Skill(battleid,casterid,attackNum,level)

end
function SK_134_Action(battleid, casterid)
	
	local skillid = 134	-- 技能id

	local level = 5
	
	local attackNum = 0 
	
	All_3_Skill(battleid,casterid,attackNum,level)

end
function SK_135_Action(battleid, casterid)
	
	
	local skillid = 135	-- 技能id

	local level = 6
	
	local attackNum = 0 
	
	All_3_Skill(battleid,casterid,attackNum,level)

end
function SK_136_Action(battleid, casterid)
	
	
	local skillid = 136	-- 技能id

	local level = 7
	
	local attackNum = 0 
	
	All_3_Skill(battleid,casterid,attackNum,level)

end
function SK_137_Action(battleid, casterid)
	
	
	local skillid = 137	-- 技能id

	local level = 8
	
	local attackNum = 0 
	
	All_3_Skill(battleid,casterid,attackNum,level)

end
function SK_138_Action(battleid, casterid)
	
	
	local skillid = 138	-- 技能id

	local level = 9
	
	local attackNum = 0 
	
	All_3_Skill(battleid,casterid,attackNum,level)

end
function SK_139_Action(battleid, casterid)
	
	
	local skillid = 139	-- 技能id

	local level = 10
	
	local attackNum = 0 
	
	All_3_Skill(battleid,casterid,attackNum,level)

end
function SK_140_Action(battleid, casterid)
	
	
	local skillid = 140	-- 技能id

	local level = 11
	
	local attackNum = 0 
	
	All_3_Skill(battleid,casterid,attackNum,level)

end
function SK_141_Action(battleid, casterid)
	
	
	local skillid = 141	-- 技能id

	local level = 12
	
	local attackNum = 0 
	
	All_3_Skill(battleid,casterid,attackNum,level)

end
function SK_142_Action(battleid, casterid)
	
	
	local skillid = 142	-- 技能id

	local level = 13
	
	local attackNum = 0 
	
	All_3_Skill(battleid,casterid,attackNum,level)

end
function SK_143_Action(battleid, casterid)
	
	
	local skillid = 143	-- 技能id

	local level = 14
	
	local attackNum = 0 
	
	All_3_Skill(battleid,casterid,attackNum,level)

end
function SK_144_Action(battleid, casterid)
	
	
	local skillid = 144	-- 技能id

	local level = 15
	
	local attackNum = 0 
	
	All_3_Skill(battleid,casterid,attackNum,level)

end
sys.log("主角 主动技能 3 结束")
