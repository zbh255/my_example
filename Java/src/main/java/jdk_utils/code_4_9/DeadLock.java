package jdk_utils.code_4_9;

// 线程死锁等待演示
public class DeadLock {

    static class SynAddRunalbe implements Runnable {
        int a,b;
        public SynAddRunalbe(int a, int b) {
            this.a = a;
            this.b = b;
        }

        @Override
        public void run() {
            synchronized (Integer.valueOf(a)) {
                synchronized (Integer.valueOf(b)) {
                    System.out.println(a + b);
                }
            }
        }
    }

    public static void main(String[] args) {
        for (int i = 0; i < 100; i++) {
            new Thread(new SynAddRunalbe(1,2)).start();
            new Thread(new SynAddRunalbe(2,1)).start();
        }
    }
}
