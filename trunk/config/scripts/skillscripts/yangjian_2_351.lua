sys.log(" 杨戬 SK_351_Action 开始")

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
-- 神目：对目标造成100%的法术伤害，无视防御
-- 增加速度视作buff

function SK_351_Action(battleid, casterid)

	Battle.TargetOn(battleid)

	local skillid = 351		-- 技能id
	
	local  t = Player.GetTarget(battleid,casterid)  --获取目标 
	
	local  truedamage  = Player.GetUnitDamage(battleid,casterid,t)    --伤害 公式（）
	
	sys.log("杨戬 神目对目标造成的物理伤害   ".. truedamage)

	local damage = ClacDamageByAllBuff(battleid,casterid,t,truedamage)

	sys.log("杨戬 神目对目标造成的最终物理伤害   ".. damage)
	
	--判断伤害
	if damage <= 0 then 
	
		damage = 0
	
	end
	
	local crit = Battle.GetCrit(skillid)   --是否暴击
	
	damage = damage * 3

	Battle.Attack(battleid,casterid,t,damage,crit)   --调用服务器   （伤害）(战斗者，释放者，承受者，伤害，暴击）
	
	Battle.TargetOver(battleid)
	
	return  true
	 
end
sys.log(" 杨戬 SK_351_Action 结束")