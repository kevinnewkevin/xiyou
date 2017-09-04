sys.log(" skill 25 start")

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
-- 龙王1号技能 标记所有敌方目标，使其受到的伤害增加20%。

-- 法术强度视作buff  Battle.buff

function SK_124_Action(battleid, casterid)

	
	local skillid = 124	-- 技能id

	local  attackNum = 0  --攻击个数

	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
		local  damage = Player.GetUnitDamage(battleid,casterid,v)  --获取攻击者属性
		Battle.Attack(battleid, casterid, v, 0, 0)
		
		Battle.AddBuff(battleid,casterid,v, 110,damage*0.2)
		Battle.TargetOver(battleid)
	
		sys.log("skill25")
	end
	
	
	return  true
	 
	 
end

sys.log( "skill 25 end")