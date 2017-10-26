using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class AudioSystem {

    static AudioSource _BackgroundSound;
    static AudioSource _VoiceSound;
    static List<EffectSound> _EffectSoundPool;

    static Dictionary<string, AudioClip> _ClipPool;

    static float _MusicVolum = 1f;
    static float _SoundVolum = 1f;
    static float _VoiceVolum = 1f;

    static public void Init()
    {
        _ClipPool = new Dictionary<string, AudioClip>();
        _EffectSoundPool = new List<EffectSound>();
    }

    static public void PlayBackground(string music)
    {
        if (_BackgroundSound == null)
        {
            GameObject go = new GameObject();
            _BackgroundSound = go.AddComponent<AudioSource>();
            MonoBehaviour.DontDestroyOnLoad(go);
        }
        else
        {
            _BackgroundSound.Stop();
            _BackgroundSound.clip = null;
        }

        _BackgroundSound.clip = GetClip(music);
        _BackgroundSound.loop = true;
        _BackgroundSound.pitch = 1f;
        _BackgroundSound.volume = _MusicVolum;
        _BackgroundSound.Play();
    }

    static public void PlayVoice(string voice)
    {
        if (_VoiceSound == null)
        {
            GameObject go = new GameObject();
            _VoiceSound = go.AddComponent<AudioSource>();
            MonoBehaviour.DontDestroyOnLoad(go);
        }
        else
        {
            _VoiceSound.Stop();
            _VoiceSound.clip = null;
        }

        _VoiceSound.clip = GetClip(voice);
        _VoiceSound.loop = true;
        _VoiceSound.pitch = 1f;
        _VoiceSound.volume = _MusicVolum;
        _VoiceSound.Play();
    }

    static public void PlayEffect(string effect)
    {
        EffectSound es = null;
        for(int i=0; i < _EffectSoundPool.Count; ++i)
        {
            if (_EffectSoundPool [i]._IsOver)
            {
                es = _EffectSoundPool [i];
                es._IsOver = false;
                break;
            }
        }
        if (es == null)
        {
            es = new EffectSound();
            _EffectSoundPool.Add(es);
        }

        es.Play(effect);
    }

    static public AudioClip GetClip(string clipName)
    {
        if (!_ClipPool.ContainsKey(clipName))
            _ClipPool.Add(clipName, AssetLoader.LoadAudio(clipName));
        return _ClipPool[clipName];
    }

    static public float MusicVolum
    {
        get
        {
            return _MusicVolum;
        }
        set
        {
            _MusicVolum = value;
            if (_BackgroundSound != null)
                _BackgroundSound.volume = _MusicVolum;
        }
    }

    static public float SoundVolum
    {
        get
        {
            return _SoundVolum;
        }
        set
        {
            _SoundVolum = value;
            for(int i=0; i < _EffectSoundPool.Count; ++i)
            {
                if(_EffectSoundPool [i] != null)
                    _EffectSoundPool [i].Volum(_SoundVolum);
            }
        }
    }

    static public float VoiceVolum
    {
        get
        {
            return _VoiceVolum;
        }
        set
        {
            _VoiceVolum = value;
            if (_VoiceSound != null)
                _VoiceSound.volume = _VoiceVolum;
        }
    }
}

public class EffectSound
{
    public AudioSource _EffectSound;
    public bool _IsOver;

    public void Play(string effect)
    {
        if (_EffectSound == null)
        {
            GameObject go = new GameObject();
            _EffectSound = go.AddComponent<AudioSource>();
            MonoBehaviour.DontDestroyOnLoad(go);
            _IsOver = false;
        }

        _EffectSound.clip = AudioSystem.GetClip(effect);
        _EffectSound.loop = false;
        _EffectSound.pitch = 1f;
        _EffectSound.volume = AudioSystem.SoundVolum;
        _EffectSound.Play();
        new Timer().Start(_EffectSound.clip.length, delegate {
            if (_EffectSound != null)
                _EffectSound.Stop();
            _IsOver = true;
        });
    }

    public void Volum(float volum)
    {
        if (!_IsOver && _EffectSound != null)
            _EffectSound.volume = volum;
    }
}
