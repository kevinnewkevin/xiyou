sys.log(" skill 6 start")

-- 技能释放 传入战斗ID和释放者的ID
-- 通过释放者和battleid取得对应的目标 单体或者多个
-- 循环/直接使用接口操控战斗 类似 战斗.攻击(战斗id, 释放者id, 承受者ID, 伤害数值, 是否暴击)
-- 
-- 
-- 所需接口
--	取得目标 （GetTarget（）  单   GetTargets（）  复）
--  取得对应属性  GetUnitProperty()
--  计算伤害数值  demage
--  计算是否暴击
--  攻击
-- 姜子牙3号技能 吸收场上所有护盾，对敌方单体造成法术强度+护盾吸收值*3的伤害。

-- 法术强度视作buff  Battle.buff
 
function SK_105_Action(battleid, casterid)

	
	local skillid = 105		-- 技能id

	local  attackNum = 0   --攻击个数

	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	
	local  caster_attack = Player.GetUnitMtk(battleid,casterid)  --获取攻击者的法强
	
	sys.log(1)
	
	local HDnum =  Player.GetUnitSheld(battleid)   --护盾值
	
	sys.log(2)
	
	--Battle.AddBuff(battleid,casterid,t,2,5)   --吸收场上所有护盾
	
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid) --清空数据
	
		local  defender_def = Player.GetCalcMagicDef(battleid,v)  --获取被攻击者的法强防御
		
		local  damage  = caster_attack+HDnum*3-defender_def    --伤害 公式
		
	
		--判断伤害
		if damage <= 0 then 
		
			damage = 1
		
		end
		
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器   （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		Battle.TargetOver(battleid)  --赋给下个目标
		
		sys.log("skill6 对id为"..v.."的目标造成"..damage.."点伤害")
	end
	
	return  true
	 
end

sys.log( "skill 6 end")