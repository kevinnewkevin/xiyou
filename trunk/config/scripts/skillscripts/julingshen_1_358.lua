sys.log(" 巨灵神 SK_358_Action 开始")

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
--巨灵神，横冲直撞：对地方单体造成100%物理伤害， 并且嘲讽目标1回合

function SK_358_Action(battleid,casterid)
	Battle.TargetOn(battleid)
	local skillid = 358
	local attackNum = 0

	local t = Player.GetTarget(battleid,casterid)

	local truedamage = Player.GetUnitDamage(battleid,casterid,t)

	sys.log("巨灵神，横冲直撞目标  "..t.."  造成的物理伤害  "..truedamage )

	local damage = ClacDamageByAllBuff(battleid,casterid,t,truedamage)

	sys.log("巨灵神，横冲直撞对目标  "..t.."  造成的最终物理伤害  "..truedamage )

	if damage <= 0 then

		damage = 0

	end

	local crit = Battle.GetCrit(skillid)
	
	Battle.Attack(battleid,casterid,t,damage,crit)

	Battle.AddBuff(battleid,casterid,t,176,casterid)

	Battle.TargetOver(battleid)
	
	return true
end
sys.log(" 巨灵神 SK_358_Action 结束")