sys.log(" skill 17 start")

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
-- 观音2号技能 为一个友方目标回复法术强度的生命值，并且使其每次受到伤害都会回复10%法术强度的生命值，持续1回合。

-- 法术强度视作buff  Battle.buff

function SK_116_Action(battleid, casterid)
	Battle.TargetOn(battleid)

	local skillid = 116		-- 技能id
	
	local  t = Player.GetFriend(battleid,casterid)  --获取目标 
	
	local  caster_attack = Player.GetUnitMtk(battleid,casterid)  --获取攻击者属性
	
	local crit = Battle.GetCrit(skillid)

	Battle.Cure(battleid, t, caster_attack, crit)    --为一个友方目标回复法术强度的生命值
	
	Battle.AddBuff(battleid,casterid,t,121,caster_attack*0.1) --每回合伤害回复法术强度10%的生命值
	
	Battle.TargetOver(battleid)
	
	
	sys.log("skill17")
	
	return  true
	 
end

sys.log( "skill 17 end")