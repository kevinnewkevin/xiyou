sys.log(" 太上老君 SK_354_Action 开始")

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
--太上老君，明道：提升己方全体目标的速度100%，持续1回合

function SK_354_Action(battleid,casterid)
	
	local skillid = 354
	local attackNum = 0

	local t = Player.GetFriends(battleid,casterid,attackNum)

	for i,v in ipairs(t)do
		Battle.TargetOn(battleid)
		Battle.Cure(battleid,v,0,0)
		sys.log("太上老君，明道给己方加一个速度的buff 开始")
		local sudu = Player.GetUnitProperty(battleid,v,"CPT_AGILE")
		Battle.AddBuff(battleid,casterid,v,173,sudu)
		sys.log("太上老君，明道给己方加一个速度的buff 结束")
		Battle.TargetOver(battleid)
	end
	
	return true
end
sys.log(" 太上老君 SK_354_Action 结束")