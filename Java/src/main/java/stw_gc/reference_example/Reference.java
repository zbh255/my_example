package stw_gc.reference_example;

import org.junit.Test;

import java.lang.ref.PhantomReference;
import java.lang.ref.ReferenceQueue;
import java.lang.ref.SoftReference;
import java.lang.ref.WeakReference;
import java.util.ArrayList;
import java.util.List;

public class Reference {

    static class ReferenceTest {
        @Override
        protected void finalize() throws Throwable {
            System.out.println("我被回收了");
        }
    }

    @Test
    // 软引用
    public void Soft() throws InterruptedException {
        SoftReference<byte[]> sr = new SoftReference<>(new byte[1024 * 1024 * 10]);
        System.out.println(sr.get());
        System.gc();
        Thread.sleep(500);
        System.out.println(sr.get());
        // 设置
        byte[] bytes = new byte[1024 * 1024 * 15];
        System.out.println(sr.get());
    }

    @Test
    // 弱引用
    public void Weak() throws InterruptedException {
        WeakReference<byte[]> wr = new WeakReference<>(new byte[1024 * 1024 * 10]);
        System.out.println(wr.get());
        System.gc();
        Thread.sleep(500);
        System.out.println(wr.get());
        // 设置
        byte[] bytes = new byte[1024 * 1024 * 15];
        System.out.println(wr.get());
    }

    @Test
    // 虚引用
    public void Phantom() throws InterruptedException {
        List<byte[]> list = new ArrayList<>();
        // 引用队列
        ReferenceQueue<ReferenceTest> queue = new ReferenceQueue<>();
        // 虚引用
        PhantomReference<ReferenceTest> pr = new PhantomReference<>(new ReferenceTest(),queue);

        System.out.println(pr.get());
        new Thread(() -> {
            while (true) {
                list.add(new byte[1024 * 1024]);
                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
                System.out.println(pr.get());
            }
        }).start();

        new Thread(() -> {
            while (true) {
                java.lang.ref.Reference<? extends ReferenceTest> poll = queue.poll();
                if (poll != null) {
                    System.out.println("虚引用对象被回收了" + poll);
                }
            }
        }).start();

        Thread.sleep(500);
    }
}
