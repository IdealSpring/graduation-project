package cn.ccut.stage03;

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
 * 第三阶段MapReduce:
 *      提取关于hwmx的属性
 *
 * @author 李程鸿
 */
public class Stage03Main {
    private static final Logger log = LoggerFactory.getLogger(Stage03Main.class);

    public static void main(String[] args) throws Exception{
        Configuration conf = new Configuration();
        Job job = Job.getInstance(conf, "MyThirdJob");

        job.setJarByClass(Stage03Main.class);
        job.setMapperClass(DetailsExtractionMapper.class);
        job.setReducerClass(DetailsExtractionReducer.class);

        job.setMapOutputKeyClass(Text.class);
        job.setMapOutputValueClass(DetailsExtraction.class);

        job.setOutputKeyClass(DetailsExtraction.class);
        job.setOutputValueClass(NullWritable.class);

        // 输入文件路径
        FileInputFormat.addInputPaths(job, args[0]);

        // 输出文件路径
        FileOutputFormat.setOutputPath(job, new Path(args[1]));

        job.waitForCompletion(true);
        log.info("Stage03Main(hwmx表属性提取完毕) Successful......");
    }
}
