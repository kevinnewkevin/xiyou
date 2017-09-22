using UnityEngine;
using FairyGUI;
using DG.Tweening;

public class EmitComponent : GComponent
{
    GLoader _symbolLoader;
    GTextField _numberText_cri;
	GTextField _numberText_norm;
	Transform _owner;

	const float OFFSET_ADDITION = 2.5f;
	static Vector2 JITTER_FACTOR = new Vector2(80, 80);

    GComponent gcom;
    GComponent criCom;
    GComponent normCom;
    GTextField crtTextfield;

	public EmitComponent()
	{
        gcom = UIPackage.CreateObject("zhandoushuzi", "ziti_com").asCom;
        this.AddChild(gcom);
        gcom.touchable = false;

        criCom = gcom.GetChild("n0").asCom;
        normCom = gcom.GetChild("n1").asCom;

        _numberText_cri = criCom.GetChild("n3").asTextField;
        _numberText_norm = normCom.GetChild("n3").asTextField;
	}

    public void SetHurt(Transform owner, int hurt, string special)
	{
		_owner = owner;

        bool isCritical = special.Equals("BE_Crit");

        crtTextfield = isCritical ? _numberText_cri : _numberText_norm;
        TextFormat tf = crtTextfield.textFormat;
        if (hurt < 0)
            tf.font = EmitManager.inst.hurtFont;
		else
            tf.font = EmitManager.inst.recoverFont;

        if (isCritical)
        {
            tf.font = EmitManager.inst.criticalFont;
            criCom.visible = true;
            normCom.visible = false;
        }
        else
        {
            criCom.visible = false;
            normCom.visible = true;
        }
        crtTextfield.textFormat = tf;
        crtTextfield.text = hurt.ToString();

//        if (!string.IsNullOrEmpty(special))
//            _symbolLoader.url = EmitManager.inst.specialSign + special;
//        else
//            _symbolLoader.url = "";

//		UpdateLayout();

//		this.alpha = 1;
		UpdatePosition(Vector2.zero);
//		Vector2 rnd = Vector2.Scale(UnityEngine.Random.insideUnitCircle, JITTER_FACTOR);
//		int toX = (int)rnd.x * 2;
//		int toY = (int)rnd.y * 2;

		EmitManager.inst.view.AddChild(this);
        new Timer().Start(2f, delegate{
            this.OnCompleted();
        });
//		DOTween.To(() => Vector2.zero, val => { this.UpdatePosition(val); }, new Vector2(toX, toY), 1f)
//			.SetTarget(this).OnComplete(this.OnCompleted);
//		this.TweenFade(0, 0.5f).SetDelay(0.5f);
	}

	void UpdateLayout()
	{
        this.SetSize(_symbolLoader.width + crtTextfield.width, Mathf.Max(_symbolLoader.height, crtTextfield.height));
        crtTextfield.SetXY(_symbolLoader.width > 0 ? (_symbolLoader.width + 2) : 0,
            (this.height - crtTextfield.height) / 2);
		_symbolLoader.y = (this.height - _symbolLoader.height) / 2;
	}

	void UpdatePosition(Vector2 pos)
	{
		Vector3 ownerPos = _owner.position;
		ownerPos.y += OFFSET_ADDITION;
		Vector3 screenPos = Camera.main.WorldToScreenPoint(ownerPos);
        screenPos.y = Screen.height - screenPos.y; //convert to Stage coordinates system

		Vector3 pt = GRoot.inst.GlobalToLocal(screenPos);
        this.SetXY(Mathf.RoundToInt(pt.x + pos.x - this.actualWidth / 2 + Random.Range(-60f, 60f)), Mathf.RoundToInt(pt.y + pos.y - this.height) + Random.Range(-60f, 60f));
	}

	void OnCompleted()
	{
		_owner = null;
		EmitManager.inst.view.RemoveChild(this);
		EmitManager.inst.ReturnComponent(this);
	}

	public void Cancel()
	{
		_owner = null;
		if (this.parent != null)
		{
			DOTween.Kill(this);
			EmitManager.inst.view.RemoveChild(this);
		}
		EmitManager.inst.ReturnComponent(this);
	}
}