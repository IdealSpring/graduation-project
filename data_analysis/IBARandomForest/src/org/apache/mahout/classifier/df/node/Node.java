package org.apache.mahout.classifier.df.node;

import org.apache.hadoop.io.Writable;
import org.apache.mahout.classifier.df.data.Instance;

import java.io.DataInput;
import java.io.DataOutput;
import java.io.IOException;

/**
 * 表示决策树的抽象节点
 *
 * @author 李程鸿
 */
public abstract class Node implements Writable {

    public enum Type {
        LEAF,
        NUMERICAL,
        CATEGORICAL
    }

    /**
     * 预测实例的标签
     *
     * @return
     *      如果标签无法预测，则为-1
     */
    public abstract double classify(Instance instance);

    /**
     * @return
     *      树的节点总数
     */
    public abstract long nbNodes();

    /**
     * @return
     *      树的最大深度
     */
    public abstract long maxDepth();

    public abstract Type getType();

    public static Node read(DataInput in) throws IOException {
        Type type = Type.values()[in.readInt()];
        Node node;

        switch (type) {
            case LEAF:
                node = new Leaf();
                break;
            case NUMERICAL:
                node = new NumericalNode();
                break;
            case CATEGORICAL:
                node = new CategoricalNode();
                break;
            default:
                throw new IllegalStateException("This implementation is not currently supported");
        }

        node.readFields(in);

        return node;
    }

    @Override
    public final String toString() {
        return getType() + ":" + getString() + ';';
    }

    protected abstract String getString();

    @Override
    public final void write(DataOutput out) throws IOException {
        out.writeInt(getType().ordinal());
        writeNode(out);
    }

    protected abstract void writeNode(DataOutput out) throws IOException;

}
