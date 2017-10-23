sys.log("龙后 SK_328_Action 开始")

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
-- 寒冰。对目标造成70%法术伤害，并有20%几率冰冻目标1回合。攻击目标一列。
-- 增加速度视作buff

function SK_328_Action(battleid, casterid)

	local skillid = 328		-- 技能id
	
	local  p = Player.GetTarget(battleid,casterid)  --获取目标

	local  t = Player.LineTraget(battleid,p)  --获取目标
	
	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid)
	
		local  truedamage  =Player.GetMagicDamage(battleid,casterid,v)  --伤害 公式（物理伤害 减 防御 ）
		
		sys.log("龙后  寒冰对目标   "..v.. " 造成 法术伤害  "..truedamage )
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		sys.log("龙后  寒冰对目标   "..v.. " 造成 最终法术伤害  "..damage )
		--判断伤害
		if damage <= 0 then 
			damage = 0
		end
		local crit = Battle.GetCrit(skillid)   --是否暴击
		damege = damage * 0.7
		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		--20%几率冰冻
		if percent() <= 20 then 
			Battle.AddBuff(battleid,casterid,v,149,0)
		end 
		Battle.TargetOver(battleid)
	
		
	end
	
	return  true
	 
end

sys.log("龙后 SK_328_Action 结束")