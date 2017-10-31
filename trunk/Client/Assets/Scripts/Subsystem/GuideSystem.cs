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
        Rect rect = aim.TransformRect(new Rect(0, 0, aim.width, aim.height), _GuideLayer);

        int plusx = 0;
        int plusy = 0;
//        if (aim.pivotX == 0.5f)
//            plusx = (int)(aim.width / 2);
//        if (aim.pivotY == 0.5f)
//            plusy = (int)(aim.height / 2);
        
        GObject window = _GuideLayer.GetChild("n5");
        window.size = new Vector2((int)rect.size.x, (int)rect.size.y);
//        window.SetXY((int)rect.position.x, (int)rect.position.y);
        window.TweenMove(new Vector2((int)rect.position.x - plusx, (int)rect.position.y - plusy), 0.5f);
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
}
