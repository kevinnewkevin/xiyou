sys.log(" 妖妃 SK_335_Action 开始")

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
--魅惑目标，使目标下一回合的攻击目标变为己方一个随机目标

function SK_335_Action(battleid,casterid)
	Battle.TargetOn(battleid)
	local skillid = 335
	local t = Player.GetTarget(battleid,casterid)
	Battle.Attack(battleid,casterid,t,0,0)
	sys.log("妖妃 魅惑目标给目标加一个攻击目标为己方的随机目标的buff 开始")
	Battle.AddBuff(battleid,casterid,t,152,0)
	sys.log("妖妃 魅惑目标给目标加一个一个攻击目标为己方的随机目标的buff 结束")
	Battle.TargetOver(battleid)
	return true
end
sys.log(" 妖妃 SK_335_Action 结束")