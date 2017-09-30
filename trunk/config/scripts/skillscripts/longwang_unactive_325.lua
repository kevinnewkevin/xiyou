sys.log("SK_325_Action")

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
--龙威。每减少10%的生命值，增加10%的法术强度。
 


function SK_325_Action(battleid, casterid)
	
	
	Battle.TargetOn(battleid)
	local skillid = 325		-- 技能id
		
		
	Battle.TargetOver(battleid)
	
	return  true
	 
end
