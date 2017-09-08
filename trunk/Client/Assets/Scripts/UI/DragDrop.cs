using FairyGUI;
using UnityEngine;

public class DragDrop {

    GLoader _AgentLoader;
    object _SourceData;

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
    }

    public GLoader dragAgent
    {
        get { return _AgentLoader; }
    }

    public bool dragging
    {
        get { return _AgentLoader.parent != null; }
    }

    public void StartDrag(GObject source, string icon, object sourceData, int touchPointID = -1)
    {
        if (_AgentLoader.parent != null)
            return;

        _SourceData = sourceData;
        _AgentLoader.url = icon;
        GRoot.inst.AddChild(_AgentLoader);
        _AgentLoader.xy = GRoot.inst.GlobalToLocal(Stage.inst.GetTouchPosition(touchPointID));
        _AgentLoader.StartDrag(touchPointID);
    }

    public void Cancel()
    {
        if (_AgentLoader.parent != null)
        {
            _AgentLoader.StopDrag();
            GRoot.inst.RemoveChild(_AgentLoader);
            _SourceData = null;
        }
    }

    void OnDragEnd(EventContext context)
    {
        if (_AgentLoader.parent == null)
            return;

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
}
