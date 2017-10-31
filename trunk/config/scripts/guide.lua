
-- UI打开时调用
function WhenUIOpen(uiName, uiWindow)
	--1.进入主界面事件(WhenUIOpen "zhujiemian")指引点击解魂模式(uiWindow.contentPane:GetChild("n43"):GetChildAt(1))
	if uiName == "zhujiemian" then
		if GuideSystem.IsNotFinish(1) then
			GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n43"):GetChildAt(1));
			GuideSystem.SetFinish(1);
		end
	end
	--2.	解魂界面打开事件(WhenUIOpen "daguanka")
	--指引点击解魂按钮(uiWindow.contentPane:GetChild("n4"):GetChild("n3"))
	if uiName == "daguanka" then
		if GuideSystem.IsNotFinish(2) then
			GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n4"):GetChild("n3"));
			GuideSystem.SetFinish(2);
		end
	end
	--3.	小关卡界面打开事件(WhenUIOpen "xiaoguanka")
	--指引点击第一小关(uiWindow.contentPane:GetChild("n22"))
	if uiName == "xiaoguanka" then
		if GuideSystem.IsNotFinish(3) then
			GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n22"));
			GuideSystem.SetFinish(3);
		end
	end
	--4.	挑战按钮打开事件(SpecialEvent "xiaoguanka_challenge")
	--指引点击挑战按钮(uiWindow.contentPane:GetChild("n26"):GetChild("n2"))
	if uiName == "xiaoguanka_challenge" then
		if GuideSystem.IsNotFinish(4) then
			GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n26"):GetChild("n2"));
			GuideSystem.SetFinish(4);
		end
	end
	
end

-- 特殊事件调用
function SpecialEvent(type, param1)
	--5.	进入战斗事件（需要回合数）(SpecialEvent “battle_start”) Battle._Turn
	--第一回合指引点击第一个主角技能按钮(uiWindow.contentPane:GetChild(“n50”):GetChildAt(0))
	if type == "battle_start" then
		if Battle._Turn == 1 then
			if GuideSystem.IsNotFinish(5) then
				GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n50"):GetChildAt(0));
				GuideSystem.SetFinish(5);
			end
		end
	end
	--6.	点击主角按钮事件(SpecialEvent “battle_roleskill”) 
	--指引点击结束回合按钮(uiWindow.contentPane:GetChild(“n16”))
	if type == "battle_roleskill" then
		if Battle._Turn == 1 then
			if GuideSystem.IsNotFinish(6) then
				GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n16"));
				GuideSystem.SetFinish(6);
			end
		end
	end
	--7.	第二回合开始(SpecialEvent “battle_start”) Battle._Turn
	--指引点击第一章卡牌(uiWindow.contentPane:GetChild(“n17”))
	if type == "battle_start" then
		if Battle._Turn == 2 then
			if GuideSystem.IsNotFinish(7) then
				GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n17"));
				GuideSystem.SetFinish(7);
			end
		end
	end
	--8.	点击卡牌事件(SpecialEvent “battle_selectcard”)
	--指引点击前排中间位置(StartGuideInScene  Battle.GetPoint(1)   width height)
	if type == "battle_selectcard" then
		if Battle._Turn == 2 then
			if GuideSystem.IsNotFinish(8) then
				GuideSystem.StartGuideInScene(Battle.GetPoint(1).gameObject,300,300);
				GuideSystem.SetFinish(8);
			end
		end
	end
	--9.	卡牌上阵事件(SpecialEvent “battle_cardonbattle”)
	--指引点击结束回合按钮
	if type == "battle_roleskill" then
		if Battle._Turn == 2 then
			if GuideSystem.IsNotFinish(9) then
				GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n16"));
				GuideSystem.SetFinish(9);
			end
		end
	end
end