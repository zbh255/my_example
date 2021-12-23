package stw_gc.code_3_9;

public class ObjectToTenured {
    // 长期存活的对象进入老年代
    private static final int _1MB = 2 << 19;

    /*
    * VM参数:  -XX:+UseSerialGC -verbose:gc -Xms20M -Xmx20M -Xmn10M -XX:+PrintGCDetails -XX:SurvivorRatio=8
    *  -Xlog:gc+age=trace
    *  以下参数二选一: -XX:MaxTenuringThreshold=1 | -XX:MaxTenuringThreshold=15
    */
    @SuppressWarnings("unused")
    public static void main(String[] args) {
        byte[] allocation1,allocation2,allocation3;
        allocation1 = new byte[_1MB / 4];

        allocation2 = new byte[4 * _1MB];
        allocation3 = new byte[4 * _1MB];
        allocation3 = null;
        allocation3 = new byte[4 * _1MB];
    }
}
