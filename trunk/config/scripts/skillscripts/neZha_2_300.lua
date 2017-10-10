sys.log("哪吒 SK_300_Action 开始")

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
-- 火尖枪。突刺敌方单体目标（主角）5次，每次造成30%物理强度的伤害，并减少目标10%的物理防御，持续2回合。
-- 增加速度视作buff

function SK_300_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	local skillid = 300		-- 技能id
	local skillAttack = 10	-- 技能攻击
	--local attackNum = 0		-- 攻击个数
	
	local t = Player.GetMainTarget(battleid, casterid)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	local  truedamage = Player.GetUnitDamage(battleid,casterid,t)  --获取伤害
	
	sys.log("哪吒火尖枪对主角目标"..t.."造成的物理伤害"..truedamage)
	local  damage = ClacDamageByAllBuff(battleid,casterid,t,truedamage)
	sys.log("哪吒火尖枪对主角目标"..t.."造成的最终物理伤害"..damage)
	
	local  defender_def = Player.GetCalcDef(battleid,t)  --获取被攻击者的物理防御
	
	local crit = Battle.GetCrit(skillid)    --是否暴击
	
	local damage_atk = damage * 1.5
	
	local atk_del = defender_def * 0.1
	
	Battle.Attack(battleid, casterid, t, damage_atk, crit)
	
	Battle.AddBuff(battleid,casterid,t,112,atk_del)
	Battle.TargetOver(battleid)
	
	return 1
end
sys.log("哪吒 SK_300_Action 结束")
