sys.log("skill 1 start")

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
-- 猴子一号技能 对单体造成本体攻击力130%的伤害, 并且增加本体速度5%
-- 增加速度视作buff

function SK_100_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	local skillid = 100		-- 技能id
	local skillAttack = 10	-- 技能攻击
	
	local t = Player.GetTarget(battleid, casterid)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	sys.log("GetTarget  ".. t)
	
	local caster_attack = Player.GetUnitProperty(battleid, casterid, "CPT_ATK")	-- 获取到攻击者的属性
	local defender_def = Player.GetUnitProperty(battleid, t, "CPT_DEF")
		
	--local damage = caster_attack * 1.3 - defender_def		-- 伤害公式
	local damage = Player.GetUnitDamage(battleid, casterid, t)
	sys.log(1)
	if damage <= 0 then 
		damage = 1
	end
	sys.log(2)
	local crit = Battle.GetCrit(skillid)   --是否暴击
	sys.log("是否暴击"..crit)
	Battle.Attack(battleid, casterid, t, damage, crit)
	Battle.AddBuff(battleid, casterid, t, 10, 5)
	Battle.TargetOver(battleid)
	
	sys.log("skill1 对id为"..t.."的目标造成"..damage.."点伤害")	
	-- 只给游戏返回 对谁造成了多少伤害
	-- 并不参与计算
	return true
end

sys.log("skill 1 end")