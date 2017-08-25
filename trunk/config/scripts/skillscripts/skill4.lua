sys.log(" skill 4 start")

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
-- 姜子牙一号技能 对敌方单目标造成等同物理攻击的伤害，并降低对手等同伤害值的法术强度，为自己增加等量的法术强度。

-- 法术强度视作buff  Battle.buff

function SK_103_Action(battleid, casterid)

	Battle.TargetOn(battleid) --清空数据
	local skillid = 103		-- 技能id
	
	local  t = Player.GetTarget(battleid,casterid)  --获取目标 
	
	local damage=Player.GetUnitDamage(battleid,casterid,t)   --获取物理伤害

	
	--判断伤害
	if damage <= 0 then 
	
		damage = 1
	
	end
	
	local crit = Battle.GetCrit(skillid)   --是否暴击
	
	Battle.AddBuff(battleid,casterid,t,3,-damage)   --降低对手同等伤害的法术强度
	
	Battle.AddBuff(battleid,casterid,casterid,3,damage)   --增加自己等量的法术强度
	
	Battle.Attack(battleid,casterid,t,damage,crit)   --调用服务器   （伤害）(战斗者，释放者，承受者，伤害，暴击）
	
	Battle.TargetOver(battleid)  --赋给下个目标
	
	
	sys.log("skill4 对id为"..t.."的目标造成"..damage.."点伤害")
	
	
	return  true
	 
end

sys.log( "skill 4 end")