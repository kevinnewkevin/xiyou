using UnityEngine;
using FairyGUI;

public class GuideSystem  {

    static GComponent _GuideLayer;

    static GObject _ContentLayer;

    static GObject _ContentBg;

    static GTextField _TextField;

    static ulong _Progress;

    static string _SpecialEvt = "";

    static public void Init()
    {
        Define.LaunchUIBundle("xinshouyindao");
        _GuideLayer = UIPackage.CreateObject("xinshouyindao", "xinshouyindao_com").asCom;
        _GuideLayer.SetSize(GRoot.inst.width, GRoot.inst.height);
        _GuideLayer.AddRelation(GRoot.inst, RelationType.Size);

        _ContentLayer = _GuideLayer.GetChild("n8");
        _ContentBg = _GuideLayer.GetChild("n6");
        _TextField = _GuideLayer.GetChild("n7").asTextField;
    }

    static public void OpenUI(string ui, Window win)
    {
        LuaManager.Call("guide.lua", "WhenUIOpen", ui, win);
    }

    static public void SpecialEvt(string type, params object[] par)
    {
        LuaManager.Call("guide.lua", "SpecialEvent", type, par);
    }

    static public void StartGuide(GObject aim, string content = "")
    {
        if (aim == null)
            return;
        
        GRoot.inst.AddChild(_GuideLayer); //!!Before using TransformRect(or GlobalToLocal), the object must be added first
        Rect rect;
        if(aim.parent != null && aim.parent.parent != null && aim.parent.parent is Window)
            rect = aim.TransformRect(new Rect(aim.pivotX * -aim.width, aim.pivotY * -aim.height, aim.width, aim.height), _GuideLayer);
        else
            rect = aim.TransformRect(new Rect(0f, 0f, aim.width, aim.height), _GuideLayer);
        
        GObject window = _GuideLayer.GetChild("n5");
        window.pivotX = 0f;
        window.pivotY = 0f;
        window.size = new Vector2((int)rect.size.x, (int)rect.size.y);
//        window.SetXY((int)rect.position.x, (int)rect.position.y);
        window.TweenMove(new Vector2((int)rect.position.x, (int)rect.position.y), 0.5f);

        TypeEffectContent(content);
    }

    static public void StartGuide(GObject aim, float width, float height, string content = "")
    {
        if (aim == null)
            return;

        GRoot.inst.AddChild(_GuideLayer); //!!Before using TransformRect(or GlobalToLocal), the object must be added first
        Rect rect;
        rect = aim.TransformRect(new Rect((aim.width - width) / 2, (aim.height - height) / 2, width, height), _GuideLayer);

        GObject window = _GuideLayer.GetChild("n5");
        window.pivot = aim.pivot;
        window.size = new Vector2((int)rect.size.x, (int)rect.size.y);
        //        window.SetXY((int)rect.position.x, (int)rect.position.y);
        window.TweenMove(new Vector2((int)rect.position.x, (int)rect.position.y), 0.5f);

        TypeEffectContent(content);
    }

    static public void StartGuideInScene(GameObject go, float width, float height, string content = "")
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

        TypeEffectContent(content);
    }

    static public void StartGuide(GObject aim, string content, string specialevt)
    {
        if (aim == null)
            return;

        GRoot.inst.AddChild(_GuideLayer); //!!Before using TransformRect(or GlobalToLocal), the object must be added first
        Rect rect;
        if(aim.parent != null && aim.parent.parent != null && aim.parent.parent is Window)
            rect = aim.TransformRect(new Rect(aim.pivotX * -aim.width, aim.pivotY * -aim.height, aim.width, aim.height), _GuideLayer);
        else
            rect = aim.TransformRect(new Rect(0f, 0f, aim.width, aim.height), _GuideLayer);

        GObject window = _GuideLayer.GetChild("n5");
        window.pivotX = 0f;
        window.pivotY = 0f;
        window.size = new Vector2((int)rect.size.x, (int)rect.size.y);
        //        window.SetXY((int)rect.position.x, (int)rect.position.y);
        window.TweenMove(new Vector2((int)rect.position.x, (int)rect.position.y), 0.5f);

        TypeEffectContent(content);
        _SpecialEvt = specialevt;
        Stage.inst.onTouchEnd.Add(OnTouchEnd);
    }

    static void TypeEffectContent(string content)
    {
        if (string.IsNullOrEmpty(content))
        {
            _ContentLayer.visible = false;
            _ContentBg.visible = false;
            _TextField.visible = false;
            return;
        }

        _ContentLayer.visible = true;
        _ContentBg.visible = true;
        _TextField.visible = true;
        TypingEffect te = new TypingEffect(_TextField);
        _TextField.text = content;
        te.Start();
        te.PrintAll(0.05f);

        new Timer().Start(content.Length * 0.05f, delegate
        {
            _ContentLayer.visible = false;
        });
    }

    static void OnTouchEnd()
    {
        ClearGuide();
        if(!string.IsNullOrEmpty(_SpecialEvt))
            SpecialEvt(_SpecialEvt);
        _SpecialEvt = "";
        Stage.inst.onTouchEnd.Remove(OnTouchEnd);
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
        SyncProgress();
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
