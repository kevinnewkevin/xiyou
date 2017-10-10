sys.log(" 老象 SK_308_Action 开始")

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
-- 神力：强制一个目标普通攻击自己，持续2回合，优先后排对面，然后是两侧，然后是前排
-- 增加速度视作buff

function SK_308_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	local skillid = 308		-- 技能id
	local skillAttack = 10	-- 技能攻击
	--local attackNum = 0		-- 攻击个数
	
	local t = Player.GetTarget(battleid, casterid)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	Battle.Attack(battleid, casterid, t, 0, 0)
	sys.log("老象 神力 给目标  "..t.." 加124buff")
	Battle.AddBuff(battleid,casterid,t,142,casterid)  
	sys.log("老象 神力 给目标  "..t.. " 加buff完成")	
	Battle.TargetOver(battleid)
	
	return 1
end
sys.log(" 老象 SK_308_Action 结束")
