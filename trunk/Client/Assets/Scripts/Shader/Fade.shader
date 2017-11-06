// Upgrade NOTE: replaced 'mul(UNITY_MATRIX_MVP,*)' with 'UnityObjectToClipPos(*)'

Shader "Custom/Fade" {  
    Properties {  
        _MainTex ("Base (RGB)", 2D) = "white" {}  
        _Color("Main Color",Color) = (0.5,0.5,0.5,0.5)  
    }  
  
    SubShader {  
        Pass {  
            CGPROGRAM  
            	#pragma vertex vert_img
            	#pragma fragment frag
            	#include "UnityCG.cginc"

            	fixed4 _Color;
            	sampler2D _MainTex;

            	float4 frag(v2f_img i) : COLOR{
            		float4 col = tex2D(_MainTex, i.uv) * _Color;
            		return col;
            	}
//  
//            #pragma vertex vert  
//            #pragma fragment frag  
//  
//            uniform sampler2D _MainTex;  
//            uniform float _Float1;   
//              
//            struct Input {  
//                float4 pos : POSITION;  
//                float2 uv : TEXCOORD0;  
//            };  
//  
//            struct v2f {  
//                float4 pos : POSITION;  
//                float2 uv : TEXCOORD0;  
//            };  
//  
//            v2f vert( Input i) {  
//                v2f o;  
//                o.pos = UnityObjectToClipPos (i.pos);  
//                o.uv = i.uv;  
//                return o;  
//            }  
//  
//            float4 frag (v2f i) : COLOR {              
//                float4 outColor;                  
//                outColor = tex2D( _MainTex, i.uv) + _Float1;  
//                return outColor;  
//            }  
            ENDCG  
        }  
    }  
    Fallback off  
}