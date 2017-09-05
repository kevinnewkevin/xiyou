using UnityEngine;
using System.Collections.Generic;

public class AssetCounter {

    static Dictionary<string, AssetInfo> _AssetCounters = new Dictionary<string, AssetInfo>();

    static public void AddRef(string name, AssetBundle bundle)
    {
        if (!_AssetCounters.ContainsKey(name))
            _AssetCounters.Add(name, new AssetInfo(name, bundle, 1));
        else
            _AssetCounters [name]._Refcount += 1;
    }

    static public bool Excist(string name)
    {
        if (_AssetCounters.ContainsKey(name))
            return _AssetCounters [name]._Refcount > 0;
        return false;
    }

    static public AssetBundle GetBundle(string name)
    {
        if (_AssetCounters.ContainsKey(name))
        {
            _AssetCounters [name]._Refcount += 1;
            return _AssetCounters [name]._Bundle;
        }
        return null;
    }

    static public void DelRef(string name)
    {
        if (_AssetCounters.ContainsKey(name))
        {
            _AssetCounters [name]._Refcount -= 1;
            if (_AssetCounters [name]._Refcount <= 0)
            {
                _AssetCounters [name].Dispose();
                _AssetCounters.Remove(name);
                Debug.Log(" Delete unused asset: " + name);
            }
        }
    }

    static public void Dispose(string name)
    {
        if (_AssetCounters.ContainsKey(name))
        {
            _AssetCounters [name].Dispose();
            _AssetCounters.Remove(name);
            Debug.Log(" Force Delete unused asset: " + name);
        }
    }
}

class AssetInfo
{
    public AssetBundle  _Bundle;
    public string       _Name;
    public int          _Refcount;

    public AssetInfo(string name, AssetBundle bundle, int refcount)
    {
        _Name = name;
        _Bundle = bundle;
        _Refcount = refcount;
    }

    public void Dispose()
    {
        if (_Bundle != null)
            _Bundle.Unload(true);
    }
}
