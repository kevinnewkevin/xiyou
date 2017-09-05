using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class AnimHolder : MonoBehaviour {

    Animation _Animation;

    Queue<KeyValuePair<float, string>> _ActionQue;

    void Start()
    {
        if (_Animation == null)
            _Animation = GetComponent<Animation>();
    }

    public void Add(float delay, string clipName)
    {
        if (_Animation == null)
        {
            if (_ActionQue == null)
                _ActionQue = new Queue<KeyValuePair<float, string>>();

            _ActionQue.Enqueue(new KeyValuePair<float, string>(delay, clipName));
        }
        else
        {
            new Timer().Start(delay, (object o) => {
                Play(o.ToString());
            }, clipName);
        }
    }

    void Update()
    {
        if (_Animation == null)
            return;

        if (_ActionQue == null)
            return;

        KeyValuePair<float, string> kv;
        while(_ActionQue.Count > 0)
        {
            kv = _ActionQue.Dequeue();
            new Timer().Start(kv.Key, (object o) =>  {
                Play(o.ToString());
            }, kv.Value);
        }
    }

    void Play(string clipName)
    {
        _Animation.CrossFade(clipName);
    }
}
