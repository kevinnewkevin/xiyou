using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class GameObjectHandler : MonoBehaviour {

    public List<Actor.CallBackHandler> callbacks = null;
    public void HandlerFunction(string index)
    {
        int idx = int.Parse(index);
        if (callbacks == null || callbacks.Count == 0)
            return;

        if (idx < 0 || idx >= callbacks.Count)
            return;

        if(callbacks[idx] != null)
            callbacks[idx]();
        callbacks.RemoveAt(idx);
    }

    public int LaunchHandler(Actor.CallBackHandler handler)
    {
        if (callbacks == null)
            callbacks = new List<Actor.CallBackHandler>();
        
        callbacks.Add(handler);
        return callbacks.Count - 1;
    }
}
