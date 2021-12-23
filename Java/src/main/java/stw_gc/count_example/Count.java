package stw_gc.count_example;

public class Count {
    static class ReferenceCountingGC {
        public Object instance = null;
        private static final int _1MB = 1024 * 1024;

        private byte[] bigSize = new byte[2 * _1MB];

        public static void testGC() {
            ReferenceCountingGC obja = new ReferenceCountingGC();
            ReferenceCountingGC objb = new ReferenceCountingGC();
            obja.instance = objb;
            objb.instance = obja;

            objb = null;
            obja = null;
            System.gc();
        }
    }

    public static void main(String[] args) {
        ReferenceCountingGC.testGC();
    }
}