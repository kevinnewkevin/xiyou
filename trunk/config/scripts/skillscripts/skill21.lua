sys.log(" skill 21 start")

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
-- 黄牙老象3号技能 造成的伤害增加20%，可叠加。

-- 法术强度视作buff  Battle.buff

function SK_120_Action(battleid, casterid)
	local skillid = 120		-- 技能id
	
	local  t = Player.GetTarget(battleid,casterid)  --获取目标 
	
	 --local  add_buff = Battle.AddBuff(battleid,del_buff*0.2)    --造成的伤害增加20%，可叠加。
		
	sys.log("skil20")
	
	
	return  true
	 
end

sys.log( "skill 21 end")