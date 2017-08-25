using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using FairyGUI;

public class HeadBar {

    GameObject _Headbar;

    GComponent _HeadBarCom;

    GProgressBar _BloodBar;

    GTextField _Name;

    GTextField _Title;

    GComponent _HeadIconCom;

    GLoader _QuestIcon;

    Actor _Root;

    GList _BuffList;

    string _BuffLoaderUrl = "ui://xuetiao/buff_loader";

    public bool _IsDirty;

    public HeadBar(Actor root, int state)
    {
        if (root._ActorObj == null)
            return;

        UIPackage.AddPackage("UI/EmitNumbers");
        Define.LaunchUIBundle("xuetiao");

        _Root = root;
        _Headbar = new GameObject();
        _Headbar.AddComponent<Billboard>();
        UIPanel headbarpanel = _Headbar.AddComponent<UIPanel>();
        headbarpanel.componentName = "xuetiao_com";
        headbarpanel.packageName = "xuetiao";
        _Headbar.transform.parent = root._ActorObj.transform;
        _Headbar.transform.localScale = Vector3.one;
        _Headbar.transform.localPosition = new Vector3(0f, 2f, 0f);
        _HeadBarCom = headbarpanel.ui;
        _BuffList = _HeadBarCom.GetChild("n7").asList;
        _BloodBar = _HeadBarCom.GetChild("n5").asProgress;
        GComponent labelCom = _HeadBarCom.GetChild("n8").asCom;
        _Name = labelCom.GetChild("n0").asTextField;
        _Title = labelCom.GetChild("n1").asTextField;
        _HeadIconCom = _HeadBarCom.GetChild("n9").asCom;
        _QuestIcon = _HeadIconCom.GetChild("n4").asLoader;
        _HeadBarCom.GetController("xuetiao").selectedIndex = state;

        _IsDirty = true;
    }

    public void PopContent(int value)
    {
        EmitManager.inst.Emit(_Root._ActorObj.transform, 0, value, UnityEngine.Random.Range(0, 10) == 5);
    }

    public void Update()
    {
        if (!_IsDirty)
            return;

        _BuffList.RemoveChildrenToPool();
        GObject item;
        BuffData data;
        if (_Root.BuffList != null)
        {
            int line = (_Root.BuffList.Count / 6) + ((_Root.BuffList.Count % 6) > 0? 1: 0);
            for(int i=0; i < _Root.BuffList.Count; ++i)
            {
                item = _BuffList.AddItemFromPool(_BuffLoaderUrl);
                data = BuffData.GetData(_Root.BuffList[i]);
                item.asCom.GetChild("n0").asLoader.url = string.Format("ui://{0}", data._Icon);
            }
            _Headbar.transform.localPosition = new Vector3(0f, 2f + (line - 1) * 0.245f, 0f);
        }

        _BloodBar.TweenValue(((float)_Root._CrtValue / (float)_Root._MaxValue * 100), 0.5f);

        _Name.text = _Root._Name;
        _Title.text = _Root._Title;

        if (!_Root.HasQuest)
            _HeadIconCom.visible = false;

        _IsDirty = false;
    }
}
