using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class LoginUI {

    static Button btn;
    static GameObject selfObj;

    static public void Start(GameObject go)
    {
        selfObj = go;
        btn = selfObj.GetComponent<Button>();
        btn.onClick.AddListener(LoginSuccess);

    }

	static public void LoginSuccess()
    {
        SceneLoader.LoadScene(Define.SCENE_MAIN);
    }
}
