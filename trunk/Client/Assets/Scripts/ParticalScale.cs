using UnityEngine;
using System.Collections;
using System.Collections.Generic;

public class ParticalScale : MonoBehaviour
{
    public float scanleSize = 1f;
    List<float> initialSize = new List<float>();

    void Awake()
    {
        ParticleSystem[] particles = gameObject.GetComponentsInChildren<ParticleSystem>();

        foreach(ParticleSystem particle in particles)
        {
            initialSize.Add(particle.startSize);
            ParticleSystemRenderer renderer = particle.GetComponent<ParticleSystemRenderer>();
            if (renderer)
            {
                initialSize.Add(renderer.lengthScale);
                initialSize.Add(renderer.velocityScale);
            }
        }
    }

    void Start()
    {
        int arrayIndex = 0;
        ParticleSystem[] particles = gameObject.GetComponentsInChildren<ParticleSystem>();
        foreach(ParticleSystem particle in particles)
        {
            particle.startSize = initialSize [arrayIndex++] * scanleSize;
//            ParticleSystemRenderer renderer = particle.GetComponent<ParticleSystemRenderer>();
//            if (renderer)
//            {
//                renderer.lengthScale = initialSize[arrayIndex++] * scanleSize;
//                renderer.velocityScale = initialSize[arrayIndex++] * scanleSize;
//            }
        }
    }
}