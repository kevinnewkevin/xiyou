local GameObject = UnityEngine.GameObject;

function BattlePanel_start()
	--SkipBtn
	local btnObj = GameObject.Find("BattlePanel(Clone)/SkipBtn");
	local btn = GameObject.GetComponent(btnObj, "Button");
	local evt = btn.onClick;
	evt.AddListener(evt, BattlePanel_skip);
	
	--HandCards
	btnObj = GameObject.Find("BattlePanel(Clone)/Card0");
	btn = GameObject.GetComponent(btnObj, "Button");
	evt = btn.onClick;
	evt.AddListener(evt, HandCards_1);
	
	btnObj = GameObject.Find("BattlePanel(Clone)/Card1");
	btn = GameObject.GetComponent(btnObj, "Button");
	evt = btn.onClick;
	evt.AddListener(evt, HandCards_2);
	
	btnObj = GameObject.Find("BattlePanel(Clone)/Card2");
	btn = GameObject.GetComponent(btnObj, "Button");
	evt = btn.onClick;
	evt.AddListener(evt, HandCards_3);
	
	btnObj = GameObject.Find("BattlePanel(Clone)/Card3");
	btn = GameObject.GetComponent(btnObj, "Button");
	evt = btn.onClick;
	evt.AddListener(evt, HandCards_4);
	
	btnObj = GameObject.Find("BattlePanel(Clone)/Card4");
	btn = GameObject.GetComponent(btnObj, "Button");
	evt = btn.onClick;
	evt.AddListener(evt, HandCards_5);
end

function BattlePanel_skip()
	Proxy4Lua.BattleSetup();
end

function BattlePanel_update()
	
end

function BattlePanel_tick()
	
end

function HandCards_1()
	Proxy4Lua.SelectCard4Ready(0);
end

function HandCards_2()
	Proxy4Lua.SelectCard4Ready(1);
end

function HandCards_3()
	Proxy4Lua.SelectCard4Ready(2);
end

function HandCards_4()
	Proxy4Lua.SelectCard4Ready(3);
end

function HandCards_5()
	Proxy4Lua.SelectCard4Ready(4);
end