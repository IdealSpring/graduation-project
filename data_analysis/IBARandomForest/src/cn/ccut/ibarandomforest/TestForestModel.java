package cn.ccut.ibarandomforest;

import org.apache.mahout.classifier.df.mapreduce.TestForest;

/**
 * 测试模型
 *
 * @author 李程鸿
 */
public class TestForestModel {
    /**
     * 需要传入参数：cn.ccut.ibarandomforest.TestForestModel
     *      1.测试文件路径
     *      2.描述文件路径
     *      3.预测文件保存
     *      4.模型保存路径
     * @param args
     */
    public static void main(String[] args) throws Exception {
        String[] paramsTestForest = null;

        if(args.length == 5) {
            if(args[4].equals("a")) {
                paramsTestForest = new String[]{
                        "-i", args[0],
                        "-ds", args[1],
                        "-o", args[2],
                        "-m", args[3],
                        "-a"
                };
            } else {
                throw new Exception("参数异常");
            }
        } else {
            paramsTestForest = new String[]{
                    "-i", args[0],
                    "-ds", args[1],
                    "-o", args[2],
                    "-m", args[3]
            };
        }

        TestForest.main(paramsTestForest);
    }
}
