sys.log("SK_1_Action")

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
--  普通单体攻击技能  物理
-- 增加速度视作buff

function SK_1_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	local skillid = 1	-- 技能id
	
	local lock_buff = Player.GetCheckSpec(battleid,casterid,"BF_LOCK")
	
	if lock_buff == 1 then
		t = Player.GetSpecialData(battleid,casterid,"BF_LOCK")
	else 
		t =  Player.GetTarget(battleid, casterid)
	end

	local truedamage = Player.GetUnitDamage(battleid, casterid, t)
	


	local damage = ClacDamageByAllBuff(battleid,casterid,t,truedamage)
	
	sys.log("攻击者的普通物理伤害    "..damage)
	
	if damage <= 0 then 
		damage = 1
	end
	
	local crit = Battle.GetCrit(skillid)   --是否暴击
	

	Battle.Attack(battleid, casterid, t, damage, crit)
	
	
	Battle.TargetOver(battleid)
	
	
	return true
end

