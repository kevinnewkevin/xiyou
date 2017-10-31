using UnityEngine;
using FairyGUI;
using DG.Tweening;

public class EmitComponent : GComponent
{
    GLoader _symbolLoader;
    GTextField _numberText_cri;
	GTextField _numberText_norm;
    GTextField _numberText_buff;
	Transform _owner;

	const float OFFSET_ADDITION = 2.5f;
    const float OFFSET_ADDITION_BUFF = 1.1f;
	static Vector2 JITTER_FACTOR = new Vector2(80, 80);

    GComponent gcom;
    GComponent criCom;
    GComponent normCom;
    GComponent buffCom;
    GTextField crtTextfield;

    public bool _IsBuff;

	public EmitComponent()
	{
        gcom = UIPackage.CreateObject("zhandoushuzi", "ziti_com").asCom;
        this.AddChild(gcom);
        gcom.touchable = false;

        criCom = gcom.GetChild("n0").asCom;
        normCom = gcom.GetChild("n1").asCom;
        buffCom = gcom.GetChild("n2").asCom;

        _numberText_cri = criCom.GetChild("n3").asTextField;
        _numberText_norm = normCom.GetChild("n3").asTextField;
        _numberText_buff = buffCom.GetChild("n12").asTextField;
	}

    public void SetHurt(Transform owner, int hurt, string special, bool isBuff = false)
	{
		_owner = owner;

        _IsBuff = isBuff;
        if (isBuff == false)
        {
            buffCom.visible = false;
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
        }
        else
        {
            _numberText_buff.text = special;
            buffCom.visible = true;
            criCom.visible = false;
            normCom.visible = false;
        }

        UpdatePosition(new Vector2(0f, EmitManager.inst._BuffCount * buffCom.height));
		EmitManager.inst.view.AddChild(this);
        new Timer().Start(2f, delegate{
            this.OnCompleted();
        });
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
        if (_IsBuff)
        {
            Vector3 ownerPos = _owner.position;
            ownerPos.y += OFFSET_ADDITION_BUFF;
            Vector3 screenPos = Camera.main.WorldToScreenPoint(ownerPos);
            screenPos.y = Screen.height - screenPos.y; //convert to Stage coordinates system

            Vector3 pt = GRoot.inst.GlobalToLocal(screenPos);
            this.SetXY(Mathf.RoundToInt(pt.x + pos.x - this.actualWidth / 2), Mathf.RoundToInt(pt.y + pos.y - this.height));
        }
        else
        {
            Vector3 ownerPos = _owner.position;
            ownerPos.y += OFFSET_ADDITION;
            Vector3 screenPos = Camera.main.WorldToScreenPoint(ownerPos);
            screenPos.y = Screen.height - screenPos.y; //convert to Stage coordinates system

            Vector3 pt = GRoot.inst.GlobalToLocal(screenPos);
            this.SetXY(Mathf.RoundToInt(pt.x + pos.x - this.actualWidth / 2 + Random.Range(-60f, 60f)), Mathf.RoundToInt(pt.y + pos.y - this.height) + Random.Range(-60f, 60f));
        }
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