using System.Collections.Generic;
using UnityEngine;
using FairyGUI;

public class EmitManager
{
	static EmitManager _instance;
	public static EmitManager inst
	{
		get
		{
			if (_instance == null)
				_instance = new EmitManager();
			return _instance;
		}
	}

	public string hurtFont;
	public string recoverFont;
    public string criticalFont;
	public string specialSign;

	public GComponent view { get; private set; }

	private readonly Stack<EmitComponent> _componentPool = new Stack<EmitComponent>();

	public EmitManager()
	{
        hurtFont = "ui://zhandoushuzi/putong";
        recoverFont = "ui://zhandoushuzi/jiaxue";
        criticalFont = "ui://zhandoushuzi/baoji";
        specialSign = "ui://zhandoushuzi/";

		view = new GComponent();
		GRoot.inst.AddChild(view);
	}

    public void Emit(Transform owner, int hurt, string special, bool isBuff = false)
	{
		EmitComponent ec;
		if (_componentPool.Count > 0)
			ec = _componentPool.Pop();
		else
			ec = new EmitComponent();
        ec.SetHurt(owner, hurt, special, isBuff);
	}

	public void ReturnComponent(EmitComponent com)
	{
		_componentPool.Push(com);
	}
}