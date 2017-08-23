sys.log(" skill 7 start")

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
-- 哪吒1号技能 自己在场时为所有友军增加物理强度，提高的数值等于哪吒物理强度的20%，可叠加。

-- 物理强度视作buff Battle.buff

function SK_106_Action(battleid, casterid)
	local skillid = 106		-- 技能id

	local  attackNum = 0   --攻击个数
	

	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标
	
	--local  p_buff = Battle.AddBuff(1)  --自己物理强度
	

	for i,v in ipairs(t) do
		--local  add_buff = Battle.AddBuff(p_buff*0.2)  --给所有友方加物理强度    （暂时么有这个函数）
		
		sys.log("skill7")
	end
	
	return  true
	 
end

sys.log( "skill 7 end")