sys.log("skill 15 start")

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
-- 女娲3号技能 使一个友方单位免疫一次会导致其死亡的伤害。

-- 物理强度视作buff Battle.buff

function SK_114_Action(battleid, casterid)
	local skillid = 114		-- 技能id
	
	local t = Player.GetTarget(battleid, casterid)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	
	--local  spell = Battle.AddBuff(1)   --（暂时么有这个函数） 造物。使一个友方单位免疫一次会导致其死亡的伤害。
	
	
	sys.log("skill15 对id为"..t.."的目标增加免疫伤害")
	
	
	return true
end

sys.log("skill 15 end")