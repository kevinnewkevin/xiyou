sys.log(" skill 26 start")

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
-- 龙王2号技能 对敌方全体造成50%法术强度的伤害，并减少20%的速度。

-- 法术强度视作buff  Battle.buff

function SK_125_Action(battleid, casterid)

	local skillid = 125	-- 技能id

	local  attackNum = 0  --攻击个数

	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	local sudu = Player.GetUnitProperty(battleid, casterid, "CPT_AGILE")
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
		
		local  truedamage  = Player.GetMagicDamage(battleid,casterid,v)
		
		sys.log("SK_125_Action 的伤害"..truedamage)
		
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		
		--判断伤害
		if damage <= 0 then 
		
			damage = 1
		
		end
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		Battle.Attack(battleid,casterid,v,damage*0.5,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		Battle.AddBuff(battleid,casterid,v,117,sudu*0.2)    -- 减少20%的速度
		Battle.TargetOver(battleid)
	
		sys.log("skill26 对id为"..v.."的目标减少"..damage.."点伤害")
	end
	
	return  true
	 
	 
end

sys.log( "skill 26 end")