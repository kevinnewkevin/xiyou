sys.log(" 主角 主动技能 7 开始")

--主角7技能  删除对手的一张手牌

function SK_106_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 106-- 技能id

	local level = 1
	
	--Battle.Addbuff(battleid,casterid,0,0,0)

	Battle.TargetOver(battleid)

end

sys.log(" 主角 主动技能 7 结束")