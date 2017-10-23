sys.log("龙王 SK_311_Action 开始")

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
-- 布雨。对敌方全体造成50%法术强度的伤害，并减少20%的速度。
-- 增加速度视作buff

function SK_311_Action(battleid, casterid)

	
	local skillid = 311		-- 技能id
	local skillAttack = 10	-- 技能攻击
	local attackNum = 0		-- 攻击个数
	
	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	local sudu = Player.GetUnitProperty(battleid, casterid, "CPT_AGILE")--获取速度
	
	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid)
		local  truedamage  = Player.GetMagicDamage(battleid,casterid,v)
		sys.log("龙王 布雨对目标  "..v..  "造成法术伤害"..truedamage)
		
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		sys.log("龙王 布雨对目标  "..v .. "造成最终法术伤害"..truedamage)
		
		--判断伤害
		if damage <= 0 then 
		
			damage = 0
		
		end
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		--50%的伤害
		local mag_damage = damage*0.5
		sys.log("龙王 布雨对目标  "..v .. "造成最终法术伤害的50%"..mag_damage)
		Battle.Attack(battleid,casterid,v,mag_damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		--减少20的速度
		local sudu_del = sudu*0.2
		Battle.AddBuff(battleid,casterid,v,117,sudu_del)    -- 减少20%的速度
		
		Battle.TargetOver(battleid)
		
	end
	return  true
	
end

sys.log("龙王 SK_311_Action 结束")