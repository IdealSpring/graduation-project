package cn.ccut.ibarandomforest;

import org.apache.mahout.classifier.df.mapreduce.BuildForest;
import org.apache.mahout.classifier.df.tools.Describe;

/**
 * 1.训练随机森林模型
 *
 * @author 李程鸿
 */
public class BuildForestModel {
    /**
     * 需要传入参数：
     *      1.训练文件路径 -p /user/hadoop/runjarin/03/3-CV_train.dat
     *      2.描述文件路径 -f /user/hadoop/runjarin/03/train.info
     *      3.模型保存路径 -o /user/hadoop/runjarin/03/forest_result
     * @param args
     */
    public static void main(String[] args) throws Exception {
        String[] paramsDescribe =new String[]{
                "-p", args[0],
                "-f", args[1],
                "-d", "I", "2", "N", "16", "C", "L"
        };

        String[] paramsBuildForest = new String[]{
                "-Drapred.max.split.size=1874231",
                "-d", args[0],
                "-ds", args[1],
                "-o", args[2],
                "-sl", "5",
                "-p", "-t", "1000"
        };

        Describe.main(paramsDescribe);
        BuildForest.main(paramsBuildForest);
    }
}
