sys.log(" skill 27 start")

-- 技能释放 传入战斗ID和释放者的ID
-- 通过释放者和battleid取得对应的目标 单体或者多个
-- 循环/直接使用接口操控战斗 类似 战斗.攻击(战斗id, 释放者id, 承受者ID, 伤害数值, 是否暴击)
-- 
-- 
-- 所需接口
--	取得目标 （GetTarget（）  单   GetTargets（）  复）
--  取得对应属性 GetUnitProperty()
--  计算伤害数值 demage
--  计算是否暴击
--  攻击
-- 龙王3号技能 对敌方全体造成50%法术强度的伤害，并清除所有敌方负向状态，每清除一个负向状态额外造成50%法术强度的伤害。

-- 法术强度视作buff  Battle.buff

function SK_126_Action(battleid, casterid)

	local skillid = 126	-- 技能id

	local  attackNum = 0  --攻击个数

	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	--local  caster_attack = Player.GetUnitMtk(battleid,casterid)  --获取攻击者属性  fashu 
	
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
		
		local  magic_damage = Player.GetMagicDamage(battleid,casterid,v)  --法术伤害
		
		local  damage  = magic_damage --伤害 公式（ ）
		
		sys.log("SK_126_Action 的伤害"..damage)
		
		local trueDamage =  ClacDamageByAllBuff(battleid,casterid,v,damage)
		
		--判断伤害
		if trueDamage <= 0 then 
		
			trueDamage = 1
		
		end
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		debuffnum = Player.PopAllBuffByDebuff(battleid,v)
			
		local pvalue = 0.5
	
		if debuffnum >0 then 
			
			local demage = magic_damage*pvalue  --额外造成50%法术强度的伤害
			
			for a=1,debuffnum,1 do
			
				trueDamage = trueDamage*pvalue + demage
	
			end
		end
		
		Battle.Attack(battleid,casterid,v,trueDamage,crit)
		
		Battle.TargetOver(battleid)
		
		sys.log("skill26 对id为"..v.."的目标减少"..trueDamage.."点伤害")
		
		
	end
	
	return  true
	 
	 
end

sys.log( "skill 27 end")