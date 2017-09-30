sys.log("SK_310_Action")

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
-- 落雷。对敌方竖排造成100%法术强度的伤害。
-- 增加速度视作buff

function SK_310_Action(battleid, casterid)

	

	local skillid = 310		-- 技能id
	
	local  p = Player.GetTarget(battleid,casterid)  --获取目标

	local  t = Player.LineTraget(battleid,p)  --获取目标
	
	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid)
	
		local  truedamage  =Player.GetMagicDamage(battleid,casterid,v)  --伤害 公式（物理伤害 减 防御 ）
		
		sys.log("雷震子  落雷对目标   "..v.. " 造成 法术伤害  "..truedamage )
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		sys.log("雷震子  落雷对目标   "..v.. " 造成 最终法术伤害  "..damage )
		--判断伤害
		if damage <= 0 then 
		
			damage = 1
		
		end
		local crit = Battle.GetCrit(skillid)   --是否暴击
		local Bhp = Player.GetUnitProperty(battleid, casterid, "CPT_HP")
		local  hp_pro= Bhp*0.1
		damege = damage + hp_pro
		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		Battle.TargetOver(battleid)
	
		
	end
	
	return  true
	 
end

