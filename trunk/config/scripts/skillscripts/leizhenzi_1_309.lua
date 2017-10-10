sys.log("雷震子 SK_309_Action 开始")

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
-- 雷击。对敌方血量最高的单体造成法术强度的伤害。并且附带持续掉血效果，每回合掉当前生命值40%血量。持续2回合
-- 增加速度视作buff

function SK_309_Action(battleid, casterid)

	Battle.TargetOn(battleid)

	local skillid = 309		-- 技能id
	
	local  t = max_property_one(battleid,casterid,"CPT_CHP")  --获取目标 
	
	local  truedamage  = Player.GetMagicDamage(battleid,casterid,t)    --伤害 公式（）
	sys.log("雷震子 雷击对目标   "..t.. " 造成 法术伤害  "..truedamage )
	
	local damage = ClacDamageByAllBuff(battleid,casterid,t,truedamage)
	sys.log("雷震子 雷击对目标   "..t.. " 造成最终 法术伤害  "..damage )
	--判断伤害
	if damage <= 0 then 
	
		damage = 1
	
	end
	
	local crit = Battle.GetCrit(skillid)   --是否暴击
	
	local hp = Player.GetUnitProperty(battleid, casterid, "CPT_CHP")
	local Bhp = Player.GetUnitProperty(battleid, casterid, "CPT_HP")
	local pvalue = 0.4
	local hp_damage = hp*pvalue
	local  hp_pro= Bhp*0.1
	damege = damage + hp_pro
	Battle.Attack(battleid,casterid,t,damage,crit)   --调用服务器   （伤害）(战斗者，释放者，承受者，伤害，暴击）
	sys.log("雷震子 雷击对目标   "..t.. "加掉血buff" )
	Battle.AddBuff(battleid,casterid, t,100, hp_damage)     --每回合掉当前生命值40%血量
	sys.log("雷震子 雷击对目标   "..t.. "加被动技能buff" )
	
	Battle.TargetOver(battleid)
	
	
	
	return  true
	 
end

sys.log("雷震子 SK_309_Action 结束")