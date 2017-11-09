#if UNITY_EDITOR && UNITY_IOS
using System.IO;
using UnityEngine;
using UnityEditor;
using UnityEditor.iOS.Xcode;
using UnityEditor.Callbacks;
using UnityEditor.XCodeEditor;

using System.Collections;

public class XCodeSettings
{
    //该属性是在build完成后，被调用的callback
    [PostProcessBuildAttribute(0)]
    public static void OnPostprocessBuild(BuildTarget buildTarget, string pathToBuiltProject)
    {
        // BuildTarget需为iOS
        if (buildTarget != BuildTarget.iOS)
            return;

        // 初始化
        var projectPath = pathToBuiltProject + "/Unity-iPhone.xcodeproj/project.pbxproj";
        PBXProject pbxProject = new PBXProject();
        pbxProject.ReadFromFile(projectPath);
        string targetGuid = pbxProject.TargetGuidByName("Unity-iPhone");

        // 添加flag
        pbxProject.AddBuildProperty(targetGuid, "OTHER_LDFLAGS", "-ObjC");
        // 关闭Bitcode
        pbxProject.SetBuildProperty(targetGuid, "ENABLE_BITCODE", "NO");

        // 添加framwrok
        pbxProject.AddFrameworkToProject(targetGuid, "Security.framework", false);
//        pbxProject.AddFrameworkToProject(targetGuid, "CoreTelephony.framework", false);
//        pbxProject.AddFrameworkToProject(targetGuid, "SystemConfiguration.framework", false);
//        pbxProject.AddFrameworkToProject(targetGuid, "CoreGraphics.framework", false);
//        pbxProject.AddFrameworkToProject(targetGuid, "ImageIO.framework", false);
//        pbxProject.AddFrameworkToProject(targetGuid, "CoreData.framework", false);

        //添加lib
//        AddLibToProject(pbxProject, targetGuid, "libYvImSdk.a");
//        AddLibToProject(pbxProject, targetGuid, "libc++.1.dylib");
//        AddLibToProject(pbxProject, targetGuid, "libz.1.dylib");

        // 应用修改
        File.WriteAllText(projectPath, pbxProject.WriteToString());

        // 修改Info.plist文件
//        var plistPath = Path.Combine(pathToBuiltProject, "Info.plist");
//        var plist = new PlistDocument();
//        plist.ReadFromFile(plistPath);

        // 插入URL Scheme到Info.plsit（理清结构）
//        var array = plist.root.CreateArray("CFBundleURLTypes");
        //插入dict
//        var urlDict = array.AddDict();
//        urlDict.SetString("CFBundleTypeRole", "Editor");
        //插入array
//        var urlInnerArray = urlDict.CreateArray("CFBundleURLSchemes");
//        urlInnerArray.AddString("blablabla");
        // 应用修改
//        plist.WriteToFile(plistPath);

        //插入代码
        //读取UnityAppController.mm文件
//        string unityAppControllerPath = pathToBuiltProject + "/Classes/UnityAppController.mm";
//        XClass UnityAppController = new XClass(unityAppControllerPath);

        //在指定代码后面增加一行代码
//        UnityAppController.WriteBelow("#include \"PluginBase/AppDelegateListener.h\"", "#import <UMSocialCore/UMSocialCore.h>");
//
//        string newCode = "\n" +
//            "    [[UMSocialManager defaultManager] openLog:YES];\n" +
//            "    [UMSocialGlobal shareInstance].type = @\"u3d\";\n" +
//            "    [[UMSocialManager defaultManager] setUmSocialAppkey:@\"" +"\"];\n" +
//            "    [[UMSocialManager defaultManager] setPlaform:UMSocialPlatformType_WechatSession appKey:@\""+"\" appSecret:@\""+ "\" redirectURL:@\"http://mobile.umeng.com/social\"];\n" +
//            "    \n"
//            ;
        //在指定代码后面增加一大行代码
//        UnityAppController.WriteBelow("// if you wont use keyboard you may comment it out at save some memory", newCode);

    }

    //添加lib方法
    static void AddLibToProject(PBXProject inst, string targetGuid, string lib)
    {
        string fileGuid = inst.AddFile("usr/lib/" + lib, "Frameworks/" + lib, PBXSourceTree.Sdk);
        inst.AddFileToBuild(targetGuid, fileGuid);
    }
}
#endif