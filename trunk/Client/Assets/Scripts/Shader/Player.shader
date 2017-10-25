// Upgrade NOTE: replaced 'mul(UNITY_MATRIX_MVP,*)' with 'UnityObjectToClipPos(*)'

Shader "Hunter/3Dmiaobian" {
 Properties {
  _Color ("Main Color", Color) = (1,1,1,1)
  _OutlineColor ("Outline Color", Color) = (0,0,0,1) //改变这个能改变轮廓边的颜色
  _Outline ("Outline width",float) = 0.007 //改变这个能改变轮廓边的粗细
  _MainTex ("Base (RGB)", 2D) = "white" { }
  _Factor("Factor",float)=1
  _HitColor("HitColor",Color) = (0,0,0,1)
 }
 
 
 SubShader {
  Tags { "Queue" = "Transparent" }  //使用透明渲染队列，这个渲染队列在几何体队列之后被渲染，采用由后
      //到前的次序。任何采用alpha混合的对象（不对深度缓冲产生写操作的着色器）应该在这里渲染。
  // note that a vertex shader is specified here but its using the one above
  LOD 200  
  Pass {
   Name "OUTLINE"
   Tags { "LightMode" = "Always" }
   Cull Front//正面剔除（后期发现，如果不剔除正面，当多个相同模型的uv重合度较大时，非边缘区域出现边
                   //缘色区域）
   ZWrite Off//关闭深度缓冲，使得绘制的区域是一个2D平面图形
   ZTEST Less//显示绘制区域像素点时，如果该像素点已经被绘制，则不再绘制该像素点，实现遮挡。
   ColorMask RGB // alpha not used
   // you can choose what kind of blending mode you want for the outline
  // Blend SrcAlpha OneMinusSrcAlpha // Normal
   //Blend One One // Additive
   //Blend One OneMinusDstColor // Soft Additive
   //Blend DstColor Zero // Multiplicative
   //Blend DstColor SrcColor // 2x Multiplicative
  CGPROGRAM
  #pragma vertex vert
  #pragma fragment frag
  #include "UnityCG.cginc"
  struct appdata {
   float4 vertex : POSITION;
   float3 normal : NORMAL;
   float4 texcoord : TEXCOORD0;  
  };
  struct v2f {
   float4 pos : SV_POSITION;
   float4 color : COLOR;
   float4 tex : TEXCOORD0; 
  };
  uniform sampler2D _MainTex;
  uniform float _Outline;
  uniform float4 _OutlineColor;
  uniform float _Factor;
  uniform float4 _Color;
  v2f vert(appdata v) {
   // just make a copy of incoming vertex data but scaled according to normal direction
    v2f o;
    o.tex = v.texcoord;
    o.pos=UnityObjectToClipPos(v.vertex);
    float3 dir=normalize(v.vertex.xyz);
    float3 dir2=v.normal;
    float D=dot(dir,dir2);
    D=(D/_Factor+1)/(1+1/_Factor);
    dir=lerp(dir2,dir,D);
    dir= mul ((float3x3)UNITY_MATRIX_IT_MV, dir);
    float2 offset = TransformViewToProjection(dir.xy);
    offset=normalize(offset);
    //o.pos.xy += offset * o.pos.z *_Outline;//乘以o.pos.z导致物体里摄像机越远，描边越粗，不想要可以
    o.color = _OutlineColor;

    return o;
  }

  float4 frag(v2f i) :COLOR {
  		float4 textureColor = tex2D(_MainTex, i.tex.xy);  
        textureColor.rgb *= _Color.rgb;
   return textureColor;
  }
  ENDCG
  }
  Pass
  {
    NAME"BASE"
    CULL BACK    
             CGPROGRAM
             #pragma vertex vert
             #pragma fragment frag
             #include "UnityCG.cginc"
             struct vertOut {
                 float4 pos:SV_POSITION;
                 float4 tex: TEXCOORD0;
             };
    float _Amount;
             vertOut vert(appdata_base v) {
                 vertOut o;
                 o.pos = UnityObjectToClipPos (v.vertex);
                 o.tex = v.texcoord;
                 return o;
             }
    float4 _Color;
    sampler2D _MainTex;
             float4 frag(vertOut i) : COLOR0
             {
			              float4 textureColor = tex2D(_MainTex, i.tex.xy);  
			        textureColor.rgb *= _Color.rgb;
			   return textureColor;
             }
            ENDCG
  }
 } 
  
 Fallback "Diffuse"
}