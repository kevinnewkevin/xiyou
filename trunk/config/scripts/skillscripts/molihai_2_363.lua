sys.log("魔礼海 SK_362_Action 开始")

-- 技能释放 传入战斗ID和释放者的ID
-- 通过释放者和battleid取得对应的目标 单体或者多个
-- 循环/直接使用接口操控战斗 类似 战斗.攻击(战斗id, 释放者id, 承受者ID, 伤害数值, 是否暴击)
-- 
-- 
-- 所需接口
--	取得目标 （GetTarget（）  单   GetTargets（）  复）
--  取得对应属性 GetUnitProperty()
--  计算伤害数值 demage
--  计算是否暴击
--  攻击
-- 琵琶毒：对地方全体造成50%法术伤害，并且中毒，每回合拾取当前生命值10%生命，持续2回合

-- 法术强度视作buff  Battle.buff

function SK_363_Action(battleid, casterid)
	
	local skillid = 363

	local attackNum = 0

	local t = Player.GetTargets(battleid,casterid,attackNum)

	for i,v in ipairs(t) do 

		Battle.TargetOn(battleid)

		local  truedamage  = Player.GetMagicDamage(battleid,casterid,v)    --伤害 公式（）
	
		sys.log("魔礼海 琵琶毒对目标造成的法术伤害   ".. truedamage)

		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)

		sys.log("魔礼海 琵琶毒对目标造成的最终法术伤害   ".. damage)
		--判断伤害
		if damage <= 0 then 

			damage = 0

		end
		local crit = Battle.GetCrit(skillid)   --是否暴击

		damage = damage * 0.5
		
		Battle.Attack(battleid,casterid,v,damage,crit)

		--掉当前血量的10%
		local hp = Player.GetUnitProperty(battleid, v, "CPT_CHP")

		local hp_damage = hp * 0.1

		Battle.AddBuff(battleid,casterid,v,179,hp_damage)

		Battle.TargetOver(battleid)
	end

	return ture 
	 
end
sys.log("魔礼海 SK_363_Action 结束")
