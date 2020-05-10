package org.apache.mahout.classifier.df.mapreduce;

import org.apache.commons.cli2.CommandLine;
import org.apache.commons.cli2.Group;
import org.apache.commons.cli2.Option;
import org.apache.commons.cli2.OptionException;
import org.apache.commons.cli2.builder.ArgumentBuilder;
import org.apache.commons.cli2.builder.DefaultOptionBuilder;
import org.apache.commons.cli2.builder.GroupBuilder;
import org.apache.commons.cli2.commandline.Parser;
import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.conf.Configured;
import org.apache.hadoop.fs.FSDataInputStream;
import org.apache.hadoop.fs.FSDataOutputStream;
import org.apache.hadoop.fs.FileSystem;
import org.apache.hadoop.fs.Path;
import org.apache.hadoop.util.Tool;
import org.apache.hadoop.util.ToolRunner;
import org.apache.mahout.classifier.ClassifierResult;
import org.apache.mahout.classifier.RegressionResultAnalyzer;
import org.apache.mahout.classifier.ResultAnalyzer;
import org.apache.mahout.classifier.df.DFUtils;
import org.apache.mahout.classifier.df.DecisionForest;
import org.apache.mahout.classifier.df.data.DataConverter;
import org.apache.mahout.classifier.df.data.Dataset;
import org.apache.mahout.classifier.df.data.Instance;
import org.apache.mahout.common.CommandLineUtil;
import org.apache.mahout.common.RandomUtils;
import org.apache.mahout.common.commandline.DefaultOptionCreator;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.*;
import java.util.*;

/**
 * 使用先前建立的决策森林对数据集进行分类的工具
 *
 * @author 李程鸿
 */
public class TestForest extends Configured implements Tool {

    private static final Logger log = LoggerFactory.getLogger(TestForest.class);

    private FileSystem dataFS;
    private Path dataPath; // 测试数据路径

    private Path datasetPath;

    private Path modelPath; // 森林储存的路径

    private FileSystem outFS;
    private Path outputPath; // 预测文件的路径，如果NULL不输出预测

    private boolean analyze;

    private boolean useMapreduce;

    @Override
    public int run(String[] args) throws IOException, ClassNotFoundException, InterruptedException {

        DefaultOptionBuilder obuilder = new DefaultOptionBuilder();
        ArgumentBuilder abuilder = new ArgumentBuilder();
        GroupBuilder gbuilder = new GroupBuilder();

        Option inputOpt = DefaultOptionCreator.inputOption().create();

        Option datasetOpt = obuilder.withLongName("dataset").withShortName("ds").withRequired(true).withArgument(
                abuilder.withName("dataset").withMinimum(1).withMaximum(1).create()).withDescription("Dataset path")
                .create();

        Option modelOpt = obuilder.withLongName("model").withShortName("m").withRequired(true).withArgument(
                abuilder.withName("path").withMinimum(1).withMaximum(1).create()).
                withDescription("Path to the Decision Forest").create();

        Option outputOpt = DefaultOptionCreator.outputOption().create();

        Option analyzeOpt = obuilder.withLongName("analyze").withShortName("a").withRequired(false).create();

        Option mrOpt = obuilder.withLongName("mapreduce").withShortName("mr").withRequired(false).create();

        Option helpOpt = DefaultOptionCreator.helpOption();

        Group group = gbuilder.withName("Options").withOption(inputOpt).withOption(datasetOpt).withOption(modelOpt)
                .withOption(outputOpt).withOption(analyzeOpt).withOption(mrOpt).withOption(helpOpt).create();

        try {
            Parser parser = new Parser();
            parser.setGroup(group);
            CommandLine cmdLine = parser.parse(args);

            if (cmdLine.hasOption("help")) {
                CommandLineUtil.printHelp(group);
                return -1;
            }

            String dataName = cmdLine.getValue(inputOpt).toString();
            String datasetName = cmdLine.getValue(datasetOpt).toString();
            String modelName = cmdLine.getValue(modelOpt).toString();
            String outputName = cmdLine.hasOption(outputOpt) ? cmdLine.getValue(outputOpt).toString() : null;
            analyze = cmdLine.hasOption(analyzeOpt);
            useMapreduce = cmdLine.hasOption(mrOpt);

            if (log.isDebugEnabled()) {
                log.debug("inout     : {}", dataName);
                log.debug("dataset   : {}", datasetName);
                log.debug("model     : {}", modelName);
                log.debug("output    : {}", outputName);
                log.debug("analyze   : {}", analyze);
                log.debug("mapreduce : {}", useMapreduce);
            }

            dataPath = new Path(dataName);
            datasetPath = new Path(datasetName);
            modelPath = new Path(modelName);
            if (outputName != null) {
                outputPath = new Path(outputName);
            }
        } catch (OptionException e) {
            log.warn(e.toString(), e);
            CommandLineUtil.printHelp(group);
            return -1;
        }

        testForest();

        return 0;
    }

    /**
     * 测试随机森林
     *
     * @throws IOException
     * @throws ClassNotFoundException
     * @throws InterruptedException
     */
    private void testForest() throws IOException, ClassNotFoundException, InterruptedException {

        // 确保输出文件不存在
        if (outputPath != null) {
            outFS = outputPath.getFileSystem(getConf());
            if (outFS.exists(outputPath)) {
                throw new IllegalArgumentException("Output path already exists");
            }
        }

        // 确保决策森林存在
        FileSystem mfs = modelPath.getFileSystem(getConf());
        if (!mfs.exists(modelPath)) {
            throw new IllegalArgumentException("The forest path does not exist");
        }

        // 确保测试数据存在
        dataFS = dataPath.getFileSystem(getConf());
        if (!dataFS.exists(dataPath)) {
            throw new IllegalArgumentException("The Test data path does not exist");
        }

        if (useMapreduce) {
            mapreduce();
        } else {
            sequential();
        }

    }

    /**
     * MapReduce部分
     *
     * @throws ClassNotFoundException
     * @throws IOException
     * @throws InterruptedException
     */
    private void mapreduce() throws ClassNotFoundException, IOException, InterruptedException {
        if (outputPath == null) {
            throw new IllegalArgumentException("You must specify the ouputPath when using the mapreduce implementation");
        }

        Classifier classifier = new Classifier(modelPath, dataPath, datasetPath, outputPath, getConf());

        classifier.run();

        if (analyze) {
            double[][] results = classifier.getResults();
            if (results != null) {
                Dataset dataset = Dataset.load(getConf(), datasetPath);
                if (dataset.isNumerical(dataset.getLabelId())) {
                    RegressionResultAnalyzer regressionAnalyzer = new RegressionResultAnalyzer();
                    regressionAnalyzer.setInstances(results);
                    log.info("{}", regressionAnalyzer);
                } else {
                    ResultAnalyzer analyzer = new ResultAnalyzer(Arrays.asList(dataset.labels()), "unknown");
                    for (double[] res : results) {
                        analyzer.addInstance(dataset.getLabelString(res[0]),
                                new ClassifierResult(dataset.getLabelString(res[1]), 1.0));
                    }
                    log.info("{}", analyzer);

                }
            }
        }
    }

    /**
     * 按次序
     *
     * @throws IOException
     */
    private void sequential() throws IOException {

        log.info("Loading the forest...");
        DecisionForest forest = DecisionForest.load(getConf(), modelPath);

        if (forest == null) {
            log.error("No Decision Forest found!");
            return;
        }

        // 加载数据集
        Dataset dataset = Dataset.load(getConf(), datasetPath);
        DataConverter converter = new DataConverter(dataset);

        log.info("Sequential classification...");
        long time = System.currentTimeMillis();

        Random rng = RandomUtils.getRandom();

        List<double[]> resList = new ArrayList<>();
        if (dataFS.getFileStatus(dataPath).isDir()) {
            // 输入是一个文件目录
            testDirectory(outputPath, converter, forest, dataset, resList, rng);
        }  else {
            // 输入是一个文件
            testFile(dataPath, outputPath, converter, forest, dataset, resList, rng);
        }

        time = System.currentTimeMillis() - time;
        log.info("Classification Time: {}", DFUtils.elapsedTime(time));

        if (analyze) {
            if (dataset.isNumerical(dataset.getLabelId())) {
                RegressionResultAnalyzer regressionAnalyzer = new RegressionResultAnalyzer();
                double[][] results = new double[resList.size()][2];
                regressionAnalyzer.setInstances(resList.toArray(results));
                log.info("{}", regressionAnalyzer);
            } else {
                ResultAnalyzer analyzer = new ResultAnalyzer(Arrays.asList(dataset.labels()), "unknown");
                for (double[] r : resList) {
                    analyzer.addInstance(dataset.getLabelString(r[0]),
                            new ClassifierResult(dataset.getLabelString(r[1]), 1.0));
                }
                log.info("{}", analyzer);

                //将fiscore写入文档
                String weightedF1score = analyzer.getWeightedF1score();
            }
        }
    }

    /**
     * 测试目录
     *
     * @param outPath
     *          输出路径
     * @param converter
     *          转换器
     * @param forest
     *          随机森林
     * @param dataset
     *          数据集
     * @param results
     *          结果
     * @param rng
     * @throws IOException
     */
    private void testDirectory(Path outPath,
                               DataConverter converter,
                               DecisionForest forest,
                               Dataset dataset,
                               Collection<double[]> results,
                               Random rng) throws IOException {
        Path[] infiles = DFUtils.listOutputFiles(dataFS, dataPath);

        for (Path path : infiles) {
            log.info("Classifying : {}", path);
            Path outfile = outPath != null ? new Path(outPath, path.getName()).suffix(".out") : null;
            testFile(path, outfile, converter, forest, dataset, results, rng);
        }
    }

    /**
     * 测试文件
     *
     * @param inPath
     *          输入路径
     * @param outPath
     *          输出路径
     * @param converter
     *          转换器
     * @param forest
     *          随机森林
     * @param dataset
     *          数据集
     * @param results
     *          结果
     * @param rng
     * @throws IOException
     */
    private void testFile(Path inPath,
                          Path outPath,
                          DataConverter converter,
                          DecisionForest forest,
                          Dataset dataset,
                          Collection<double[]> results,
                          Random rng) throws IOException {
        // 创建预测文件
        FSDataOutputStream ofile = null;

        if (outPath != null) {
            ofile = outFS.create(outPath);
        }

        try (FSDataInputStream input = dataFS.open(inPath)){
            Scanner scanner = new Scanner(input, "UTF-8");

            while (scanner.hasNextLine()) {
                String line = scanner.nextLine();
                if (!line.isEmpty()) {

                    Instance instance = converter.convert(line);
                    double prediction = forest.classify(dataset, rng, instance);

                    int predictionInt = 0;
                    if(prediction == 1.0) {
                        predictionInt = 1;
                    } else {
                        predictionInt = 0;
                    }

                    String nsr_id = line.substring(0, line.indexOf(","));

                    if (ofile != null) {
                        ofile.writeBytes(nsr_id + "\t" + Integer.toString(predictionInt) + "\n");
                    }
                    results.add(new double[]{dataset.getLabel(instance), prediction});
                }
            }
            scanner.close();
        }
    }

    public static void main(String[] args) throws Exception {
        ToolRunner.run(new Configuration(), new TestForest(), args);
    }

}
