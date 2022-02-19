package high_group_test;

import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.ThreadFactory;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;

public class ThreadPool {
    public static void main(String[] args) {
        ArrayBlockingQueue<Runnable> queue = new ArrayBlockingQueue<>(100);
        ThreadPoolExecutor pool = new ThreadPoolExecutor(100,
                300,
                10,
                TimeUnit.SECONDS,
                queue,
                new ThreadFactory() {
                    @Override
                    public Thread newThread(Runnable r) {
                        return null;
                    }
                },
                new ThreadPoolExecutor.CallerRunsPolicy());
    }
}
