sys.log(" 杨戬 SK_350_Action 开始")

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
-- 放狗：对敌方后排造成100%物理伤害
-- 增加速度视作buff

function SK_350_Action(battleid, casterid)

	local skillid = 350		-- 技能id
	
	local  t = Player.BackTarget(battleid,casterid)  --获取目标 

	for i,v in ipairs(t)do
		Battle.TargetOn(battleid)
		local  truedamage  = Player.GetUnitDamage(battleid,casterid,v)    --伤害 公式（）
	
		sys.log("杨戬 放狗对目标造成的物理伤害   ".. truedamage)

		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)

		sys.log("杨戬 放狗对目标造成的最终物理伤害   ".. damage)
		
		--判断伤害
		if damage <= 0 then 
		
			damage = 0
		
		end
		
		local crit = Battle.GetCrit(skillid)   --是否暴击

		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器   （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		Battle.TargetOver(battleid)

	end
	
	
	return  true
	 
end
sys.log(" 杨戬 SK_350_Action 结束")