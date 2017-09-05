using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;

public class CameraEffect {

    static public Material _Mat;
    static bool _IsPlaying;
    static float _V;
    static bool _FadeIn;

    public delegate void FadeingCallback();
    static FadeingCallback fadeCallback;

    public static void Init()
    {
        _Mat = Resources.Load<Material>("Material/Fade");
        SceneManager.sceneLoaded += (arg0, arg1) => {
            SearchAllCamera();
        };
    }

    static void SearchAllCamera()
    {
        Camera[] cameras = GameObject.FindObjectsOfType<Camera>();
        CameraFade cf;
        for(int i=0; i < cameras.Length; ++i)
        {
            cf = cameras [i].GetComponent<CameraFade>();
            if(cf == null)
                cf = cameras [i].gameObject.AddComponent<CameraFade>();
        }
    }

    public static void Fade(FadeingCallback callback)
    {
        fadeCallback = callback;
        SearchAllCamera();
        _IsPlaying = true;
        _FadeIn = true;
    }

    public static void Update()
    {
        if (!_IsPlaying)
            return;
        
        if (_FadeIn)
        {
            _V -= Time.deltaTime;
            if (_V < -1f)
            {
                Debug.Log("Do something");
                if (fadeCallback != null)
                {
                    fadeCallback();
                    fadeCallback = null;
                }
            }
        }
        else
        {
            _V += Time.deltaTime;
            if (_V > 0f)
            {
                Debug.Log("over");
                _V = 0f;
                _IsPlaying = false;
            }
        }
        _Mat.SetFloat("_Float1", _V);
    }

    public static void Continue()
    {
        if (!_IsPlaying)
            return;
        
        _FadeIn = false;
    }
}
