sys.log(" 魔礼青 SK_367_Action 开始")

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
--火：对地方全体造成30%法术伤害，无视法术防御

function SK_367_Action(battleid,casterid)
	
	local skillid = 367
	local attackNum = 0
	local t = Player.GetTargets(battleid,casterid,attackNum)

	for  i,v in ipairs(t)do

		Battle.TargetOn(battleid)
	
		local truedamage = Player.GetMagicDamage(battleid,casterid,v)

		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		
		if damage <= 0 then

			damage = 0
			
		end

		damage = damage * 0.9

		local crit = Battle.GetCrit(skillid)

		Battle.Attack(battleid,casterid,v,damage,crit)

		Battle.TargetOver(battleid)
	end
	return true
end
sys.log(" 魔礼青 SK_367_Action 结束")