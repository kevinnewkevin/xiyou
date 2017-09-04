sys.log(" skill 18 start")

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
-- 观音3号技能 驱散我方全体所有负面效果，并且减少伤害20%，持续1回合。

-- 物理强度视作buff Battle.buff

function SK_117_Action(battleid, casterid)
	
	local skillid = 117		-- 技能id

	local  attackNum = 0   --攻击个数

	local  t = Player.GetFriends(battleid,casterid,attackNum)  --获取目标
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
		Player.PopAllBuffByDebuff(battleid,v)
		local  attack_damage = Player.GetUnitDamage(battleid,casterid,v)  --获取伤害
		Battle.Cure(battleid,v,0,0)
		Battle.AddBuff(battleid,casterid,v,123, attack_damage*0.2)      --公式(减少20%的伤害）
		Battle.TargetOver(battleid)
	
		sys.log("skill18")
	end
	
	
	return  true
	 
end

sys.log( "skill 18 end")