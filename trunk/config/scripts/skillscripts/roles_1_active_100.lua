sys.log("主角 主动技能 1 开始")

--主角1技能  对敌方全体造成物理攻击

function SK_100_Action(battleid, casterid)
	
	local skillid = 100	-- 技能id

	local level = 1
	local attackNum = 0		-- 攻击个数
	
	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标

	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid) -- 清空数据
	
	    local  truedamage  = Player.GetUnitDamage(battleid,casterid,v)       --伤害 公式（）
		
		sys.log("主角对目标造成的物理伤害    "..truedamage)
		
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		sys.log("主角对目标造成的最终物理伤害    "..truedamage)
		--判断伤害
		if damage <= 0 then 
		
			damage = 0
		
		end
		
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		
		Battle.TargetOver(battleid)  -- 赋给下个目标
		
	end

end

sys.log("主角 主动技能 1  结束")