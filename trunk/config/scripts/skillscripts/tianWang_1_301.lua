sys.log("天王 SK_301_Action 开始")

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
-- 天罡刀。对敌方前排造成物理强度的伤害。30%概率可以使击中的目标无法行动1回合
-- 增加速度视作buff
set_random_seed()
function SK_301_Action(battleid, casterid)

	local skillid = 301		-- 技能id
	local skillAttack = 10	-- 技能攻击
	--local attackNum = 0		-- 攻击个数
	
	local  t = Player.FrontTarget(battleid,casterid)  --获取目标
	
	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid) -- 清空数据
	
	    local  truedamage  = Player.GetUnitDamage(battleid,casterid,v)       --伤害 公式（）
		
		sys.log("天王对目标造成的物理伤害    "..truedamage)
		
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		sys.log("天王对目标造成的最终物理伤害    "..truedamage)
		--判断伤害
		if damage <= 0 then 
		
			damage = 0
		
		end
		
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		if percent() <= 30 then 						
			Battle.AddBuff(battleid,casterid,v,165,0)  --30%概率可以使击中的目标无法行动1回合
		end
	
		
		Battle.TargetOver(battleid)  -- 赋给下个目标
		
		
	end
	
	return  true
	
	
end

sys.log("天王 SK_301_Action 结束")