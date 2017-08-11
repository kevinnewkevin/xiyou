sys.log(" skill 6 start")

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
-- 姜子牙3号技能 吸收场上所有护盾，对敌方单体造成法术强度+护盾吸收值*3的伤害。

-- 法术强度视作buff

function  skill_6_Action (battleid, casterid)

	local  num = 0

	local  p = player.GetTargets(battleid,casterid,num)  --获取目标
	
	local  _property = player.GetUnitProperty(battleid,casterid,"CPT_ATK")  --获取攻击者属性
	
	
	for i,v in ipairs(p) do
	
		--local  p_property = Battle.AddBuff(1)  --吸收场上所有盾牌（暂时么有这个函数）
		
	end
	
	for i,v in ipairs(p) do
	
		local defender_def = Player.GetUnitProperty(battleid, v, "CPT_DEF")  --获取防御
		
		--local  spell = Battle.AddBuff(1)   （暂时么有这个函数）  法术强度
	
		local  demage  = spell+ p_property*3-defender_def    --伤害 公式
	
		--判断伤害
		if demage <= 0 then 
		
			demage = 1
		
		end
		
		Battle.Attack(battleid,casterid,demage,true)   --调用服务器   （伤害）
	end
	
	return  true
	 
end

sys.log( "skill 6 end")