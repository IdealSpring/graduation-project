package org.apache.mahout.classifier;

import com.google.common.base.Preconditions;
import java.util.Collection;
import java.util.Collections;
import java.util.HashMap;
import java.util.Iterator;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.Map.Entry;
import org.apache.commons.lang3.StringUtils;
import org.apache.commons.math3.stat.descriptive.moment.Mean;
import org.apache.mahout.cf.taste.impl.common.FullRunningAverageAndStdDev;
import org.apache.mahout.cf.taste.impl.common.RunningAverageAndStdDev;
import org.apache.mahout.math.DenseMatrix;
import org.apache.mahout.math.Matrix;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * 计算各类评测指标
 *
 * @author 李程鸿
 */
public class ConfusionMatrix {
    private static final Logger LOG = LoggerFactory.getLogger(ConfusionMatrix.class);
    private final Map<String, Integer> labelMap = new LinkedHashMap();
    private final int[][] confusionMatrix;
    private int samples = 0;
    private String defaultLabel = "unknown";

    /**
     * 混淆矩阵
     *
     * @param labels
     *          标签
     * @param defaultLabel
     *          默认标签
     */
    public ConfusionMatrix(Collection<String> labels, String defaultLabel) {
        this.confusionMatrix = new int[labels.size() + 1][labels.size() + 1];
        this.defaultLabel = defaultLabel;
        int i = 0;
        Iterator i$ = labels.iterator();

        while(i$.hasNext()) {
            String label = (String)i$.next();
            this.labelMap.put(label, i++);
        }

        this.labelMap.put(defaultLabel, i);
    }

    /**
     * 混淆矩阵
     *
     * @param m
     */
    public ConfusionMatrix(Matrix m) {
        this.confusionMatrix = new int[m.numRows()][m.numRows()];
        this.setMatrix(m);
    }

    public int[][] getConfusionMatrix() {
        return this.confusionMatrix;
    }

    public Collection<String> getLabels() {
        return Collections.unmodifiableCollection(this.labelMap.keySet());
    }

    private int numLabels() {
        return this.labelMap.size();
    }

    /**
     * 计算准确率
     *
     * @param label
     *          标签
     * @return
     */
    public double getAccuracy(String label) {
        int labelId = (Integer)this.labelMap.get(label);
        int labelTotal = 0;
        int correct = 0;

        for(int i = 0; i < this.numLabels(); ++i) {
            labelTotal += this.confusionMatrix[labelId][i];
            if (i == labelId) {
                correct += this.confusionMatrix[labelId][i];
            }
        }

        return 100.0D * (double)correct / (double)labelTotal;
    }

    /**
     * 计算准确率
     *
     * @return
     */
    public double getAccuracy() {
        int total = 0;
        int correct = 0;

        for(int i = 0; i < this.numLabels(); ++i) {
            for(int j = 0; j < this.numLabels(); ++j) {
                total += this.confusionMatrix[i][j];
                if (i == j) {
                    correct += this.confusionMatrix[i][j];
                }
            }
        }

        return 100.0D * (double)correct / (double)total;
    }

    /**
     * 获取类的测试实例数
     *
     * @param label
     *          标签
     * @return
     */
    private int getActualNumberOfTestExamplesForClass(String label) {
        int labelId = (Integer)this.labelMap.get(label);
        int sum = 0;

        for(int i = 0; i < this.numLabels(); ++i) {
            sum += this.confusionMatrix[labelId][i];
        }

        return sum;
    }

    /**
     * 计算精度
     *
     * @param label
     * @return
     */
    public double getPrecision(String label) {
        int labelId = (Integer)this.labelMap.get(label);
        int truePositives = this.confusionMatrix[labelId][labelId];
        int falsePositives = 0;

        for(int i = 0; i < this.numLabels(); ++i) {
            if (i != labelId) {
                falsePositives += this.confusionMatrix[i][labelId];
            }
        }

        if (truePositives + falsePositives == 0) {
            return 0.0D;
        } else {
            return (double)truePositives / (double)(truePositives + falsePositives);
        }
    }

    /**
     * 计算加权精度
     *
     * @return
     */
    public double getWeightedPrecision() {
        double[] precisions = new double[this.numLabels()];
        double[] weights = new double[this.numLabels()];
        int index = 0;

        for(Iterator i$ = this.labelMap.keySet().iterator(); i$.hasNext(); ++index) {
            String label = (String)i$.next();
            precisions[index] = this.getPrecision(label);
            weights[index] = (double)this.getActualNumberOfTestExamplesForClass(label);
        }

        return (new Mean()).evaluate(precisions, weights);
    }

    /**
     * 计算召回率
     *
     * @param label
     *          标签
     * @return
     */
    public double getRecall(String label) {
        int labelId = (Integer)this.labelMap.get(label);
        int truePositives = this.confusionMatrix[labelId][labelId];
        int falseNegatives = 0;

        for(int i = 0; i < this.numLabels(); ++i) {
            if (i != labelId) {
                falseNegatives += this.confusionMatrix[labelId][i];
            }
        }

        if (truePositives + falseNegatives == 0) {
            return 0.0D;
        } else {
            return (double)truePositives / (double)(truePositives + falseNegatives);
        }
    }

    /**
     * 计算加权召回率
     *
     * @return
     */
    public double getWeightedRecall() {
        double[] recalls = new double[this.numLabels()];
        double[] weights = new double[this.numLabels()];
        int index = 0;

        for(Iterator i$ = this.labelMap.keySet().iterator(); i$.hasNext(); ++index) {
            String label = (String)i$.next();
            recalls[index] = this.getRecall(label);
            weights[index] = (double)this.getActualNumberOfTestExamplesForClass(label);
        }

        return (new Mean()).evaluate(recalls, weights);
    }

    /**
     * 计算F1-score
     *
     * @param label
     * @return
     */
    public double getF1score(String label) {
        double precision = this.getPrecision(label);
        double recall = this.getRecall(label);
        return precision + recall == 0.0D ? 0.0D : 2.0D * precision * recall / (precision + recall);
    }

    /**
     * 计算加权F1-score
     *
     * @return
     */
    public double getWeightedF1score() {
        double[] f1Scores = new double[this.numLabels()];
        double[] weights = new double[this.numLabels()];
        int index = 0;

        for(Iterator i$ = this.labelMap.keySet().iterator(); i$.hasNext(); ++index) {
            String label = (String)i$.next();
            f1Scores[index] = this.getF1score(label);
            weights[index] = (double)this.getActualNumberOfTestExamplesForClass(label);
        }

        return (new Mean()).evaluate(f1Scores, weights);
    }

    /**
     * 计算可靠性
     *
     * @return
     */
    public double getReliability() {
        int count = 0;
        double accuracy = 0.0D;

        for(Iterator i$ = this.labelMap.keySet().iterator(); i$.hasNext(); ++count) {
            String label = (String)i$.next();
            if (!label.equals(this.defaultLabel)) {
                accuracy += this.getAccuracy(label);
            }
        }

        return accuracy / (double)count;
    }

    /**
     * 计算kappa系数
     *
     * @return
     */
    public double getKappa() {
        double a = 0.0D;
        double b = 0.0D;

        for(int i = 0; i < this.confusionMatrix.length; ++i) {
            a += (double)this.confusionMatrix[i][i];
            double br = 0.0D;

            for(int j = 0; j < this.confusionMatrix.length; ++j) {
                br += (double)this.confusionMatrix[i][j];
            }

            double bc = 0.0D;
            int[][] arr$ = this.confusionMatrix;
            int len$ = arr$.length;

            for(int i$ = 0; i$ < len$; ++i$) {
                int[] vec = arr$[i$];
                bc += (double)vec[i];
            }

            b += br * bc;
        }

        return ((double)this.samples * a - b) / ((double)(this.samples * this.samples) - b);
    }

    public RunningAverageAndStdDev getNormalizedStats() {
        RunningAverageAndStdDev summer = new FullRunningAverageAndStdDev();

        for(int d = 0; d < this.confusionMatrix.length; ++d) {
            double total = 0.0D;

            for(int j = 0; j < this.confusionMatrix.length; ++j) {
                total += (double)this.confusionMatrix[d][j];
            }

            summer.addDatum((double)this.confusionMatrix[d][d] / (total + 1.0E-6D));
        }

        return summer;
    }

    /**
     * 得到正常的id
     *
     * @param label
     *      标签
     * @return
     */
    public int getCorrect(String label) {
        int labelId = (Integer)this.labelMap.get(label);
        return this.confusionMatrix[labelId][labelId];
    }

    /**
     * 得到总数
     *
     * @param label
     *          标签
     * @return
     */
    public int getTotal(String label) {
        int labelId = (Integer)this.labelMap.get(label);
        int labelTotal = 0;

        for(int i = 0; i < this.labelMap.size(); ++i) {
            labelTotal += this.confusionMatrix[labelId][i];
        }

        return labelTotal;
    }

    /**
     * 增加实例
     *
     * @param correctLabel
     * @param classifiedResult
     */
    public void addInstance(String correctLabel, ClassifierResult classifiedResult) {
        ++this.samples;
        this.incrementCount(correctLabel, classifiedResult.getLabel());
    }

    /**
     * 增加实例
     *
     * @param correctLabel
     * @param classifiedLabel
     */
    public void addInstance(String correctLabel, String classifiedLabel) {
        ++this.samples;
        this.incrementCount(correctLabel, classifiedLabel);
    }

    /**
     * 得到总数
     *
     * @param correctLabel
     * @param classifiedLabel
     * @return
     */
    public int getCount(String correctLabel, String classifiedLabel) {
        if (!this.labelMap.containsKey(correctLabel)) {
            LOG.warn("Label {} did not appear in the training examples", correctLabel);
            return 0;
        } else {
            Preconditions.checkArgument(this.labelMap.containsKey(classifiedLabel), "Label not found: " + classifiedLabel);
            int correctId = (Integer)this.labelMap.get(correctLabel);
            int classifiedId = (Integer)this.labelMap.get(classifiedLabel);
            return this.confusionMatrix[correctId][classifiedId];
        }
    }

    public void putCount(String correctLabel, String classifiedLabel, int count) {
        if (!this.labelMap.containsKey(correctLabel)) {
            LOG.warn("Label {} did not appear in the training examples", correctLabel);
        } else {
            Preconditions.checkArgument(this.labelMap.containsKey(classifiedLabel), "Label not found: " + classifiedLabel);
            int correctId = (Integer)this.labelMap.get(correctLabel);
            int classifiedId = (Integer)this.labelMap.get(classifiedLabel);
            if ((double)this.confusionMatrix[correctId][classifiedId] == 0.0D && count != 0) {
                ++this.samples;
            }

            this.confusionMatrix[correctId][classifiedId] = count;
        }
    }

    /**
     * 得到默认标签
     *
     * @return
     */
    public String getDefaultLabel() {
        return this.defaultLabel;
    }

    public void incrementCount(String correctLabel, String classifiedLabel, int count) {
        this.putCount(correctLabel, classifiedLabel, count + this.getCount(correctLabel, classifiedLabel));
    }

    public void incrementCount(String correctLabel, String classifiedLabel) {
        this.incrementCount(correctLabel, classifiedLabel, 1);
    }

    public ConfusionMatrix merge(ConfusionMatrix b) {
        Preconditions.checkArgument(this.labelMap.size() == b.getLabels().size(), "The label sizes do not match");
        Iterator i$ = this.labelMap.keySet().iterator();

        while(i$.hasNext()) {
            String correctLabel = (String)i$.next();
            i$ = this.labelMap.keySet().iterator();

            while(i$.hasNext()) {
                String classifiedLabel = (String)i$.next();
                this.incrementCount(correctLabel, classifiedLabel, b.getCount(correctLabel, classifiedLabel));
            }
        }

        return this;
    }

    /**
     * 得到矩阵
     *
     * @return
     */
    public Matrix getMatrix() {
        int length = this.confusionMatrix.length;
        Matrix m = new DenseMatrix(length, length);

        for(int r = 0; r < length; ++r) {
            for(int c = 0; c < length; ++c) {
                m.set(r, c, (double)this.confusionMatrix[r][c]);
            }
        }

        Map<String, Integer> labels = new HashMap();
        Iterator i$ = this.labelMap.entrySet().iterator();

        while(i$.hasNext()) {
            Entry<String, Integer> entry = (Entry)i$.next();
            labels.put(entry.getKey(), entry.getValue());
        }

        m.setRowLabelBindings(labels);
        m.setColumnLabelBindings(labels);
        return m;
    }

    /**
     * 设置矩阵
     *
     * @param m
     */
    public void setMatrix(Matrix m) {
        int length = this.confusionMatrix.length;
        if (m.numRows() != m.numCols()) {
            throw new IllegalArgumentException("ConfusionMatrix: matrix(" + m.numRows() + ',' + m.numCols() + ") must be square");
        } else {
            for(int r = 0; r < length; ++r) {
                for(int c = 0; c < length; ++c) {
                    this.confusionMatrix[r][c] = (int)Math.round(m.get(r, c));
                }
            }

            Map<String, Integer> labels = m.getRowLabelBindings();
            if (labels == null) {
                labels = m.getColumnLabelBindings();
            }

            if (labels != null) {
                String[] sorted = sortLabels(labels);
                verifyLabels(length, sorted);
                this.labelMap.clear();

                for(int i = 0; i < length; ++i) {
                    this.labelMap.put(sorted[i], i);
                }
            }

        }
    }

    /**
     * 把标签分类
     *
     * @param labels
     *          标签
     * @return
     */
    private static String[] sortLabels(Map<String, Integer> labels) {
        String[] sorted = new String[labels.size()];

        Entry entry;
        for(Iterator i$ = labels.entrySet().iterator(); i$.hasNext(); sorted[(Integer)entry.getValue()] = (String)entry.getKey()) {
            entry = (Entry)i$.next();
        }

        return sorted;
    }

    /**
     * 验证标签
     *
     * @param length
     * @param sorted
     */
    private static void verifyLabels(int length, String[] sorted) {
        Preconditions.checkArgument(sorted.length == length, "One label, one row");

        for(int i = 0; i < length; ++i) {
            if (sorted[i] == null) {
                Preconditions.checkArgument(false, "One label, one row");
            }
        }

    }

    public String toString() {
        StringBuilder returnString = new StringBuilder(200);
        returnString.append("=======================================================").append('\n');
        returnString.append("Confusion Matrix\n");
        returnString.append("-------------------------------------------------------").append('\n');
        int unclassified = this.getTotal(this.defaultLabel);
        Iterator i$ = this.labelMap.entrySet().iterator();

        while(true) {
            Entry entry;
            do {
                if (!i$.hasNext()) {
                    returnString.append("<--Classified as").append('\n');
                    i$ = this.labelMap.entrySet().iterator();

                    label47:
                    while(true) {
                        do {
                            if (!i$.hasNext()) {
                                if (unclassified > 0) {
                                    returnString.append("Default Category: ").append(this.defaultLabel).append(": ").append(unclassified).append('\n');
                                }

                                returnString.append('\n');
                                return returnString.toString();
                            }

                            entry = (Entry)i$.next();
                        } while(((String)entry.getKey()).equals(this.defaultLabel) && unclassified == 0);

                        String correctLabel = (String)entry.getKey();
                        int labelTotal = 0;
                        i$ = this.labelMap.keySet().iterator();

                        while(true) {
                            String classifiedLabel;
                            do {
                                if (!i$.hasNext()) {
                                    returnString.append(" |  ").append(StringUtils.rightPad(String.valueOf(labelTotal), 6)).append('\t').append(StringUtils.rightPad(getSmallLabel((Integer)entry.getValue()), 5)).append(" = ").append(correctLabel).append('\n');
                                    continue label47;
                                }

                                classifiedLabel = (String)i$.next();
                            } while(classifiedLabel.equals(this.defaultLabel) && unclassified == 0);

                            returnString.append(StringUtils.rightPad(Integer.toString(this.getCount(correctLabel, classifiedLabel)), 5)).append('\t');
                            labelTotal += this.getCount(correctLabel, classifiedLabel);
                        }
                    }
                }

                entry = (Entry)i$.next();
            } while(((String)entry.getKey()).equals(this.defaultLabel) && unclassified == 0);

            returnString.append(StringUtils.rightPad(getSmallLabel((Integer)entry.getValue()), 5)).append('\t');
        }
    }

    static String getSmallLabel(int i) {
        int val = i;
        StringBuilder returnString = new StringBuilder();

        do {
            int n = val % 26;
            returnString.insert(0, (char)(97 + n));
            val /= 26;
        } while(val > 0);

        return returnString.toString();
    }
}
