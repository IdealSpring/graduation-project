package org.apache.mahout.classifier.df.node;

import org.apache.mahout.classifier.df.data.Instance;

import java.io.DataInput;
import java.io.DataOutput;
import java.io.IOException;

/**
 * 表示使用数值属性拆分的节点
 *
 * @author 李程鸿
 */
public class NumericalNode extends Node {
    // 分裂的数值属性
    private int attr;

    // 分裂价值
    private double split;

    // 属性值<拆分值时的子节点
    private Node loChild;

    // 子节点当属性值>＝分裂值时
    private Node hiChild;

    public NumericalNode() { }

    public NumericalNode(int attr, double split, Node loChild, Node hiChild) {
        this.attr = attr;
        this.split = split;
        this.loChild = loChild;
        this.hiChild = hiChild;
    }

    @Override
    public double classify(Instance instance) {
        if (instance.get(attr) < split) {
            return loChild.classify(instance);
        } else {
            return hiChild.classify(instance);
        }
    }

    @Override
    public long maxDepth() {
        return 1 + Math.max(loChild.maxDepth(), hiChild.maxDepth());
    }

    @Override
    public long nbNodes() {
        return 1 + loChild.nbNodes() + hiChild.nbNodes();
    }

    @Override
    public Type getType() {
        return Type.NUMERICAL;
    }

    @Override
    public boolean equals(Object obj) {
        if (this == obj) {
            return true;
        }
        if (!(obj instanceof NumericalNode)) {
            return false;
        }

        NumericalNode node = (NumericalNode) obj;

        return attr == node.attr && split == node.split && loChild.equals(node.loChild) && hiChild.equals(node.hiChild);
    }

    @Override
    public int hashCode() {
        return attr + (int) Double.doubleToLongBits(split) + loChild.hashCode() + hiChild.hashCode();
    }

    @Override
    protected String getString() {
        return loChild.toString() + ',' + hiChild.toString();
    }

    @Override
    public void readFields(DataInput in) throws IOException {
        attr = in.readInt();
        split = in.readDouble();
        loChild = Node.read(in);
        hiChild = Node.read(in);
    }

    @Override
    protected void writeNode(DataOutput out) throws IOException {
        out.writeInt(attr);
        out.writeDouble(split);
        loChild.write(out);
        hiChild.write(out);
    }



    public int getAttr() {
        return attr;
    }

    public void setAttr(int attr) {
        this.attr = attr;
    }

    public double getSplit() {
        return split;
    }

    public void setSplit(double split) {
        this.split = split;
    }

    public Node getLoChild() {
        return loChild;
    }

    public void setLoChild(Node loChild) {
        this.loChild = loChild;
    }

    public Node getHiChild() {
        return hiChild;
    }

    public void setHiChild(Node hiChild) {
        this.hiChild = hiChild;
    }
}
