package stw_gc.code_3_8;

public class BigObjectToTenured {
    // 大对象直接进入老年代
    private static final int _1MB = 2 << 19;

    /**
     *  VM参数: -XX:+UseSerialGC -verbose:gc -Xms20M -Xmx20M -Xmn10M -XX:+PrintGCDetails -XX:SurvivorRatio=8
     *  -XX:PretenureSizeThreshold=3145782
     */
    public static void main(String[] args) {
        byte[] allocation;
        allocation = new byte[4 * _1MB]; // 直接分配在老年代中
    }
}
