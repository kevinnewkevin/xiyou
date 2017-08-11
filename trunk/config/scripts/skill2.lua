sys.log("skill 2 start")

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
-- 猴子2号技能 对三个目标造成150%伤害并减少对方双防15%，持续2回合。
-- 增加速度视作buff

function SK_2_Action(battleid, casterid)
	local skillid = 2		-- 技能id
	local skillAttack = 10	-- 技能攻击
	local attackNum = 3		-- 攻击个数
	
	local t = Player.GetTargets(battleid, casterid, attackNum)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	local caster_attack = Player.GetUnitProperty(battleid, casterid, "CPT_ATK")	-- 获取到攻击者的属性
	
	sys_log(t)
	for i,v in ipairs(t)	do
		local defender_def = Player.GetUnitProperty(battleid, v, "CPT_DEF")
		local damage = caster_attack * 1.5 - defender_def		-- 伤害公式
	
		if damage <= 0 then 
			damage = 1
		end
	
		Battle.Attack(battleid, casterid, v, damage, true)
		
		sys.log("skill2 对id为"..v.."的目标造成"..damage.."点伤害")
	end
	
	return 1
end

sys.log("skill 2 end")