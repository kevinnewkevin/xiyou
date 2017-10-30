
-- UI打开时调用
function WhenUIOpen(uiName, uiWindow)
	--[[if uiName == "denglu" then
		if GuideSystem.IsNotFinish(1) then
			GuideSystem.StartGuide(uiWindow.contentPane:GetChild("n10"):GetChild("n11"));
			GuideSystem.SetFinish(1);
		end
	end--]]
end

-- 特殊事件调用
function SpecialEvent(type, param1)
	--[[if type == "battlestart" then
		if Battle._Turn == 1 then
			if GuideSystem.IsNotFinish(2) then
				GuideSystem.StartGuide();
				GuideSystem.SetFinish(2);
			end
		end
	end--]]
end