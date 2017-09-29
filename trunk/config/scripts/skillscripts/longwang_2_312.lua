sys.log("skill 32 start")

-- 技能释放 传入战斗ID和释放者的ID
-- 通过释放者和battleid取得对应的目标 单体或者多个
-- 循环/直接使用接口操控战斗 类似 战斗.攻击(战斗id, 释放者id, 承受者ID, 伤害数值, 是否暴击)
-- 
-- 
-- 所需接口
--  取得对应属性
--  计算伤害数值
--  计算是否暴击
--  攻击
--消灾。对敌方全体造成50%法术强度的伤害，并清除所有敌方负向状态，每清除一个负向状态额外增加50%法术强度的伤害。伤害只结算一次，根据debuff数量直接算伤害，每消除一个增加50%伤害
-- 增加速度视作buff

function SK_312_Action(battleid, casterid)

	
	local skillid = 312		-- 技能id
	local skillAttack = 10	-- 技能攻击
	local attackNum = 0		-- 攻击个数
	
	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	--local  caster_attack = Player.GetUnitMtk(battleid,casterid)  --获取攻击者属性  fashu 
	
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
		
		local  magic_damage = Player.GetMagicDamage(battleid,casterid,v)  --法术伤害
		sys.log("龙王 消灾对目标  "..v..  "造成法术伤害"..magic_damage)
		local  damage  = magic_damage --伤害 公式（ ）
		sys.log("龙王 消灾对目标  "..v..  "造成法术伤害"..damage)
		local trueDamage =  ClacDamageByAllBuff(battleid,casterid,v,damage)
		sys.log("龙王 消灾对目标  "..v .. "造成法术最终伤害"..trueDamage)
		--判断伤害
		if trueDamage <= 0 then 
		
			trueDamage = 1
		
		end
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		debuffnum = Player.PopAllBuffByDebuff(battleid,v)
		sys.log("龙王 消灾对目标  "..v..  "清除所有负向状态")
		local pvalue = 0.5
	
		if debuffnum >0 then 
			
			local demage = magic_damage*pvalue  --额外造成50%法术强度的伤害
			mag_damage = 0
			for a=1,debuffnum,1 do
			
				mag_damage = mag_damage + demage
	
			end
			trueDamage = trueDamage*pvalue + mag_damage
			
			sys.log("龙王 消灾对目标  "..v .. "清除所有负向状态的伤害"..trueDamage)
		else 
		
			trueDamage = trueDamage*pvalue
			
		end
		sys.log("龙王 消灾对目标  "..v .. " 伤害 "..trueDamage)
		
		Battle.Attack(battleid,casterid,v,trueDamage,crit)
		
		Battle.TargetOver(battleid)
		
		
	end
	
	return  true
end

sys.log("skill 32 end")