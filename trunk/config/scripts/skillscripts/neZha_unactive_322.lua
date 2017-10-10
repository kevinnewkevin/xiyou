sys.log("哪吒 被动技能 SK_322_Action 开始")

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
--三头六臂。装备获得的连击属性加倍。
 


function SK_322_Action(battleid, casterid)
	
	local skillid = 322		-- 技能id
	Battle.TargetOn(battleid)

		
		
	Battle.TargetOver(battleid)
	
	return  true
	 
end
sys.log("哪吒 被动技能 SK_322_Action 结束")