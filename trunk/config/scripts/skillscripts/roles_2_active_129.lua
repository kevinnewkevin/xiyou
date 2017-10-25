sys.log("主角 主动技能 2 开始")

--主角2技能 对敌方法攻，每升一级加伤5%
--公用函数
function All_2_Skill(battleid,casterid,attackNum,level,skillid)
	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标

	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid) -- 清空数据
	
	    local  truedamage  = Player.GetMagicDamage(battleid,casterid,v)       --伤害 公式（）
		
		sys.log("主角对目标造成的法术伤害    "..truedamage)
		
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		sys.log("主角对目标造成的最终法术伤害    "..damage)
		--判断伤害
		if damage <= 0 then 
		
			damage = 0
		
		end

		sys.log("zhujue 2 333"  ..  damage)

		local per = 0.05 * level
		
		local crit = Battle.GetCrit(skillid)   --是否暴击

		damage = damage + damage *per

		sys.log("zhujue 2 444"  ..  damage)
		
		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
	
		Battle.TargetOver(battleid)  -- 赋给下个目标
		
		
	end
end

function SK_115_Action(battleid, casterid)
	
	
	local skillid = 115	-- 技能id

	local level = 1

	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end

function SK_116_Action(battleid, casterid)
	
	
	local skillid = 116	-- 技能id

	local level = 2
	
	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end
function SK_117_Action(battleid, casterid)
	
	
	local skillid = 117	-- 技能id

	local level = 3
	
	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end
function SK_118_Action(battleid, casterid)
	
	local skillid = 118	-- 技能id

	local level = 4
	
	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end
function SK_119_Action(battleid, casterid)
	
	
	local skillid = 119	-- 技能id

	local level = 5
	
	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end
function SK_120_Action(battleid, casterid)
	
	
	local skillid = 120	-- 技能id

	local level = 6
	
	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end
function SK_121_Action(battleid, casterid)
	
	
	local skillid = 121-- 技能id

	local level = 7
	
	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end
function SK_122_Action(battleid, casterid)
	
	
	local skillid = 122	-- 技能id

	local level = 8
	
	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end
function SK_123_Action(battleid, casterid)
	
	
	local skillid = 123	-- 技能id

	local level = 9
	
	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end
function SK_124_Action(battleid, casterid)
	
	local skillid = 124	-- 技能id

	local level = 10
	
	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end
function SK_125_Action(battleid, casterid)
	
	
	local skillid = 125	-- 技能id

	local level = 11
	
	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end
function SK_126_Action(battleid, casterid)
	
	
	local skillid = 126	-- 技能id

	local level = 12
	
	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end
function SK_127_Action(battleid, casterid)
	
	local skillid = 127	-- 技能id

	local level = 13
	
	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end
function SK_128_Action(battleid, casterid)
	
	
	local skillid = 128	-- 技能id

	local level = 14
	
	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end
function SK_129_Action(battleid, casterid)
	
	
	local skillid = 129	-- 技能id

	local level = 15
	
	local  attackNum = 0
	
	All_2_Skill(battleid,casterid,attackNum,level,skillid)

end
sys.log("主角 主动技能 2 结束")