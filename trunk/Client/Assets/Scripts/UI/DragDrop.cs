using FairyGUI;
using UnityEngine;

public class DragDrop {

    GLoader _AgentLoader;
    GGraph _AgentGraph;
    object _SourceData;

    string _AssetPath;

    static DragDrop _Inst;
    static public DragDrop inst
    {
        get
        {
            if (_Inst == null)
                _Inst = new DragDrop();
            return _Inst;
        }
    }

    public DragDrop()
    {
        _AgentLoader = (GLoader)UIObjectFactory.NewObject("loader");
        _AgentLoader.gameObjectName = "DragDropAgent";
        _AgentLoader.SetHome(GRoot.inst);
        _AgentLoader.touchable = false;
        _AgentLoader.draggable = true;
        _AgentLoader.SetSize(100, 100);
        _AgentLoader.SetPivot(0.5f, 0.5f, true);
        _AgentLoader.align = AlignType.Center;
        _AgentLoader.verticalAlign = VertAlignType.Middle;
        _AgentLoader.sortingOrder = int.MaxValue;
        _AgentLoader.onDragEnd.Add(OnDragEnd);
        _AgentLoader.alpha = 0.7f;

        _AgentGraph = (GGraph)UIObjectFactory.NewObject("graph");
        _AgentGraph.gameObjectName = "DragDropAgent";
        _AgentGraph.SetHome(GRoot.inst);
        _AgentGraph.touchable = false;
        _AgentGraph.draggable = true;
        _AgentGraph.SetSize(100, 100);
        _AgentGraph.SetPivot(0.5f, 1f, true);
        _AgentGraph.sortingOrder = int.MinValue;
        _AgentGraph.onDragEnd.Add(OnDragEnd);
    }

    public GLoader dragAgentLoader
    {
        get { return _AgentLoader; }
    }

    public GGraph dragAgentGraph
    {
        get { return _AgentGraph; }
    }

    public bool dragging
    {
        get { return _AgentLoader.parent != null || _AgentGraph.parent != null; }
    }

    public void StartDrag(GObject source, string path, object sourceData, int touchPointID = -1)
    {
        if (path.Contains("ui://"))
        {
            if (_AgentLoader.parent == null)
            {
                _SourceData = sourceData;
                _AgentLoader.url = path;
                GRoot.inst.AddChild(_AgentLoader);
                _AgentLoader.xy = GRoot.inst.GlobalToLocal(Stage.inst.GetTouchPosition(touchPointID));
                _AgentLoader.StartDrag(touchPointID);
            }
        }
        else
        {
            if (_AgentGraph.parent == null)
            {
                _SourceData = sourceData;
                _AgentGraph.SetNativeObject(Proxy4Lua.GetAssetGameObject(path));
                GRoot.inst.AddChild(_AgentGraph);
                _AgentGraph.xy = GRoot.inst.GlobalToLocal(Stage.inst.GetTouchPosition(touchPointID));
                _AgentGraph.StartDrag(touchPointID);
                _AssetPath = path;
            }
        }
    }

    public void Cancel()
    {
        if (_AgentLoader.parent != null)
        {
            _AgentLoader.StopDrag();
            GRoot.inst.RemoveChild(_AgentLoader);
            _SourceData = null;
        }

        if (_AgentGraph.parent != null)
        {
            _AgentGraph.visible = true;
            _AgentGraph.StopDrag();
            GRoot.inst.RemoveChild(_AgentGraph);
            _SourceData = null;
        }
    }

    void OnDragEnd(EventContext context)
    {
        if (_AgentLoader.parent != null)
        {
            GRoot.inst.RemoveChild(_AgentLoader);

            object sourcedata = _SourceData;
            _SourceData = null;

            GObject obj = GRoot.inst.touchTarget;
            while(obj != null)
            {
                if (obj is GComponent)
                {
                    if (!((GComponent)obj).onDrop.isEmpty)
                    {
                        obj.RequestFocus();
                        ((GComponent)obj).onDrop.Call(sourcedata);
                        return;
                    }
                }
                obj = obj.parent;
            }
        }

        if (_AgentGraph.parent != null)
        {
            _AgentGraph.visible = true;
            GRoot.inst.RemoveChild(_AgentGraph);

            object sourcedata = _SourceData;
            _SourceData = null;

            GObject obj = GRoot.inst.touchTarget;
            while(obj != null)
            {
                if (obj is GComponent)
                {
                    if (!((GComponent)obj).onDrop.isEmpty)
                    {
                        obj.RequestFocus();
                        ((GComponent)obj).onDrop.Call(sourcedata);
                        return;
                    }
                }
                obj = obj.parent;
            }
        }

        if (string.IsNullOrEmpty(_AssetPath))
        {
            AssetLoader.UnloadAsset(_AssetPath);
            _AssetPath = "";
        }

        //傻逼的代码by nick
        UIWindow win = UIManager.GetUI("BattlePanel");
        if(win != null)
            win.Call("NormalCard");

        if (dragAgentGraph != null && dragAgentGraph.visible)
            Battle._SelectedHandCardInstID = 0;
    }
}
