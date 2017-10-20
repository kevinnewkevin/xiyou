sys.log("子牙 SK_298_Action 开始")

-- 技能释放 传入战斗ID和释放者的ID
-- 通过释放者和battleid取得对应的目标 单体或者多个
-- 循环/直接使用接口操控战斗 类似 战斗.攻击(战斗id, 释放者id, 承受者ID, 伤害数值, 是否暴击)
-- 
-- 
-- 所需接口
--  取得对应属性
--  计算伤害数值
--  计算是否暴击
--  攻击
-- 封神：给主角增加一个可以吸收8%最大血量的护盾，这个buff不会消失，可以叠加。每次释放增加8%，直到被驱散或者姜子牙死亡，buff消失
-- 增加速度视作buff

function SK_298_Action(battleid, casterid)

	Battle.TargetOn(battleid)
	local skillid = 298		-- 技能id
	local skillAttack = 10	-- 技能攻击
	local attackNum = 0		-- 攻击个数
	
	local t = Player.GetMainFriend(battleid, casterid)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	local hp = Player.GetUnitProperty(battleid, casterid, "CPT_HP")	-- 获取到攻击者的属性
	
	local damage = hp * 0.08
	
	Battle.Cure(battleid, t, 0, 0)
	
	local buffid = 103
	
	Battle.AddBuff(battleid,casterid,t,buffid,damage)  --给友方分别提供一个盾牌
	
	Battle.TargetOver(battleid)
	
	sys.log("姜子牙2号能")
	
	return  true
end
sys.log("子牙 SK_298_Action 结束")
