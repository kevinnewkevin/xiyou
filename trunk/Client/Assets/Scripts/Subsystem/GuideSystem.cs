using FairyGUI;

public class GuideSystem  {

    static public void Init()
    {
        
    }

    static public void OpenUI(string ui, Window win)
    {
        LuaManager.Call("global.lua", "WhenUIOpen", ui, win);
    }

    static public void StartGuide(GObject aim)
    {
        
    }

    static public bool IsNotFinish(int idx)
    {
        return true;
    }

    static public void SetFinish(int idx)
    {
        
    }
}
