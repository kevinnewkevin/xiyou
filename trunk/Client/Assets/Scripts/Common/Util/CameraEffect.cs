using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
using FairyGUI;

public class CameraEffect {

    static public GGraph _Mat;
    static bool _IsPlaying;
    static float _V;
    static bool _FadeIn;

    public delegate void FadeingCallback();
    static FadeingCallback fadeCallback;

    public delegate void FadedCallback();
    static FadedCallback fadeCallback2;

    public static void Init()
    {
        _Mat = new GGraph();
        _Mat.DrawRect(GRoot.inst.width, GRoot.inst.height, 0, new Color(0f, 0f, 0f, 0f), new Color(0f, 0f, 0f, 0f));
        GRoot.inst.AddChild(_Mat);
        _Mat.touchable = false;
        _Mat.sortingOrder = int.MaxValue;
        _Mat.AddRelation(GRoot.inst, RelationType.Size);
        SceneManager.sceneLoaded += OnSceneLoaded;
    }

    static void SearchAllCamera()
    {
//        Camera[] cameras = GameObject.FindObjectsOfType<Camera>();
//        CameraFade cf;
//        for(int i=0; i < cameras.Length; ++i)
//        {
//            cf = cameras [i].GetComponent<CameraFade>();
//            if(cf == null)
//                cf = cameras [i].gameObject.AddComponent<CameraFade>();
//        }
    }

    public static void Fade(FadeingCallback callback, FadedCallback callback2 = null)
    {
        fadeCallback = callback;
        fadeCallback2 = callback2;
//        SearchAllCamera();
        _IsPlaying = true;
        _FadeIn = true;
    }

    public static void Update()
    {
        if (_Mat == null)
        {
            Fini();
            Init();
            return;
        }
        if (!_IsPlaying)
            return;

        float alpha;
        if (_FadeIn)
        {
            _V += Time.deltaTime;
            if (_V > 1f)
            {
                if (fadeCallback != null)
                {
                    fadeCallback();
                    fadeCallback = null;
                }
            }
        }
        else
        {
            _V -= Time.deltaTime;
            if (_V < 0f)
            {
                _V = 0f;
                _IsPlaying = false;
                if (fadeCallback2 != null)
                {
                    fadeCallback2();
                    fadeCallback2 = null;
                }
            }
        }

        _Mat.color = new Color(0f, 0f, 0f, _V);
    }

    public static void OnSceneLoaded(Scene arg0, LoadSceneMode arg1)
    {
//        SearchAllCamera();
        Continue();
    }

    public static void Continue()
    {
        if (!_IsPlaying)
            return;
        
        _FadeIn = false;
    }

    public static void Fini()
    {
        SceneManager.sceneLoaded -= OnSceneLoaded;
        fadeCallback = null;
        fadeCallback2 = null;
    }
}
