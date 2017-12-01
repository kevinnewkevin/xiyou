sys.log(" 太白金星 SK_348_Action 开始")

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
--太白金星1 号 技能，议和：降低对方单体双攻100%，持续1回合

function SK_348_Action(battleid,casterid)
	Battle.TargetOn(battleid)
	local skillid = 348

	local t = Player.GetTarget(battleid,casterid)

	local mtk = Player.GetUnitMtk(battleid,t)

	local atk = Player.GetUnitAtk(battleid,t)

	sys.log("太白金星1 号"..mtk.."太白金星1 号"..atk)

	Battle.Attack(battleid,casterid,t,0,0)

	Battle.AddBuff(battleid,casterid,t,171,atk)

	Battle.AddBuff(battleid,casterid,t,172,mtk)
	
	Battle.TargetOver(battleid)

	return true
end
sys.log(" 太白金星 SK_348_Action 结束")