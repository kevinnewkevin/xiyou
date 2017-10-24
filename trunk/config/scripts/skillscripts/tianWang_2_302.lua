sys.log("天王 SK_302_Action 开始")

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
-- 玲珑塔。敌方单体2回合无法行动，并对目标造成法术强度的伤害。
-- 增加速度视作buff

function SK_302_Action(battleid, casterid)

	Battle.TargetOn(battleid)
	local skillid = 302		-- 技能id
	local skillAttack = 10	-- 技能攻击
	--local attackNum = 0		-- 攻击个数
	
	local t = Player.GetTarget(battleid, casterid)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	local  truedamage  = Player.GetMagicDamage(battleid,casterid,t)       --伤害 公式（）
	sys.log("天王 玲珑塔对目标造成的物理伤害    "..truedamage)	
	local damage = ClacDamageByAllBuff(battleid,casterid,t,truedamage)
	sys.log("天王  玲珑塔对目标造成的最终物理伤害    "..truedamage)
	--判断伤害
	if damage <= 0 then 
		
		damage = 0
		
	end
		
	local crit = Battle.GetCrit(skillid)   --是否暴击
	
	Battle.Attack(battleid,casterid,t,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
	
	Battle.AddBuff(battleid,casterid,t,166,0)
	
	Battle.TargetOver(battleid)
	return  true
end

sys.log("天王 SK_302_Action 结束")