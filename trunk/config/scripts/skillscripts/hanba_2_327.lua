sys.log(" 旱魃 SK_327_Action 开始")

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
-- 焚天。攻击目标全体
-- 增加速度视作buff

function SK_327_Action(battleid, casterid)

	local skillid = 327		-- 技能id

	local  attackNum = 0   --攻击个数
	
	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标 
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
		local  truedamage  = Player.GetMagicDamage(battleid,casterid,v)    --伤害 公式（）
		sys.log("旱魃 焚天对目标造成的法术伤害   ".. truedamage)
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		sys.log("旱魃 焚天对目标造成的最终法术伤害   ".. damage)
		--判断伤害
		if damage <= 0 then 
			damage = 0
		end
		local crit = Battle.GetCrit(skillid)   --是否暴击
		damage = damage * 0.55
		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器   （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		--20%几率获得点燃效果
		local hp = Player.GetUnitProperty(battleid, casterid, "CPT_CHP")
		local hp_damage = hp * 0.1
		if percent() <= 20 then 
			Battle.AddBuff(battleid,casterid,v,162,hp_damage)
		end

		Battle.TargetOver(battleid)
	
	end
	
	return  true
	 
end
sys.log(" 旱魃 SK_327_Action 结束")