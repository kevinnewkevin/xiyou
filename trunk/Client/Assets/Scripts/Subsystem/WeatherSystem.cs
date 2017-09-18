using System;
using UnityEngine;

public class WeatherSystem {

    static string _CrtWeather;
    static GameObject _WeatherObj;

    static float _MinGap;

    public static void Init()
    {
        _CrtWeather = "";
        _MinGap = Define.GetFloat("WeatherCheckTime");
    }

    public static void Update()
    {
        _MinGap -= Time.deltaTime;
        if (_MinGap < 0f)
        {
            CheckWeather();
            _MinGap = Define.GetFloat("WeatherCheckTime");
        }
    }

    static public void ForceCheck()
    {
        CheckWeather();
    }

    static void CheckWeather()
    {
        if (Camera.main == null)
            return;
        
        int hour = DateTime.Now.Hour;
        string nextWeather = WeatherData.GetWeather(hour);
        if (!_CrtWeather.Equals(nextWeather))
        {
            if (!string.IsNullOrEmpty(_CrtWeather))
                AssetLoader.UnloadAsset(_CrtWeather, true);

            if (_WeatherObj != null)
                GameObject.Destroy(_WeatherObj);
            
            _CrtWeather = "";
            _WeatherObj = null;

            _WeatherObj = AssetLoader.LoadAsset(nextWeather);
            if (_WeatherObj != null)
            {
                _WeatherObj.transform.parent = Camera.main.transform;
                _WeatherObj.transform.localPosition = Vector3.zero;
                _WeatherObj.transform.localScale = Vector3.one;
                _WeatherObj.transform.localRotation = Quaternion.identity;
                _CrtWeather = nextWeather;
            }
        }
    }
}
