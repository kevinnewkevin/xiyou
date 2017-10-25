sys.log("主角 主动技能 1 开始")

--主角1技能  对敌方全体物理攻击，每升一级伤害加5%

--公用函数
function AllSkill(battleid,casterid,attackNum,level,skillid)

	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标

	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid) -- 清空数据
	
	    local  truedamage  = Player.GetUnitDamage(battleid,casterid,v)       --伤害 公式（）
		
		sys.log("主角对目标造成的物理伤害    "..truedamage)
		
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		sys.log("主角对目标造成的最终物理伤害    "..truedamage)
		--判断伤害
		if damage <= 0 then 
		
			damage =0
		
		end
		
		local crit = Battle.GetCrit(skillid)   --是否暴击

		sya.log("主角造成的伤害111    ".. damage)

		local per = 0.05 * level

		damage = damage + damage *per

		sya.log("主角造成的伤害222     ".. damage)
		
		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		
		Battle.TargetOver(battleid)  -- 赋给下个目标
		
	end
end

function SK_100_Action(battleid, casterid)
	local skillid = 100	-- 技能id

	local level = 1
	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)

end
--主角1技能  增加自己所有属性

function SK_101_Action(battleid, casterid)
	local skillid = 101	-- 技能id

	local level = 2
	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)

end
function SK_102_Action(battleid, casterid)
	local skillid = 102	-- 技能id

	local level = 3
	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)

end



function SK_103_Action(battleid, casterid)
	local skillid = 103	-- 技能id

	local level = 4
	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)

end



function SK_104_Action(battleid, casterid)
	local skillid = 104	-- 技能id

	local level = 5
	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)
end



function SK_105_Action(battleid, casterid)
	local skillid = 105	-- 技能id

	local level = 6
	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)

end



function SK_106_Action(battleid, casterid)
	
	
	local skillid = 106	-- 技能id

	local level = 7

	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)

end



function SK_107_Action(battleid, casterid)
	
	
	local skillid = 107	-- 技能id

	local level = 8
	
	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)

end

--主角1技能  增加自己所有属性

function SK_108_Action(battleid, casterid)
	
	
	local skillid = 108	-- 技能id

	local level = 9
	
	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)

end

--主角1技能  增加自己所有属性

function SK_109_Action(battleid, casterid)
	
	
	local skillid = 109	-- 技能id

	local level = 10
	
	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)

end

--主角1技能  增加自己所有属性

function SK_110_Action(battleid, casterid)
	
	
	local skillid = 110	-- 技能id

	local level = 11
	
	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)

end

--主角1技能  增加自己所有属性

function SK_111_Action(battleid, casterid)
	
	
	local skillid = 111	-- 技能id

	local level = 12
	
	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)

end

--主角1技能  增加自己所有属性

function SK_112_Action(battleid, casterid)
	
	
	local skillid = 112	-- 技能id

	local level = 13
	
	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)

end

--主角1技能  增加自己所有属性

function SK_113_Action(battleid, casterid)
	
	local skillid = 113	-- 技能id

	local level = 14
	
	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)

end

--主角1技能  增加自己所有属性

function SK_114_Action(battleid, casterid)
	
	local skillid = 114	-- 技能id

	local level = 15
	
	local attackNum = 0		-- 攻击个数
	
	AllSkill(battleid,casterid,attackNum,level,skillid)

end
sys.log("主角 主动技能 1  结束")