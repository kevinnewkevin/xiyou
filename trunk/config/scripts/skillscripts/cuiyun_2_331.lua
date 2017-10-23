sys.log(" 翠云 SK_331_Action 开始")

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
--对目标造成150%物理伤害，并使目标中毒，每回合掉当前生命值的10%血量
function SK_331_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	local skillid = 331
	local t = Player.GetTarget(battleid,casterid)
	local  truedamage  = Player.GetUnitDamage(battleid,casterid,t)    --伤害 公式（）
	
	sys.log("翠云 碧云决对目标造成的法术伤害   ".. truedamage)

	local damage = ClacDamageByAllBuff(battleid,casterid,t,truedamage)

	sys.log("翠云 碧云决对目标造成的最终法术伤害   ".. damage)
	--判断伤害
	if damage <= 0 then 
		damage = 0
	end
	local crit = Battle.GetCrit(skillid)   --是否暴击
	damage = damage * 1.5
	Battle.Attack(battleid,casterid,t,damage,crit)

	--掉当前血量的10%
	local hp = Player.GetUnitProperty(battleid, casterid, "CPT_CHP")
	local hp_damage = hp * 0.1
	Battle.AddBuff(battleid,casterid,t,150,hp_damage)

	Battle.TargetOver(battleid)

	return ture 
end
sys.log("翠云 SK_331_Action 结束")