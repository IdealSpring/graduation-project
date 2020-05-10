package cn.ccut.stage01;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.Path;
import org.apache.hadoop.io.NullWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Job;
import org.apache.hadoop.mapreduce.lib.input.FileInputFormat;
import org.apache.hadoop.mapreduce.lib.output.FileOutputFormat;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * 第一阶段MapReduce主类
 *      提取前18个属性
 *
 * @author 李程鸿
 */
public class Stage01Main {
    private static final Logger log = LoggerFactory.getLogger(Stage01Main.class);

    public static void main(String[] args) throws Exception{
        Configuration conf = new Configuration();
        //向服务器提交执行
        Job job = Job.getInstance(conf, "MyFirstJob");

        job.setJarByClass(Stage01Main.class);
        job.setMapperClass(EnterpriesMapper.class);
        job.setReducerClass(EnterpriesReducer.class);

        job.setMapOutputKeyClass(Text.class);
        job.setMapOutputValueClass(Enterprise.class);

        job.setOutputKeyClass(Enterprise.class);
        job.setOutputValueClass(NullWritable.class);

        // 输入文件路径
        FileInputFormat.addInputPaths(job, args[0]);

        // 输出文件路径
        FileOutputFormat.setOutputPath(job, new Path(args[1]));

        job.waitForCompletion(true);
        log.info("Stage01Main Successful......");
    }
}
