sys.log(" 王母 SK_356_Action 开始")

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
--王母，主持公道：对地方单体目标造成100%法术伤害，并且沉默2回合。沉默阶段只能释放普通攻击

function SK_356_Action(battleid,casterid)
	Battle.TargetOn(battleid)
	local skillid = 356
	local attackNum = 0

	local t = Player.GetTarget(battleid,casterid)

	local truedamage = Player.GetMagicDamage(battleid,casterid,t)

	sys.log("王母，主持公道对目标  "..t.."  造成的法术伤害  "..truedamage )

	local damage = ClacDamageByAllBuff(battleid,casterid,t,truedamage)

	sys.log("王母，主持公道对目标  "..t.."  造成的最终法术伤害  "..truedamage )

	if damage <= 0 then

		damage = 0

	end

	local crit = Battle.GetCrit(skillid)
	
	Battle.Attack(battleid,casterid,t,damage,crit)

	Battle.AddBuff(battleid,casterid,t,175,0)

	Battle.TargetOver(battleid)
	
	return true
end
sys.log(" 王母 SK_356_Action 结束")