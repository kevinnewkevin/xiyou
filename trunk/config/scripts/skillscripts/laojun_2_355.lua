sys.log(" 太上老君 SK_355_Action 开始")

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
--太上老君 道德经：提升己方所有目标荆棘值100点，持续2回合

function SK_355_Action(battleid,casterid)
	
	local skillid = 355
	local attackNum = 0
	local t = Player.GetFriends(battleid,casterid,attackNum)

	for  i,v in ipairs(t)do
		Battle.TargetOn(battleid)
		Battle.Cure(battleid,v,0,0)
		sys.log("太上老君 道德经给己方加一个免伤的buff 开始")
		Battle.AddSkillBuff(battleid,casterid,casterid,174,100)
		sys.log("太上老君 道德经给己方加一个免伤的buff 结束")
		Battle.TargetOver(battleid)
	end
	return true
end
sys.log(" 太上老君 SK_355_Action 结束")