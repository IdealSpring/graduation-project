package org.apache.mahout.classifier.df.node;

import org.apache.mahout.classifier.df.data.Instance;

import java.io.DataInput;
import java.io.DataOutput;
import java.io.IOException;

/**
 * 表示叶节点
 *
 * @author 李程鸿
 */
public class Leaf extends Node {
    private static final double EPSILON = 1.0e-6;

    private double label;

    Leaf() { }

    public Leaf(double label) {
        this.label = label;
    }

    @Override
    public double classify(Instance instance) {
        return label;
    }

    @Override
    public long maxDepth() {
        return 1;
    }

    @Override
    public long nbNodes() {
        return 1;
    }

    @Override
    public Type getType() {
        return Type.LEAF;
    }

    @Override
    public boolean equals(Object obj) {
        if (this == obj) {
            return true;
        }
        if (!(obj instanceof Leaf)) {
            return false;
        }

        Leaf leaf = (Leaf) obj;

        return Math.abs(label - leaf.label) < EPSILON;
    }

    @Override
    public int hashCode() {
        long bits = Double.doubleToLongBits(label);
        return (int)(bits ^ (bits >>> 32));
    }

    @Override
    protected String getString() {
        return "";
    }

    @Override
    public void readFields(DataInput in) throws IOException {
        label = in.readDouble();
    }

    @Override
    protected void writeNode(DataOutput out) throws IOException {
        out.writeDouble(label);
    }

    public static double getEPSILON() {
        return EPSILON;
    }

    public double getLabel() {
        return label;
    }

    public void setLabel(double label) {
        this.label = label;
    }
}
