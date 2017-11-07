
-- UI打开时调用
function WhenUIOpen(uiName, uiWindow)
	--1.进入主界面事件(WhenUIOpen "zhujiemian")指引点击解魂模式(uiWindow.contentPane:GetChild("n43"):GetChildAt(1))
	if uiName == "zhujiemian" then
		if GuideSystem.IsNotFinish(1) then
			GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n43"):GetChildAt(1), 300, 300,"欢迎来到小西游，请进入你的解魂世界！",1300,0);
			GuideSystem.SetFinish(1);
		end
	end
	--2.	解魂界面打开事件(WhenUIOpen "daguanka")
	--指引点击解魂按钮(uiWindow.contentPane:GetChild("n4"):GetChild("n3"))
	if uiName == "daguanka" then
		if GuideSystem.IsNotFinish(2) then
			GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n4"):GetChild("n3"),"通过解魂可以获得卡牌！",1138,409);
			GuideSystem.SetFinish(2);
		
		elseif not GuideSystem.IsNotFinish(2) and GuideSystem.IsNotFinish(10) then
			GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n6"),"去看看下一个英雄！",1138,409);
			GuideSystem.SetFinish(10);
		end
	end
	--3.	小关卡界面打开事件(WhenUIOpen "xiaoguanka")
	--指引点击第一小关(uiWindow.contentPane:GetChild("n22"))
	if uiName == "xiaoguanka" then
		if GuideSystem.IsNotFinish(3) then
			GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n22"),"从第一关开始吧！",689,214);
			GuideSystem.SetFinish(3);

		elseif not GuideSystem.IsNotFinish(3) and  GuideSystem.IsNotFinish(12) then
			GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n30"),150,150,"可以上阵新的卡牌了",284,403);
			GuideSystem.SetFinish(12);
		end
	end
	
	if uiName == "paiku" then
		if GuideSystem.IsNotFinish(13) then
			GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n6").asCom:GetChild("n27").asList:GetChildAt(0),"上阵新的卡牌",430,0,756,72);
			GuideSystem.SetFinish(13);
		end	
	end
	if uiName == "xiangxiziliao" then
		if GuideSystem.IsNotFinish(14) then
			GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n61"),"上阵！",0,0);
			GuideSystem.SetFinish(14);
		end	
	end
end

-- UI关闭时调用
function WhenUIClose(uiName)
	if uiName == "xiangxiziliao" then
		if not GuideSystem.IsNotFinish(999) then
			GuideSystem.ClearGuide();
		end
	end
end

-- 特殊事件调用
function SpecialEvent(type, param1)
	--4.	挑战按钮打开事件(SpecialEvent "xiaoguanka_challenge")
	--指引点击挑战按钮(uiWindow.contentPane:GetChild("n26"):GetChild("n2"))
	if type == "xiaoguanka_challenge" then
		if GuideSystem.IsNotFinish(4) then
			GuideSystem.StartGuide(UIManager.GetWindow("xiaoguanka").contentPane:GetChild("n26"):GetChild("n2"),"让我看看你的实力，挑战！",631,460);
			GuideSystem.SetFinish(4);
		end
	end
	
	--5.	进入战斗事件（需要回合数）(SpecialEvent “battle_start”) Battle._Turn
	--第一回合指引点击第一个主角技能按钮(uiWindow.contentPane:GetChild(“n50”):GetChildAt(0))
	if type == "battle_start" then
		if Battle._Turn == 1 then
			if GuideSystem.IsNotFinish(5) then
				GuideSystem.StartGuide(UIManager.GetWindow("BattlePanel").contentPane:GetChild("n37"):GetChild("n1"),"每回合你会获得一个魂！","dianji1",1375,525);
				--GuideSystem.SetFinish(5);
			end
		end
	end
	if type == "dianji1" then
		if Battle._Turn == 1 then
			if GuideSystem.IsNotFinish(5) then
				GuideSystem.StartGuide(UIManager.GetWindow("BattlePanel").contentPane:GetChild("n50"):GetChildAt(0),"消耗魂可以释放主角技能！",322,329);
				GuideSystem.SetFinish(5);
			end
		end
	end
	
	
	--6.	点击主角按钮事件(SpecialEvent “battle_roleskill”) 
	--指引点击结束回合按钮(uiWindow.contentPane:GetChild(“n16”))
	if type == "battle_roleskill" then
		if Battle._Turn == 1 then
			if GuideSystem.IsNotFinish(6) then
				GuideSystem.StartGuide(UIManager.GetWindow("BattlePanel").contentPane:GetChild("n16"),"操作过后，请结束回合！",1107,598);
				GuideSystem.SetFinish(6);
			end
		end
	end
	--7.	第二回合开始(SpecialEvent “battle_start”) Battle._Turn
	--指引点击第一章卡牌(uiWindow.contentPane:GetChild(“n17”))
	if type == "battle_start" then
		if Battle._Turn == 2 then
			if GuideSystem.IsNotFinish(7) then
				GuideSystem.StartGuide(UIManager.GetWindow("BattlePanel").contentPane:GetChild("n38"):GetChild("n1"),"第二回合，又增加了一个魂！","dianji2",1375,525);
				--GuideSystem.SetFinish(5);
			end
		end
	end
	if type == "ondragend" then
		if Battle._Turn == 2 then
			if GuideSystem.IsNotFinish(7) then
				GuideSystem.StartGuide(UIManager.GetWindow("BattlePanel").contentPane:GetChild("n17"),"消耗2个魂可以上阵这张卡牌！每张卡牌消耗的魂显示在卡牌右下角！",492,552);
			end
		end
	end
	if type == "dianji2" then
		if Battle._Turn == 2 then
			if GuideSystem.IsNotFinish(7) then
				GuideSystem.StartGuide(UIManager.GetWindow("BattlePanel").contentPane:GetChild("n17"),"消耗2个魂可以上阵这张卡牌！每张卡牌消耗的魂显示在卡牌右下角！",492,552);
			end
		end
	end
	--8.	点击卡牌事件(SpecialEvent “battle_selectcard”)
	--指引点击前排中间位置(StartGuideInScene  Battle.GetPoint(1)   width height)
	if type == "battle_selectcard" then
		if Battle._Turn == 2 then
			if GuideSystem.IsNotFinish(8) then
				GuideSystem.StartGuideInScene(Battle.GetPoint(1).gameObject,300,300,"选择卡牌上场的位置！",984,247);
			end
		end
	end
	--9.	卡牌上阵事件(SpecialEvent “battle_cardonbattle”)
	--指引点击结束回合按钮
	if type == "battle_cardonbattle" then
		if Battle._Turn == 2 then
			if GuideSystem.IsNotFinish(9) then
				GuideSystem.StartGuide(UIManager.GetWindow("BattlePanel").contentPane:GetChild("n16"),"操作过后，请结束回合！",1107,598);
				GuideSystem.SetFinish(7);
				GuideSystem.SetFinish(8);
				GuideSystem.SetFinish(9);
			end
		end
	end
	--向右翻页
	if type == "daguanka_rightclick" then
		if GuideSystem.IsNotFinish(11) then
			GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n4"):GetChild("n3"),"通过解魂可以获得卡牌！",1138,409);
			GuideSystem.SetFinish(11);
		end
	end
end