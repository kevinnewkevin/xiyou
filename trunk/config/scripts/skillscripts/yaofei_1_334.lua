sys.log(" 妖妃 SK_334_Action 开始")

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
--楚楚可怜，使自己该回合不受到伤害，免疫一切攻击

function SK_334_Action(battleid,casterid)
	Battle.TargetOn(battleid)
	local skillid = 334
	Battle.Cure(battleid,casterid,0,0)
	sys.log("妖妃 楚楚可怜给自己加一个免伤的buff 开始")
	Battle.AddSkillBuff(battleid,casterid,casterid,151,0)
	sys.log("妖妃 楚楚可怜给自己加一个免伤的buff 结束")
	Battle.TargetOver(battleid)
	return true
end
sys.log(" 妖妃 SK_334_Action 结束")