local GameObject = UnityEngine.GameObject;

function LoginPanel_start()
	local btnObj = GameObject.Find("LoginPanel(Clone)");
	local btn = GameObject.GetComponent(btnObj, "Button");
	local evt = btn.onClick;
	evt.AddListener(evt, LoginPanel_Test);
	print("Login start");
end

function LoginPanel_Test()
    Proxy4Lua.CreatePlayer(1, "guowengui");
end

function LoginPanel_update()
	
end

function LoginPanel_tick()
	print("Login tick");
end