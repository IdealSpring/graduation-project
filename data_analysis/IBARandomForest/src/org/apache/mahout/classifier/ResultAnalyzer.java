package org.apache.mahout.classifier;

import org.apache.commons.lang3.StringUtils;
import org.apache.mahout.cf.taste.impl.common.RunningAverageAndStdDev;
import org.apache.mahout.math.stats.OnlineSummarizer;

import java.text.DecimalFormat;
import java.text.NumberFormat;
import java.util.Collection;

/**
 * 结果分析器捕获分类统计并以表格方式显示。
 *
 * @author 李程鸿
 * */
public class ResultAnalyzer {

    private final ConfusionMatrix confusionMatrix;
    private final OnlineSummarizer summarizer;
    private boolean hasLL;

    private int correctlyClassified;
    private int incorrectlyClassified;

    public ResultAnalyzer(Collection<String> labelSet, String defaultLabel) {
        confusionMatrix = new ConfusionMatrix(labelSet, defaultLabel);
        summarizer = new OnlineSummarizer();
    }

    public ConfusionMatrix getConfusionMatrix() {
        return this.confusionMatrix;
    }

    /**
     * 增加实例
	 *
     * @param correctLabel
     *          正确的标签
     * @param classifiedResult
     *          分类结果
     * @return 实例是否正确
     */
    public boolean addInstance(String correctLabel, ClassifierResult classifiedResult) {
        boolean result = correctLabel.equals(classifiedResult.getLabel());
        if (result) {
            correctlyClassified++;
        } else {
            incorrectlyClassified++;
        }
        confusionMatrix.addInstance(correctLabel, classifiedResult);
        if (classifiedResult.getLogLikelihood() != Double.MAX_VALUE) {
            summarizer.add(classifiedResult.getLogLikelihood());
            hasLL = true;
        }
        return result;
    }

    @Override
    public String toString() {
        StringBuilder returnString = new StringBuilder();

        returnString.append('\n');
        returnString.append("=======================================================\n");
        returnString.append("Summary\n");
        returnString.append("-------------------------------------------------------\n");
        int totalClassified = correctlyClassified + incorrectlyClassified;
        double percentageCorrect = (double) 100 * correctlyClassified / totalClassified;
        double percentageIncorrect = (double) 100 * incorrectlyClassified / totalClassified;
        NumberFormat decimalFormatter = new DecimalFormat("0.####");

        returnString.append(StringUtils.rightPad("Correctly Classified Instances", 40)).append(": ").append(
                StringUtils.leftPad(Integer.toString(correctlyClassified), 10)).append('\t').append(
                StringUtils.leftPad(decimalFormatter.format(percentageCorrect), 10)).append("%\n");
        returnString.append(StringUtils.rightPad("Incorrectly Classified Instances", 40)).append(": ").append(
                StringUtils.leftPad(Integer.toString(incorrectlyClassified), 10)).append('\t').append(
                StringUtils.leftPad(decimalFormatter.format(percentageIncorrect), 10)).append("%\n");
        returnString.append(StringUtils.rightPad("Total Classified Instances", 40)).append(": ").append(
                StringUtils.leftPad(Integer.toString(totalClassified), 10)).append('\n');
        returnString.append('\n');

        returnString.append(confusionMatrix);
        returnString.append("=======================================================\n");
        returnString.append("Statistics\n");
        returnString.append("-------------------------------------------------------\n");

        RunningAverageAndStdDev normStats = confusionMatrix.getNormalizedStats();
        returnString.append(StringUtils.rightPad("Kappa", 40)).append(
                StringUtils.leftPad(decimalFormatter.format(confusionMatrix.getKappa()), 10)).append('\n');
        returnString.append(StringUtils.rightPad("Accuracy", 40)).append(
                StringUtils.leftPad(decimalFormatter.format(confusionMatrix.getAccuracy()), 10)).append("%\n");
        returnString.append(StringUtils.rightPad("Reliability", 40)).append(
                StringUtils.leftPad(decimalFormatter.format(normStats.getAverage() * 100.00000001), 10)).append("%\n");
        returnString.append(StringUtils.rightPad("Reliability (standard deviation)", 40)).append(
                StringUtils.leftPad(decimalFormatter.format(normStats.getStandardDeviation()), 10)).append('\n');
        returnString.append(StringUtils.rightPad("Weighted precision", 40)).append(
                StringUtils.leftPad(decimalFormatter.format(confusionMatrix.getWeightedPrecision()), 10)).append('\n');
        returnString.append(StringUtils.rightPad("Weighted recall", 40)).append(
                StringUtils.leftPad(decimalFormatter.format(confusionMatrix.getWeightedRecall()), 10)).append('\n');
        returnString.append(StringUtils.rightPad("Weighted F1 score", 40)).append(
                StringUtils.leftPad(decimalFormatter.format(confusionMatrix.getWeightedF1score()), 10)).append('\n');

        if (hasLL) {
            returnString.append(StringUtils.rightPad("Log-likelihood", 30)).append("mean      : ").append(
                    StringUtils.leftPad(decimalFormatter.format(summarizer.getMean()), 10)).append('\n');
            returnString.append(StringUtils.rightPad("", 30)).append(StringUtils.rightPad("25%-ile   : ", 10)).append(
                    StringUtils.leftPad(decimalFormatter.format(summarizer.getQuartile(1)), 10)).append('\n');
            returnString.append(StringUtils.rightPad("", 30)).append(StringUtils.rightPad("75%-ile   : ", 10)).append(
                    StringUtils.leftPad(decimalFormatter.format(summarizer.getQuartile(3)), 10)).append('\n');
        }

        return returnString.toString();
    }

	/**
	 * 得到F1-score
	 *
	 * @return
	 */
	public String getWeightedF1score() {
        NumberFormat decimalFormatter = new DecimalFormat("0.####");
        String weightedF1score = StringUtils.leftPad(decimalFormatter.format(confusionMatrix.getWeightedF1score()), 10);
        return weightedF1score;
    }
}
