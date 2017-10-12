sys.log("猴子 SK_296_Action 开始")

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
-- 猴子 2号 大闹天宫 召唤金箍棒变大从天砸下，对敌方全体目标造成230%的伤害并有20%几率眩晕对方1回合，并有20%几率造成双倍伤害。
-- 增加速度视作buff
set_random_seed()
function SK_296_Action(battleid, casterid)
	local skillid = 296		-- 技能id
	local skillAttack = 10	-- 技能攻击
	local attackNum = 0		-- 攻击个数
	
	local t = Player.GetTargets(battleid, casterid, attackNum)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	local caster_attack = Player.GetUnitProperty(battleid, casterid, "CPT_ATK")	-- 获取到攻击者的属性
	
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
		--local defender_def = Player.GetUnitProperty(battleid, v, "CPT_DEF")
		--local damage = caster_attack * 2.3 - defender_def		-- 伤害公式
		local truedamage = Player.GetUnitDamage(battleid, casterid, v)
		
		sys.log("猴子 大闹天宫的物理伤害   "..truedamage)
		
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		
		sys.log("猴子大闹天宫 最终的物理伤害   "..damage)
		
		if percent() <= 20 then 						-- 低于20%就会触发双倍伤害
			damage = damage * 2
		end
		
		if damage <= 0 then
			damage = 0
		end
		
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		local atk_damage = damage*0.8
		
		Battle.Attack(battleid, casterid, v, atk_damage, crit)
		
		if percent() <= 20 then 						-- 眩晕对方1回合
			Battle.AddBuff(battleid, casterid, v, 104, 0)
		end
		if crit == 1 then 
			Battle.AddSkillBuff(battleid,casterid,casterid,138,50)
		end
		Battle.TargetOver(battleid)
		
		
	end
	
	return 1
end
sys.log("猴子 SK_296_Action 结束")