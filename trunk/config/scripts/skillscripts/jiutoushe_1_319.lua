sys.log("九头蛇 SK_319_Action 开始")

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
-- 酸雾喷吐。使敌方全体目标受到30%法术强度的伤害，受到该伤害的目标无法回复生命值（所有回复生命值类型的属性和技能都不会生效。回复量为0），持续3回合。
-- 增加速度视作buff

function SK_319_Action(battleid, casterid)
	local skillid = 319		-- 技能id
	local skillAttack = 10	-- 技能攻击
	local attackNum = 0		-- 攻击个数
	
	local t = Player.GetTargets(battleid, casterid, attackNum)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	local caster_attack = Player.GetUnitProperty(battleid, casterid, "CPT_ATK")	-- 获取到攻击者的属性
	
	
	for i,v in ipairs(t)	do
		Battle.TargetOn(battleid)
		
		local truedamage = Player.GetMagicDamage(battleid, casterid, v)
		sys.log("九头蛇 群蛇乱舞对目标   "..t.. " 造成 法术伤害  "..truedamage )
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		sys.log("九头蛇群蛇乱舞对目标   "..t.. " 造成 最终法术伤害  "..damage )
		if damage <= 0 then
			damage = 0
		end
		
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		local atk_damage = damage*0.3
		
		Battle.Attack(battleid, casterid, v, atk_damage, crit)
		
		Battle.AddBuff(battleid,casterid,v,113,0)  --受到该伤害的目标无法回复生命值
		
		Battle.TargetOver(battleid)
		
		
	end
	
	return 1
end
sys.log("九头蛇 SK_319_Action 结束")