sys.log(" skill 24 start")

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
-- 雷震子3号技能 使下一次雷击必定溅射，下一次落雷必定连击。

-- 法术强度视作buff  Battle.buff

function SK_123_Action(battleid, casterid)

	local  p = Player.GetTarget(battleid,casterid)  --获取目标 

	--local  add_buff = Battle.AddBuff(battleid)    -- 使下一次雷击必定溅射
	
	--local  add_buff = Battle.AddBuff(battleid)    -- 下一次落雷必定连击
	
	sys.log("skill20")
	
	return  true

end



sys.log(" skill 24 end")