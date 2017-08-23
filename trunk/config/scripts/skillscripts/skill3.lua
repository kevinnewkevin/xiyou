sys.log("skill 3 start")

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
-- 猴子3号技能 对所有敌人造成230%的伤害并有几率眩晕对方1回合，并有20%几率造成双倍伤害。
-- 增加速度视作buff

-- 所有有概率的技能在使用概率之前都要使用这个函数,作用是通过服务器的时间戳来获取到一个随机种子
set_random_seed()

function SK_102_Action(battleid, casterid)
	local skillid = 102		-- 技能id
	local skillAttack = 10	-- 技能攻击
	local attackNum = 0		-- 攻击个数
	
	local t = Player.GetTargets(battleid, casterid, attackNum)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	local caster_attack = Player.GetUnitProperty(battleid, casterid, "CPT_ATK")	-- 获取到攻击者的属性
	
	sys_log(t)
	for i,v in ipairs(t)	do
		local defender_def = Player.GetUnitProperty(battleid, v, "CPT_DEF")
		local damage = caster_attack * 2.3 - defender_def		-- 伤害公式
		
		if percent() <= 20 then 						-- 低于20%就会触发双倍伤害
			damage = damage * 2
		end
	
		if damage <= 0 then 
			damage = 1
		end
		sys.log(1)
		local crit = Battle.GetCrit(skillid)   --是否暴击
		Battle.Attack(battleid, casterid, v, damage, crit)
		Battle.AddBuff(battleid, casterid, v, 3, 5)
		
		sys.log("skill3 对id为"..v.."的目标造成"..damage.."点伤害")
	end
	
	return 1
end

sys.log("skill 3 end")