sys.log(" skill 23 start")

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
-- 雷震子2号技能 对敌方全体造成50%法术强度的伤害。

-- 法术强度视作buff  Battle.buff

function SK_122_Action(battleid, casterid)

	
	local skillid = 122	-- 技能id

	local  attackNum = 0  --攻击个数

	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	--local  caster_attack = Player.GetUnitMtk(battleid,casterid)  --获取攻击者属性  fashu 
	
	for i,v in ipairs(t) do
		
		Battle.TargetOn(battleid)
		
		--local defender_def = Player.GetCalcMagicDef(battleid, v)
		
		local  magic_damage = Player.GetMagicDamage(battleid,casterid,v)  --法术伤害
	
		local  truedamage  = magic_damage  --伤害 公式（物理伤害 减 防御 ）
		
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		
		
		--判断伤害
		if damage <= 0 then 
		
			damage = 1
		
		end
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		local mag_damage = damage*0.5
		
		Battle.Attack(battleid,casterid,v,mag_damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		local  _BoolCombo = Player.GetCheckSpec(battleid,casterid,"BF_COMBO")
	
		if _BoolCombo == 1 then 
		
			local mag_damage = mag_damage*2
			
			sys.log("SK_122_Action   连击伤害"..mag_damage)
			
			Battle.Attack(battleid,casterid,v,mag_damage,crit)
			
			local  buffid = Player.GetOneSpecial(battleid,casterid,"BF_COMBO")
			
			if buffid == 0 then
				
				sys.log("SK_122_Action 111"..buffid)
				return false
				
			else
				sys.log("SK_122_Action 清除"..buffid)
				Player.PopSpec(battleid,casterid,buffid,"BF_COMBO")
			
			end
		end
		Battle.TargetOver(battleid)
	
		sys.log("skill23 对id为"..v.."的目标减少"..mag_damage.."点伤害")
	end
	
	
	return  true
	 
	 
end

sys.log( "skill 23 end")