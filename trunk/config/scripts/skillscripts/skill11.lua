sys.log(" skill 11 start")

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
-- 托塔李天王2号技能 对敌方前排造成物理强度的伤害，如果敌方目标有负面效果，则必定暴击。

-- 物理强度视作buff Battle.buff

function SK_110_Action(battleid, casterid)

	Battle.TargetOn(battleid) -- 清空数据
	local skillid = 110		-- 技能id
	

	local  attackNum = 3   --攻击个数

	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	local  caster_attack = Player.GetUnitAtk(battleid,casterid)  --获取攻击者属性  物理

	
	for i,v in ipairs(t) do
	
		local defender_def = Player.GetCalcDef(battleid, v)  -- 防御
	
	    local  damage  = caster_attack-defender_def        --伤害 公式（）
	
		--判断伤害
		if damage <= 0 then 
		
			damage = 1
		
		end
		
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		local HasDebuff = Battle.HasDebuff(battleid,v)   --负面效果   
		
		--如果敌方目标有负面效果，则必定暴击。
		
		if HasDebuff == true  then
			
			--Player.ChangeSpecial(battleid,casterid,1,"BF_CRIT")  -- casterid 单元id    v buff实例id
			
			crit = 0
			
		end
		
		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
		Battle.TargetOver(battleid)  -- 赋给下个目标
		
		sys.log("skill11 对id为"..v.."的目标造成"..damage.."点伤害")
	end
	
	return  true
	 
end

sys.log( "skill 11 end")