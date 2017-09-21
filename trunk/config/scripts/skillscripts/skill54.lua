sys.log(" skill 54 start")

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
-- 主角必杀技3技能 删除己身上所有增益buff,成功概率5*技能等级


function SK_153_Action(battleid, casterid)
	
	local skillid = 153		-- 技能id

	local  attackNum = 0   --攻击个数
	
	local level = 1

	local  t = Player.GetFriends(battleid,casterid,attackNum)  --获取目标
	
	for i,v in ipairs(t) do
		Battle.TargetOn(battleid)
		
		local per = rand.Intn(100)
		
		local skill = 5 * level
		
		if per < skill then 
			
			Player.PopAllBuffByDebuff(battleid,v)
		
		end
		
		Battle.TargetOver(battleid)
	
		sys.log("skill154")
	end
	
	
	return  true
	 
end

sys.log( "skill 54 end")