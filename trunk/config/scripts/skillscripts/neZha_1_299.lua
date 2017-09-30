sys.log("SK_299_Action")

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
-- 乾坤圈。掷出乾坤圈，击中敌人主角后选择下一个敌人弹射，弹射3次，对每个敌人造成50%物理强度的伤害。每击中一个敌人增加自己10%的物理强度，持续2回合。
-- 增加速度视作buff

function mulatk(atk)
	return atk * 0.1
end 

function SK_299_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	local skillid = 299		-- 技能id
	local skillAttack = 10	-- 技能攻击
	--local attackNum = 0		-- 攻击个数
	
	
	local t =Player.GetMainTarget(battleid, casterid)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	--local caster_attack = Player.GetUnitProperty(battleid, casterid, "CPT_ATK")	-- 获取到攻击者的属性
	
	local  damage = Player.GetUnitDamage(battleid,casterid,t)  --获取物理伤害
	
	sys.log("哪吒对主角目标"..t.."造成的物理伤害"..damage)
	
		
	damage = ClacDamageByAllBuff(battleid,casterid,t,damage)
	sys.log("哪吒对主角目标"..t.."造成的最终物理伤害"..damage)
	
	
	damage = mul(damage,0.5)
	
	--判断伤害
	if damage <= 0 then 
	
		damage = 1
	
	end
	
	local crit = Battle.GetCrit(skillid)   --是否暴击
	
	local atk =  Player.GetUnitAtk(battleid,t)
	
	Battle.Attack(battleid,casterid,t,damage,crit)
	
	atk = mul(atk,0.1)
	
	Battle.AddBuff(battleid,casterid,casterid,110,atk)
	
	Battle.TargetOver(battleid)
	
	local  aroud_target = Player.GetTargetsAround(battleid,t)
	
	sys.log("nezha  zhujue")
	
	for i,v in ipairs(aroud_target)	do
	
		Battle.TargetOn(battleid)
		
		sys.log("nezha  zhujue"..v)
		
		damage = Player.GetUnitDamage(battleid,casterid,v)  --获取物理伤害
		
		
		damage = ClacDamageByAllBuff(battleid,casterid,v,damage)
		
		
		--判断伤害
		if damage <= 0 then 
		
			damage = 1
		
		end
		
		crit = Battle.GetCrit(skillid)   --是否暴击
		
		--damage = damage*0.5
		
		Battle.Attack(battleid,casterid,v,damage,crit)
		
		atk =  Player.GetUnitAtk(battleid,t) 
		
		--atk = atk* 0.1
		
		Battle.AddBuff(battleid,casterid,casterid,110,atk)
		
		Battle.TargetOver(battleid)
		

	end
	
	return 1
end
