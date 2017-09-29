sys.log(" skill 44 start")

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
--被动技能：鼓舞。女娲在场时，每回合为所有友方单位增加女娲10%的法术强度，可叠加。
 


function SK_324_Action(battleid, casterid)
	
	local skillid = 324		-- 技能id
	
	local attackNum = 0
	
	local t = Player.GetFriends(battleid,casterid,attackNum)
	
	local mtk =  Player.GetUnitMtk(battleid,casterid)
	
	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid)
		
		local mag_atk =  mtk * 0.1
		
		Battle.Cure(battleid,casterid,0,0)
		
		Battle.AddBuff(battleid,casterid,v,140,mag_atk)
			
		Battle.TargetOver(battleid)
	
	end
	
	return  true
	 
end
sys.log(" skill 44 end")