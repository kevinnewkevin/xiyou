sys.log(" skill 28 start")

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
-- 蛟魔王1号技能 使自己受到伤害增加30%，造成的伤害增加30%，对敌方一个目标造成物理强度的伤害，持续3回合。

-- 法术强度视作buff  Battle.buff

function SK_127_Action(battleid, casterid)

	Battle.TargetOn(battleid)
	
	local skillid = 127		-- 技能id
	
	local  t = Player.GetTarget(battleid,casterid)  --获取目标 
	
	local caster_damage = Player.GetUnitDamage(battleid,casterid,t)  -- 获取伤害

	sys.log(1)
	
	--判断伤害
	if caster_damage <= 0 then 
	
		caster_damage = 1
	
	end
	sys.log(5)
	local crit = Battle.GetCrit(skillid)   --是否暴击
	sys.log(3)
	Battle.Attack(battleid,casterid,t,caster_damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
	
	Battle.AddBuff(battleid, casterid,casterid,110, caster_damage*0.3)  --使自己受到伤害增加30%
	
	Battle.AddBuff(battleid,casterid, t, 122,caster_damage*0.3)  --造成的伤害增加30%
	
	sys.log(4)
	Battle.TargetOver(battleid)
	
	sys.log("skil22 对id为"..t.."的目标造成"..caster_damage.."点伤害")
	
	return  true
	 
end

sys.log( "skill 28 end")
