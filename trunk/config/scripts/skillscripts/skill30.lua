sys.log(" skill 30 start")

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
-- 蛟魔王3号技能 对敌方一个目标造成物理强度的伤害，并解除所有负面效果，每解除一个负面效果进行一次额外攻击，额外攻击造成50%物理强度的伤害。

-- 法术强度视作buff  Battle.buff

function SK_129_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	local skillid = 129		-- 技能id
	
	local  t = Player.GetTarget(battleid,casterid)  --获取目标 

	local  caster_attack = Player.GetUnitAtk(battleid,casterid)  --获取攻击者属性
		
	local defender_def = Player.GetCalcDef(battleid,t)   -- 防御
	
	local  damage = caster_attack-defender_def
		
	debuffnum = Player.PopAllBuffByDebuff(battleid,t)
		
		--判断伤害
	if damage <= 0 then 
		
		damage = 1
		
	end
	local crit = Battle.GetCrit(skillid)   --是否暴击
	
	
	
	if debuffnum > 0 then 
		demage = int(caster_attack / 2) --额外造成50%物理强度的伤害
		
		for a=1,debuffnum,1 do
		
			--Battle.Attack(battleid,casterid,t,demage,crit)
			
			damage = damage + demage
	
		end
	
	end
	
	Battle.Attack(battleid,casterid,t,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
	
	Battle.TargetOver(battleid)
	
	sys.log("skil30 对id为"..t.."的目标造成"..damage.."点伤害")
	
	return  true
	 
end

sys.log( "skill 30 end")