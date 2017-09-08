sys.log(" skill 37 start")

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
-- 九头蛇3技能 对敌方随机目标进行9次攻击，每次攻击造成30%法术强度伤害，每次都到伤害降低10%速度，持续1回合。。

-- 法术 物理强度视作buff  Battle.buff

function SK_136_Action(battleid, casterid)

	Battle.TargetOn(battleid)
	
	local skillid = 136	-- 技能id
	
	local  t = Player.GetTarget(battleid,casterid)  --获取目标 
	
	sys.log("SK_136_Action"..t)

	--local  caster_attack = Player.GetUnitMtk(battleid,casterid)  --获取攻击者属性
		
	--local defender_def = Player.GetCalcMagicDef(battleid,t)   -- 防御
	
	local sudu = Player.GetUnitProperty(battleid, casterid, "CPT_AGILE")
	
	local magic_damage=Player.GetMagicDamage(battleid,casterid,t)
	
	local  true_damage = magic_damage*9
	
	sys.log("SK_136_Action 的伤害"..true_damage)
	
	
	local damage = ClacDamageByAllBuff(battleid,casterid,t,true_damage)
	
		
	--判断伤害
	if damage <= 0 then 
		
		damage = 1
		
	end
	local crit = Battle.GetCrit(skillid)   --是否暴击
	
	local mag_damage = damage*0.3
	
	lcoal sudu_del = sudu*0.1
		
	Battle.Attack(battleid,casterid,t,mag_damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
	Battle.AddBuff(battleid,casterid,t,117,sudu_del)  --每次都到伤害降低10%速度
		
	Battle.TargetOver(battleid)
		
	sys.log("skil36对id为"..t.."的目标造成"..damage.."点伤害")
	
	return  true
	 
end

sys.log( "skill 37 end")