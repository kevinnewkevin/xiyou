sys.log("老象 SK_307_Action 开始")

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
-- 厚皮。下3次受到的伤害降低40%。
-- 增加速度视作buff

function SK_307_Action(battleid, casterid)
	
	Battle.TargetOn(battleid)
	
	local skillid = 307		-- 技能id
	
	--local  t = Player.GetTarget(battleid,casterid)  --获取目标 
	
	Battle.Cure(battleid,casterid,0,0)
	sys.log("老象 厚皮 加buff")
	Battle.AddBuff(battleid,casterid,casterid,123,40)    --下3次受到的伤害降低40%。
	sys.log("老象 厚皮 加buff完成")
	Battle.TargetOver(battleid)
	

	
	return  true
end

sys.log("老象 SK_307_Action 结束")