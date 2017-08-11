sys.log(" skill 4 start")

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
-- 姜子牙一号技能 对敌方单目标造成等同物理攻击的伤害，并降低对手等同伤害值的法术强度，为自己增加等量的法术强度。

-- 法术强度视作buff

function  skill_4_Action (battleid, casterid)

	local  p = player.GetTarget(battleid,casterid)  --获取目标
	
	local  _property = player.GetUnitProperty(battleid,casterid,"CPT_ATK")  --获取攻击者属性
	
	local defender_def = Player.GetUnitProperty(battleid, casterid, "CPT_DEF")  --获取防御属性
	
	--local  p_property = Battle.AddBuff(1)   --降低对手同等伤害的法术强度
	
	--local  p_property = Battle.AddBuff(2)   --增加自己等量的法术强度
	
	local  demage  = _property --伤害 公式（攻击属性）
	
	--判断伤害
	if demage <= 0 then 
	
		demage = 1
	
	end
	
	Battle.Attack(battleid,casterid,demage,true)   --调用服务器   （伤害）
	
	return  true
	 
end

sys.log( "skill 4 end")