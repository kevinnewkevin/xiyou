sys.log(" 翠云 SK_330_Action 开始")

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
--攻击两次，对目标造成100%物理伤害，并且降低目标防御50点，持续2回合
function SK_330_Action(battleid, casterid)
	Battle.TargetOn(battleid)

	local skillid = 330
	local t = Player.GetTarget(battleid,casterid)
	local  truedamage  = Player.GetUnitDamage(battleid,casterid,t)    --伤害 公式（）
	
	sys.log("翠云 飞叶斩 对目标造成的法术伤害   ".. truedamage)

	local damage = ClacDamageByAllBuff(battleid,casterid,t,truedamage)

	sys.log("翠云 飞叶斩 对目标造成的最终法术伤害   ".. damage)
	
	--判断伤害
	if damage <= 0 then 
	
		damage = 0
	
	end
	
	local crit = Battle.GetCrit(skillid)   --是否暴击

	Battle.Attack(battleid,casterid,t,damage,crit)

	Battle.AddBuff(battleid,casterid,t,112,50)

	Battle.TargetOver(battleid)

	return  true
end
sys.log(" 翠云 SK_330_Action 结束")