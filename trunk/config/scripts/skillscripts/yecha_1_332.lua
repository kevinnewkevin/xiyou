sys.log(" 夜叉 SK_332_Action 开始")

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
--将目标击退，有50%几率让目标下回合不能行动
function SK_332_Action(battleid,casterid)
	Battle.TargetOn(battleid)

	local skillid = 332

	local  t = Player.GetTarget(battleid,casterid)

	local truedamage = Player.GetUnitDamage(battleid,casterid,t)

	sys.log("夜叉 1号技能  对目标  "..t..  "造成的伤害"..truedamage)

	local damage = ClacDamageByAllBuff(battleid,casterid,t,truedamage)

	sys.log("夜叉 1号技能  对目标  "..t..  "造成的最终伤害"..truedamage)

	if damage <= 0 then 

		damage = 0

	end

	local crit = Battle.GetCrit(skillid)

	Battle.Attack(battleid,casterid,t,damage,crit)

	--50%的几率无法行动
	if percent() <= 50 then

		Battle.AddBuff(battleid,casterid,t,167,0)

	end	

	Battle.TargetOver(battleid)

	return true

end
sys.log(" 夜叉 SK_332_Action 结束")