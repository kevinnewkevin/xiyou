using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using FairyGUI;

public class HeadBar {

    public GameObject _Headbar;

    GComponent _HeadBarCom;

    GProgressBar _BloodBar;

    Transition _Transition;

    GTextField _SkillName;

    GTextField _Name;

    GTextField _Title;

    GTextField _Lv;

    GComponent _HeadIconCom;

    GLoader _QuestIcon;

    GLoader _Turn;

    Actor _Root;

    GList _BuffList;

    float _HeightAdjust;

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
        DisplayData dData = DisplayData.GetData(root._DisplayID);
        if (dData != null)
            _HeightAdjust = dData._HeadBarHeight;
        _Headbar.transform.localPosition = new Vector3(0f, _HeightAdjust, 0f);
        _HeadBarCom = headbarpanel.ui;
        _BuffList = _HeadBarCom.GetChild("n7").asList;
        _BloodBar = _HeadBarCom.GetChild("n5").asProgress;
        _Lv = _HeadBarCom.GetChild("n3").asTextField;
        GComponent labelCom = _HeadBarCom.GetChild("n8").asCom;
        _Name = labelCom.GetChild("n0").asTextField;
        _Title = labelCom.GetChild("n1").asTextField;
        _HeadIconCom = _HeadBarCom.GetChild("n9").asCom;
        _QuestIcon = _HeadIconCom.GetChild("n4").asLoader;
        _Turn = _HeadBarCom.GetChild("n10").asLoader;
        _HeadBarCom.GetController("xuetiao").selectedIndex = state;
        //_Transition = _HeadBarCom.GetTransition("t0");
        //_SkillName = _HeadBarCom.GetChild("n11").asTextField;

        _IsDirty = true;
    }

    public void PopContent(int value, string special)
    {
        EmitManager.inst.Emit(_Root._ActorObj.transform, value, special);
    }

    public void DisplaySkill(string skillName)
    {
        if (_Transition.playing)
            _Transition.Stop();
        //_SkillName.text = skillName;
        _Transition.Play();
    }

    public void Update()
    {
        if (!_IsDirty)
            return;

        _BuffList.RemoveChildrenToPool();
        GObject item;
        BuffData data;

        string tail = "";
        int gapTurn = (Battle._Turn - _Root._BornTurn) % 3;
        if (gapTurn == 0)
            tail = "weiba_yi";
        else if(gapTurn == 1)
            tail = "weiba_er";
        else
            tail = "weiba_san";
        _Turn.url = "ui://xuetiao/" + tail;

        if (_Root.BuffList != null)
        {
            int line = (_Root.BuffList.Count / 6) + ((_Root.BuffList.Count % 6) > 0 ? 1 : 0);
            for (int i = 0; i < _Root.BuffList.Count; ++i)
            {
                item = _BuffList.AddItemFromPool(_BuffLoaderUrl);
                data = BuffData.GetData(_Root.BuffList [i]);
                item.asCom.GetChild("n0").asLoader.url = string.Format("ui://{0}", data._Icon);
            }
            if (line <= 0)
                line = 1;
            _Headbar.transform.localPosition = new Vector3(0f, _HeightAdjust + (line - 1) * 0.245f, 0f);
        }

        _BloodBar.TweenValue(((float)_Root._CrtValue / (float)_Root._MaxValue * 100), 0.5f);

        _Name.text = _Root._Name;
        _Title.text = _Root._Title;
        _Lv.text = _Root._StrLv.ToString();

        if (!_Root.HasQuest)
            _HeadIconCom.visible = false;

        _IsDirty = false;
    }

    public bool Visible
    {
        set
        {
            if(_Headbar != null)
                _Headbar.SetActive(value);
        }
    }
}
