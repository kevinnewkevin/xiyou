sys.log(" 巨灵神 SK_359_Action 开始")

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
--巨灵神 夺命重击：对地方后排目标造成50%物理伤害，并且迫使这些目标在下一回合强制攻击自己

function SK_359_Action(battleid,casterid)
	
	local skillid = 359
	local attackNum = 0
	local t = Player.BackTarget	(battleid,casterid)

	for  i,v in ipairs(t)do

		Battle.TargetOn(battleid)
	
		local truedamage = Player.GetUnitDamage(battleid,casterid,v)

		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		
		if damage <= 0 then

			damage = 0

		end

		damage = damage * 0.5

		local crit = Battle.GetCrit(skillid)

		Battle.Attack(battleid,casterid,v,damage,crit)

		Battle.AddBuff(battleid,casterid,v,177,casterid)
		
		Battle.TargetOver(battleid)
	end
	return true
end
sys.log(" 巨灵神 SK_359_Action 结束")