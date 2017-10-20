sys.log(" 龙后 SK_329_Action 开始")

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
-- 水龙卷。攻击目标单体，对目标造成150%法术伤害，并冰冻目标2回合。
-- 增加速度视作buff

function SK_329_Action(battleid, casterid)

	Battle.TargetOn(battleid)

	local skillid = 329		-- 技能id
	
	local  t = Player.GetTarget(battleid,casterid)  --获取目标 
	
	local  truedamage  = Player.GetMagicDamage(battleid,casterid,t)    --伤害 公式（）
	
	sys.log("龙后 水龙卷对目标造成的法术伤害   ".. truedamage)

	local damage = ClacDamageByAllBuff(battleid,casterid,t,truedamage)

	sys.log("龙后 水龙卷对目标造成的最终法术伤害   ".. damage)
	
	--判断伤害
	if damage <= 0 then 
	
		damage = 0
	
	end
	
	local crit = Battle.GetCrit(skillid)   --是否暴击
	
	damage = damage * 1.5

	Battle.Attack(battleid,casterid,t,damage,crit)   --调用服务器   （伤害）(战斗者，释放者，承受者，伤害，暴击）
	if percent() <= 20 then 

			Battle.AddBuff(battleid,casterid,v,148,0)

		end 
	Battle.TargetOver(battleid)
	
	return  true
	 
end
sys.log(" 龙后 SK_329_Action 结束")