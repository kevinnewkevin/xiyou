sys.log(" skill 41 start")

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
--愿者上钩。所有友方角色增加姜子牙物理强度和法术强度的20%。
 


function SK_321_Action(battleid, casterid)
	
	local skillid = 321		-- 技能id

	local  attackNum = 0   --攻击个数
	
	local level = 1

	local  t = Player.GetFriends(battleid,casterid,attackNum)  --获取目标
	
	local atk = Player.GetUnitAtk(battleid,casterid)
	
	local mtk = Player.GetUnitMtk(battleid,casterid)
	
	local per = 0.2
	
	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid)
		
		local atk_damage = atk * per
		
		local mag_atk = mtk * per
		
		Battle.Cure(battleid,v,0,0)
		
		Battle.AddBuff(battleid,casterid,v,139,atk_damage)
		
		Battle.AddBuff(battleid,casterid,v,140,mag_atk)
		
		Battle.TargetOver(battleid)
	
		
	end
	return  true
	 
end
sys.log(" skill 41 end")