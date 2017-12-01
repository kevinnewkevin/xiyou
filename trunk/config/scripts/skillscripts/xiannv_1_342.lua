sys.log(" 仙女 SK_342_Action 开始")

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
--仙女1 号 技能，给友方血量最少的单位恢复自身法术强度 的生命值

function SK_342_Action(battleid,casterid)
	Battle.TargetOn(battleid)
	local skillid = 342

	local t = minPropertyOne(battleid,casterid,"CPT_CHP")

	local crit = Battle.GetCrit(skillid)

	local damage = Player.GetUnitMtk(battleid,casterid)

	Battle.Cure(battleid,t,damage,crit)
	
	Battle.TargetOver(battleid)

	return true
end
sys.log(" 仙女 SK_342_Action 结束")