
sys.log(" skill 8 start")

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
-- 哪吒2号技能 掷出乾坤圈，击中敌人后选择下一个敌人弹射，弹射3次，对每个敌人造成50%物理强度的伤害。每击中一个敌人增加自己10%的物理强度，持续1回合

-- 物理强度视作buff Battle.buff

function SK_107_Action(battleid, casterid)

	

	local skillid = 107	-- 技能id

	local  attackNum = 3   --攻击个数

	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	local  caster_attack = Player.GetUnitAtk(battleid,casterid)  --获取攻击者属性  物理
	
	
	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid) --清空数据
		
		local  truedamage = Player.GetUnitDamage(battleid,casterid,v)  --获取物理伤害
		
		sys.log("SK_107_Action 的伤害"..truedamage)
		
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		
		--判断伤害
		if damage <= 0 then 
		
			damage = 1
		
		end
	
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		local atk_damage = damage*0.5
		
		local strong_damage = caster_attack*0.1
		
		Battle.Attack(battleid,casterid,v,atk_damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		Battle.AddBuff(battleid,casterid,casterid,110,strong_damage) --每击中一个敌人增加自己10%的物理强度
		
		Battle.TargetOver(battleid)  --赋给下个目标
		
		sys.log("skill8 对id为"..v.."的目标造成"..damage.."点伤害")  
		
	end
	
	sys.log("skill8"..caster_attack)
	return  true
	 
end

sys.log( "skill 8 end")