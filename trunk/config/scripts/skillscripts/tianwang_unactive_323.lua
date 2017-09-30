sys.log(" SK_323_Action")

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
--护法。场上每个友方角色为托塔李天王增加法术强度，增加值等于友方角色法术强度的20%。
 


function SK_323_Action(battleid, casterid)
	
	local skillid = 323		-- 技能id
	
	local attackNum = 0
	
	local t = Player.GetFriends(battleid,casterid,attackNum)
	
	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid)
	
		local mtk =  Player.GetUnitMtk(battleid,v)
		
		local mag_atk =  mtk * 0.2
		
		Battle.Cure(battleid,casterid,0,0)
		
		Battle.AddBuff(battleid,v,casterid,140,mag_atk)
			
		Battle.TargetOver(battleid)
	
	end
	
	
	return  true
	 
end
