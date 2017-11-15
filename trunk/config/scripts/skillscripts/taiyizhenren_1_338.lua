sys.log(" 太乙真人 SK_338_Action 开始")

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
--太乙真人1号技能 磨砺：对友方所有目标（除自己）造成当前生命值10%的伤害，并且使这些目标下次攻击额外造成30%伤害

function SK_338_Action(battleid,casterid)

	local skillid = 338
	local attackNum = 0
	local t = Player.GetFriends(battleid,casterid,attackNum)

	for v,k in ipairs(t) do
		if k == casterid then
			k = Player.GetFriend(battleid,casterid)
		end
		Battle.TargetOn(battleid)

		local hp = Player.GetUnitProperty(battleid,k,"CPT_CHP")

		local truedamage = hp * 0.1

		local demage = ClacDamageByAllBuff(battleid,casterid,k,truedamage)
		
		local crit = Battle.GetCrit(skillid)

		Battle.Attack(battleid,casterid,k,demage,crit)

		Battle.AddBuff(battleid,casterid,k,169,30)

		Battle.TargetOver(battleid)

	end
	return true
end
sys.log(" 太乙真人 SK_338_Action 结束")
