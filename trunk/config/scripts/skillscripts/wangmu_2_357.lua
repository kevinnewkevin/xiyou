sys.log(" 王母 SK_357_Action 开始")

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
--王母 母仪天下：恢复己方所有目标最大生命值20%的生命值，并且使这些目标再接下来的2回合内，每回合恢复10%生命值

function SK_357_Action(battleid,casterid)
	
	local skillid = 357
	local attackNum = 0
	local t = Player.GetFriends(battleid,casterid,attackNum)

	for  i,v in ipairs(t)do
		Battle.TargetOn(battleid)

		local max_hp = Player.GetUnitProperty(battleid,v,"CPT_HP") 

		local crit= Battle.GetCrit(skillid)

		max_hp = max_hp * 0.2 

		Battle.Cure(battleid,v,max_hp,crit)

		hp = max_hp * 0.1

		Battle.AddBuff(battleid,casterid,v,101,hp)
		
		Battle.TargetOver(battleid)
	end
	return true
end
sys.log(" 王母 SK_357_Action 结束")