sys.log(" skill 20 start")

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
-- 黄牙老象2号技能 下3次受到的伤害降低40%。

-- 法术强度视作buff  Battle.buff

function SK_119_Action(battleid, casterid)
	local skillid = 119		-- 技能id
	
	local  t = Player.GetTarget(battleid,casterid)  --获取目标 

	--local  add_buff = Battle.AddBuff(battleid,del_buff*0.4)    --下3次受到的伤害降低40%。
	
	sys.log("skill20")
	
	return  true
	 
end

sys.log( "skill 20 end")