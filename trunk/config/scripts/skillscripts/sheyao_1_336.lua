sys.log("蛇妖 SK_336_Action 开始")

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
-- 攻击敌方全体，造成50%法术伤害，并且使目标中毒，每回合拾取蛇妖当前生命值的20%，持续2回合
-- 增加速度视作buff

function SK_336_Action(battleid, casterid)

	local skillid = 336		-- 技能id

	local attackNum = 0
	
	local  p = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	for i,v in ipairs(p) do
		Battle.TargetOn(battleid)
		local  truedamage  =Player.GetMagicDamage(battleid,casterid,v)  --伤害 公式（物理伤害 减 防御 ）
		
		sys.log("蛇妖 1号技能 对目标   "..v.. " 造成 法术伤害  "..truedamage )
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		sys.log("蛇妖  1号技能 对目标   "..v.. " 造成 最终法术伤害  "..damage )
		--判断伤害
		if damage <= 0 then 
		
			damage = 0
		
		end
		damege = damage * 0.5
		local crit = Battle.GetCrit(skillid)   --是否暴击
		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		local hp = Player.GetUnitProperty(battleid, casterid, "CPT_CHP")
		local hp_damage = hp * 0.2
		Battle.AddBuff(battleid,casterid,v,150,0)
		Battle.TargetOver(battleid)
	
	end
	return  true
end

sys.log("蛇妖 SK_336_Action 结束")