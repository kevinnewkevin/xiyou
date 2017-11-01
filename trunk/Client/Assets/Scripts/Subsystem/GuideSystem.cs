using UnityEngine;
using FairyGUI;

public class GuideSystem  {

    static GComponent _GuideLayer;

    static ulong _Progress;

    static public void Init()
    {
        Define.LaunchUIBundle("xinshouyindao");
        _GuideLayer = UIPackage.CreateObject("xinshouyindao", "xinshouyindao_com").asCom;
        _GuideLayer.SetSize(GRoot.inst.width, GRoot.inst.height);
        _GuideLayer.AddRelation(GRoot.inst, RelationType.Size);
    }

    static public void OpenUI(string ui, Window win)
    {
        LuaManager.Call("guide.lua", "WhenUIOpen", ui, win);
    }

    static public void SpecialEvt(string type, params object[] par)
    {
        LuaManager.Call("guide.lua", "SpecialEvent", type, par);
    }

    static public void StartGuide(GObject aim)
    {
        if (aim == null)
            return;
        
        GRoot.inst.AddChild(_GuideLayer); //!!Before using TransformRect(or GlobalToLocal), the object must be added first
        Rect rect;
        if(aim.parent != null && aim.parent.parent != null && aim.parent.parent is Window)
            rect = aim.TransformRect(new Rect(aim.pivotX * -aim.width, aim.pivotY * -aim.height, aim.width, aim.height), _GuideLayer);
        else
            rect = aim.TransformRect(new Rect(0f, 0f, aim.width, aim.height), _GuideLayer);
//        Rect rect = new Rect(aim.position.x, aim.position.y, aim.width, aim.height);
        
        GObject window = _GuideLayer.GetChild("n5");
        window.pivotX = 0f;
        window.pivotY = 0f;
        window.size = new Vector2((int)rect.size.x, (int)rect.size.y);
//        window.SetXY((int)rect.position.x, (int)rect.position.y);
        window.TweenMove(new Vector2((int)rect.position.x, (int)rect.position.y), 0.5f);
    }

    static public void StartGuide(GObject aim, float width, float height)
    {
        if (aim == null)
            return;

        GRoot.inst.AddChild(_GuideLayer); //!!Before using TransformRect(or GlobalToLocal), the object must be added first
        GRoot.inst.AddChild(_GuideLayer); //!!Before using TransformRect(or GlobalToLocal), the object must be added first
        Rect rect;
//        if(aim.parent != null && aim.parent.parent != null && aim.parent.parent is Window)
//            rect = aim.TransformRect(new Rect(aim.pivotX * -aim.width, aim.pivotY * -aim.height, width, height), _GuideLayer);
//        else
        rect = aim.TransformRect(new Rect((aim.width - width) / 2, (aim.height - height) / 2, width, height), _GuideLayer);
//        
//        Rect rect = new Rect(aim.x + aim.width / 2, aim.y + aim.height / 2, width, height);

        GObject window = _GuideLayer.GetChild("n5");
        window.pivot = aim.pivot;
        window.size = new Vector2((int)rect.size.x, (int)rect.size.y);
        //        window.SetXY((int)rect.position.x, (int)rect.position.y);
        window.TweenMove(new Vector2((int)rect.position.x, (int)rect.position.y), 0.5f);
    }

    static public void StartGuideInScene(GameObject go, float width, float height)
    {
        if (go == null)
            return;
        
        GRoot.inst.AddChild(_GuideLayer); //!!Before using TransformRect(or GlobalToLocal), the object must be added first

        Vector3 ownerPos = go.transform.position;
        Vector3 screenPos = Camera.main.WorldToScreenPoint(ownerPos);
        screenPos.y = Screen.height - screenPos.y; //convert to Stage coordinates system
        Vector3 pt = GRoot.inst.GlobalToLocal(screenPos);

        GObject window = _GuideLayer.GetChild("n5");
        window.size = new Vector2((int)width, (int)height);
        //        window.SetXY((int)rect.position.x, (int)rect.position.y);
        window.TweenMove(new Vector2((int)pt.x - width / 2, (int)pt.y - height / 2), 0.5f);
    }

    static public void ClearGuide()
    {
        _GuideLayer.RemoveFromParent();
    }

    static public bool IsNotFinish(int idx)
    {
        return (_Progress >> idx) % 2 == 0;
    }

    static public void SetFinish(int idx)
    {
        _Progress |= (ulong)1 << idx;
    }

    static public void SetProgress(ulong progress)
    {
        _Progress = progress;
    }

    static public void SyncProgress()
    {
        NetWoking.S.NewPlayerGuide(_Progress);
    }
}
