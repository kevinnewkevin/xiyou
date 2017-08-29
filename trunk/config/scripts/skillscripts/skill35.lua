sys.log(" skill 35 start")

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
-- 鲨将军技能 使敌方一个目标受到法术强度的伤害，受到该伤害的目标无法回复生命值，持续3回合。

-- 法术 物理强度视作buff  Battle.buff

function SK_134_Action(battleid, casterid)

	Battle.TargetOn(battleid)
	local skillid = 134	-- 技能id
	
	local  t = Player.GetTarget(battleid,casterid)  --获取目标 

	local  caster_attack = Player.GetUnitMtk(battleid,casterid)  --获取攻击者属性
		
		
	local defender_def = Player.GetCalcMagicDef(battleid,t)   -- 防御
	
	local  damage = caster_attack-defender_def
		
	
	--判断伤害
	if damage <= 0 then 
		
		damage = 1
	
	end
	local crit = Battle.GetCrit(skillid)   --是否暴击
		
	Battle.Attack(battleid,casterid,t,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
	Battle.AddBuff(battleid,casterid,t,1,0)  --受到该伤害的目标无法回复生命值
		
	Battle.TargetOver(battleid)
		
	sys.log("skil35 对id为"..t.."的目标造成"..damage.."点伤害")
	
	return  true
	 
end

sys.log( "skill 35 end")