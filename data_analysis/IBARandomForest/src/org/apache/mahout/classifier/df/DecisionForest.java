package org.apache.mahout.classifier.df;

import com.google.common.base.Preconditions;
import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.FSDataInputStream;
import org.apache.hadoop.fs.FileSystem;
import org.apache.hadoop.fs.Path;
import org.apache.hadoop.io.Writable;
import org.apache.mahout.classifier.df.data.Data;
import org.apache.mahout.classifier.df.data.DataUtils;
import org.apache.mahout.classifier.df.data.Dataset;
import org.apache.mahout.classifier.df.data.Instance;
import org.apache.mahout.classifier.df.node.Node;

import java.io.DataInput;
import java.io.DataOutput;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Random;

/**
 * 表示决策树的森林
 *
 * @author 李程鸿
 */
public class DecisionForest implements Writable {

    private final List<Node> trees;
    private int flog = 1;
    private int index = 0;

    private DecisionForest() {
        trees = new ArrayList<>();
    }

    public DecisionForest(List<Node> trees) {
        Preconditions.checkArgument(trees != null && !trees.isEmpty(), "trees argument must not be null or empty");

        this.trees = trees;
    }

    public List<Node> getTrees() {
        return trees;
    }

    /**
     * 对数据进行分类并调用每个分类的回调
     */
    public void classify(Data data, double[][] predictions) {
        Preconditions.checkArgument(data.size() == predictions.length, "predictions.length must be equal to data.size()");

        if (data.isEmpty()) {
            return; // nothing to classify
        }

        int treeId = 0;
        for (Node tree : trees) {
            for (int index = 0; index < data.size(); index++) {
                if (predictions[index] == null) {
                    predictions[index] = new double[trees.size()];
                }
                predictions[index][treeId] = tree.classify(data.get(index));
            }
            treeId++;
        }
    }

    /**
     * 预测实例的标签
     *
     * @param rng
     *          随机数发生器
     * @return NaN
     *          如果标签不能预测
     */
    public double classify(Dataset dataset, Random rng, Instance instance) {
        if (dataset.isNumerical(dataset.getLabelId())) {
            double sum = 0;
            int cnt = 0;
            for (Node tree : trees) {
                double prediction = tree.classify(instance);
                if (!Double.isNaN(prediction)) {
                    sum += prediction;
                    cnt++;
                }
            }

            if (cnt > 0) {
                return sum / cnt;
            } else {
                return Double.NaN;
            }
        } else {
            int[] predictions = new int[dataset.nblabels()];

            for (Node tree : trees) {
                double prediction = tree.classify(instance);

                if (!Double.isNaN(prediction)) {
                    predictions[(int) prediction]++;
                }
            }

            if (DataUtils.sum(predictions) == 0) {
                return Double.NaN; // no prediction available
            }

            return DataUtils.maxindex(rng, predictions);
        }
    }

    /**
     * @return 平均树数
     */
    public long meanNbNodes() {
        long sum = 0;

        for (Node tree : trees) {
            sum += tree.nbNodes();
        }

        return sum / trees.size();
    }

    /**
     * @return 所有树中的节点总数
     */
    public long nbNodes() {
        long sum = 0;

        for (Node tree : trees) {
            sum += tree.nbNodes();
        }

        return sum;
    }

    /**
     * @return 单株平均最大深度
     */
    public long meanMaxDepth() {
        long sum = 0;

        for (Node tree : trees) {
            sum += tree.maxDepth();
        }

        return sum / trees.size();
    }

    @Override
    public boolean equals(Object obj) {
        if (this == obj) {
            return true;
        }
        if (!(obj instanceof DecisionForest)) {
            return false;
        }

        DecisionForest rf = (DecisionForest) obj;

        return trees.size() == rf.getTrees().size() && trees.containsAll(rf.getTrees());
    }

    @Override
    public int hashCode() {
        return trees.hashCode();
    }

    @Override
    public void write(DataOutput dataOutput) throws IOException {
        dataOutput.writeInt(trees.size());
        for (Node tree : trees) {
            tree.write(dataOutput);
        }
    }

    /**
     * 从输入中读取树，并将它们添加到现有的树中。
     */
    @Override
    public void readFields(DataInput dataInput) throws IOException {
        int size = dataInput.readInt();
        for (int i = 0; i < size; i++) {
            trees.add(Node.read(dataInput));
        }
    }

    /**
     * 从输入流读取森林
	 *
     * @param dataInput - input forest
     * @return {@link org.apache.mahout.classifier.df.DecisionForest}
     * @throws IOException
     */
    public static DecisionForest read(DataInput dataInput) throws IOException {
        DecisionForest forest = new DecisionForest();
        forest.readFields(dataInput);
        return forest;
    }

    /**
     * 从单个文件或文件目录加载随机森林
	 *
     * @throws IOException
     */
    public static DecisionForest load(Configuration conf, Path forestPath) throws IOException {
        FileSystem fs = forestPath.getFileSystem(conf);
        Path[] files;
        if (fs.getFileStatus(forestPath).isDir()) {
            files = DFUtils.listOutputFiles(fs, forestPath);
        } else {
            files = new Path[]{forestPath};
        }

        DecisionForest forest = null;
        for (Path path : files) {
            try (FSDataInputStream dataInput = new FSDataInputStream(fs.open(path))) {
                if (forest == null) {
                    forest = read(dataInput);
                } else {
                    forest.readFields(dataInput);
                }
            }
        }

        return forest;

    }

}
