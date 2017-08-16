sys.log(" skill 16 start")

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
-- 观音一号技能 对一个敌方目标造成法术强度的伤害，并降低目标40%的物理强度和法术强度，将其转移给一个友方目标。

-- 法术强度视作buff  Battle.buff

function SK_115_Action(battleid, casterid)
	local skillid = 115		-- 技能id
	
	local  t = Player.GetTarget(battleid,casterid)  --获取目标 
	
	
	
	local  caster_attack = Player.GetUnitProperty(battleid,casterid,"CPT_ATK")  --获取攻击者属性
	
	
	local defender_def = Player.GetUnitProperty(battleid, casterid, "CPT_DEF")  --获取防御属性
	
	--local  add_buff = Battle.AddBuff(1)     --法术强度
	
	--local  del_buff = Battle.AddBuff(1)     --降低目标法术强度
	
	--local  p_property = Battle.AddBuff(2)     --降低目标40%物理强度
	
	--local  demage  = add_buff-defender_def    --伤害 公式（）
	
	local  damage  = 4 --伤害 公式（攻击属性） --测试
	
	
	--判断伤害
	if damage <= 0 then 
	
		damage = 1
	
	end
	
	local crit = Battle.GetCrit(skillid)   --是否暴击
	
	Battle.Attack(battleid,casterid,t,damage,crit)   --调用服务器   （伤害）(战斗者，释放者，承受者，伤害，暴击）
	
	sys.log("skill16 对id为"..t.."的目标造成"..damage.."点伤害")
	
	return  true
	 
end

sys.log( "skill 16 end")