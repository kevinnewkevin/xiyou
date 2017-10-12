sys.log("猴子 SK_295_Action 开始")

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
-- 猴子  1号  化身三千 召唤2个分身和自己同时攻击对面三个前排目标，造成150%伤害并减少对方双防15%，持续2回合。。
-- 增加速度视作buff

function SK_295_Action(battleid, casterid)
	local skillid = 295		-- 技能id
	local skillAttack = 10	-- 技能攻击
	--local attackNum = 3		-- 攻击个数
	
	local t = Player.FrontTarget(battleid, casterid)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	local caster_attack = Player.GetUnitProperty(battleid, casterid, "CPT_ATK")	-- 获取到攻击者的属性
	
	
	for i,v in ipairs(t)	do
		Battle.TargetOn(battleid)
		local defender_atk = Player.GetUnitProperty(battleid, v, "CPT_DEF")
		local defender_mtk = Player.GetUnitProperty(battleid, v, "CPT_MAGIC_DEF")
		--local damage = caster_attack * 2.3 - defender_def		-- 伤害公式
		local truedamage = Player.GetUnitDamage(battleid, casterid, v)
		
		sys.log("猴子 化身三千的物理伤害   "..truedamage)
		
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		
		sys.log("猴子 化身三千最终的物理伤害   "..damage)
		
		if damage <= 0 then
			damage = 0
		end
		
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		local atk_damage = damage*1
		
		local per = 0.15
		
		local defender_def = defender_atk * per
		
		local defender_mag = defender_mtk * per
		
		Battle.Attack(battleid, casterid, v, atk_damage, crit)
		Battle.AddBuff(battleid, casterid, v, 112,defender_def )
		Battle.AddBuff(battleid, casterid, v, 119,defender_mag )
		if crit == 1 then 
			Battle.AddSkillBuff(battleid,casterid,casterid,138,50)
		end
		
		Battle.TargetOver(battleid)
		
		sys.log("猴子 对id为"..v.."的目标造成"..damage.."点伤害")
	end
	
	return 1
end

sys.log("猴子 SK_295_Action 结束")