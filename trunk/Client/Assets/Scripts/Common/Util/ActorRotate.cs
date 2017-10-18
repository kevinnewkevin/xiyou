using UnityEngine;
using System.Collections;
using FairyGUI;

public class ActorRotate : MonoBehaviour {

	// Use this for initialization
	void Start () {
        Stage.inst.onTouchBegin.Add(OnPress);
        Stage.inst.onTouchEnd.Add(OnEnd);
	}

    float originX_;
    float currentX_;

    float time_;

    bool beginRot_ = false;

    void OnPress()
    {
        originX_ = Input.mousePosition.x;
        beginRot_ = true;
    }

    void OnEnd()
    {
        currentX_ = 0f;
        originX_ = 0f;
        time_ = 0f;

        beginRot_ = false;
    }
	
	// Update is called once per frame
	void Update ()
    {
        if (beginRot_)
        {
            currentX_ = Input.mousePosition.x;
            if (Mathf.Approximately(currentX_, originX_))
                return;

            if (gameObject != null)
            {
//                float rotV = currentX_ > originX_ ? -1f : 1f;
//                gameObject.transform.Rotate(Vector3.up, Mathf.Rad2Deg * 0.02f * Mathf.PI * rotV);
                gameObject.transform.Rotate(Vector3.up, (originX_ - currentX_) % 360f * 0.3f);
            }
            else
            {
                Stage.inst.onTouchBegin.Remove(OnPress);
                Stage.inst.onTouchEnd.Remove(OnEnd);
            }
            originX_ = currentX_;
        }
	}

    void OnDestroy()
    {
        Stage.inst.onTouchBegin.Remove(OnPress);
        Stage.inst.onTouchEnd.Remove(OnEnd);
    }
}
