sys.log(" skill 27 start")

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
-- 龙王3号技能 对敌方全体造成50%法术强度的伤害，并清除所有敌方负向状态，每清除一个负向状态额外造成50%法术强度的伤害。

-- 法术强度视作buff  Battle.buff

function SK_126_Action(battleid, casterid)

	Battle.TargetOn(battleid)
	local skillid = 126	-- 技能id

	local  attackNum = 0  --攻击个数

	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	local  caster_attack = Player.GetUnitMtk(battleid,casterid)  --获取攻击者属性  fashu 
	
	for i,v in ipairs(t) do
		
		local defender_def = Player.GetCalcMagicDef(battleid, v)  --防御
		
	
		local  damage  = caster_attack*0.5-defender_def  --伤害 公式（ ）
		
		
		--判断伤害
		if damage <= 0 then 
		
			damage = 1
		
		end
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
	
		sys.log("skill26 对id为"..v.."的目标减少"..damage.."点伤害")
	end
	
	debuffnum = Player.PopAllBuffByDebuff(battleid,t)
	
		if debuffnum >0 then 
			demage = int(caster_attack / 2)
			
			for a=1,debuffnum,1 do
			
				Battle.Attack(battleid,casterid,t,demage,crit)
		
			end
		
		end
		
		Battle.TargetOver(battleid)
	
	return  true
	 
	 
end

sys.log( "skill 27 end")