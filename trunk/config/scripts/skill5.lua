sys.log(" skill 5 start")

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
-- 猴子2号技能 为周围的友方单位分别提供一个护盾，护盾抵挡的伤害值相当于姜子牙法术强度的10%。

-- 法术强度视作buff Battle.buff

function  skill_5_Action (battleid, casterid)

	local  num = 0

	local  p = player.GetTargets(battleid,casterid,num)  --获取目标
	
	local  _property = player.GetUnitProperty(battleid,v,"CPT_ATK")  --获取攻击者属性
	
	--local  p_property = Battle.AddBuff(1) 攻击者法术
	
	for i,v in ipairs(p) do
		local defender_def = Player.GetUnitProperty(battleid, v, "CPT_DEF")
	
		--local  p_property = Battle.AddBuff(1)  --给友方分别提供一个盾牌（暂时么有这个函数）
	
		local  demage  = p_property*0.1-defender_def  --伤害 公式（）
	
		--判断伤害
		if demage <= 0 then 
		
			demage = 1
		
		end
		
		Battle.Attack(battleid,casterid,v,demage,true)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
	end
	
	return  true
	 
end

sys.log( "skill 5 end")