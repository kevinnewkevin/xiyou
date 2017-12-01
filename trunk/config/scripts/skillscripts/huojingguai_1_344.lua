sys.log(" 火精怪 SK_344_Action 开始")

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
--火精怪1 号 技能，对单体目标造成100%法术伤害

function SK_344_Action(battleid,casterid)
	Battle.TargetOn(battleid)
	local skillid = 344

	local t = Player.GetTarget(battleid,casterid)

	local truedamage = Player.GetMagicDamage(battleid,casterid,t)

	local damage = ClacDamageByAllBuff(battleid,casterid,t,truedamage)

	if damage <= 0 then

		damage = 0

	end

	local crit = Battle.GetCrit(skillid)

	Battle.Attack(battleid,casterid,t,damage,crit)
	
	Battle.TargetOver(battleid)

	return true
end
sys.log(" 火精怪 SK_344_Action 结束")