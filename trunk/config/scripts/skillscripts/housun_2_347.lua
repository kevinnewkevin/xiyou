sys.log(" 猴孙 SK_347_Action 开始")

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
--猴孙 2号技能。连击，攻击敌方一个目标造成70%伤害，再次攻击另一个目标，造成70%伤害

function SK_347_Action(battleid,casterid)
	
	local skillid = 347
	local attackNum = 2		-- 攻击个数
	local t = Player.GetTargets(battleid,casterid,attackNum)
	
	 for i,v in ipairs(t)do 

	 	Battle.TargetOn(battleid)

	 	local truedamage = Player.GetUnitDamage(battleid,casterid,v)
	 	sys.log("猴孙 2号对目标造成的物理伤害   ".. truedamage)


	 	local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)

	 	sys.log("猴孙 2号对目标造成的最终物理伤害   ".. damage)

	 	if damage <= 0 then

	 		damage = 0

	 	end
	 	sys.log("猴孙 11111" )
	 	
		damage = damage * 0.7

		sys.log("猴孙 222222" )

	 	local crit = Battle.GetCrit(skillid)

	 	sys.log("猴孙 333333" )

	 	Battle.Attack(battleid,casterid,v,damage,crit)
	 	
	 	Battle.TargetOver(battleid)

	 end

	return true
end
sys.log(" 猴孙 SK_347_Action 结束")