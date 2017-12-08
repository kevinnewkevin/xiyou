using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ExceptionHandle {
    
    static public void Update()
    {
        if (NetWoking._LastErrorCode != 0)
        {
            LuaManager.Call("global.lua", "NetWorkException", NetWoking._LastErrorCode);
            NetWoking._LastErrorCode = 0;
        }
    }
}
