sys.log("如来 SK_360_Action 开始")

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
-- 千手观音：无数手掌击打对方随机目标9次
-- 增加速度视作buff

function SK_360_Action(battleid, casterid)
	
	local skillid = 360		-- 技能id
	local skillAttack = 10	-- 技能攻击
	--local attackNum = 0		-- 攻击个数

	local t = Player.GetTargetsRandom(battleid, casterid,9)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择


	for i,v in ipairs(t) do

		local IsDead  = Player.CheckUnitDead(battleid,v)

		if IsDead == true then

			v =  Player.RandomTarget(battleid, casterid)
		end
	
		if v == -1 then 
			break
		end

		Battle.TargetOn(battleid)
	
		local true_damage=Player.GetUnitDamage(battleid,casterid,v)
		
		sys.log("如来 千手观音对目标   "..v.. " 造成 物理伤害  "..true_damage )
		
		local damage = ClacDamageByAllBuff(battleid,casterid,v,true_damage)

		sys.log("如来 千手观音对目标   "..v.. " 造成 最终物理伤害  "..damage )
		
		--判断伤害
		if damage <= 0 then
			
			damage = 0
			
		end
		local crit = Battle.GetCrit(skillid)   --是否暴击

		damage = damage * 0.2

		sys.log("如来 千手观音对目标   "..v.. " 造成 最终物理伤害 20%  " ..damage )

		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		Battle.TargetOver(battleid)
	end
	
	return 1
end

sys.log("如来 SK_360_Action 结束")