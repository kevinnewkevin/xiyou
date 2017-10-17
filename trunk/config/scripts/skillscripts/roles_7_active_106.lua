sys.log(" 主角 主动技能 7 开始")

--主角7技能  删除对手的一张手牌

function SK_190_Action(battleid, casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 190-- 技能id

	local level = 1
	
	sys.log("casterid "..casterid)
	local t = Player.GetMainTarget(battleid, casterid, false)
	sys.log("target main id "..t)
	--获取扔的卡牌
	local throwCard = Player.ThrowCard(battleid, casterid, t)
	
	sys.log("throwCard "..throwCard)
	--扔掉卡牌
	Player.Throw(battleid, t, throwCard)
	
	Battle.TargetOver(battleid)

end

sys.log(" 主角 主动技能 7 结束")