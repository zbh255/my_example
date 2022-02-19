package high_group_test.api_limiting;

import org.junit.Test;

import java.util.Map;
import java.util.Random;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.CountDownLatch;

public class TestTb {

    @Test
    // 模拟测试单机的情况
    public void testSingleTb() {
        TokenBucket singleTb = new TokenBucket(20,1,2);
        CountDownLatch group = new CountDownLatch(singleTb.getCap() * 2);
        for (int i = 0; i < singleTb.getCap() * 2; i++) {
            Thread t = new Thread(() -> {
                try {
                    if (singleTb.takeToken()) System.out.println("我拿到令牌了");
                    else System.out.println("我没拿到令牌");
                } finally {
                    group.countDown();
                }
            });
            t.start();
        }
        try {
            group.await();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

    @Test
    // 模拟测试多机环境
    public void testMultiTb() {
        Map<Integer,TokenBucket> multiTb = new ConcurrentHashMap<>();
        Random rand = new Random();
        CountDownLatch group = new CountDownLatch(10 * 20);
        for (int i = 0; i < 10; i++) {
            TokenBucket tb = new TokenBucket(20,1,2);
            multiTb.put(i,tb);
        }
        for (int i = 0; i < 10 * 20; i++) {
            Thread t = new Thread(() -> {
                Integer n = rand.nextInt(10);
                TokenBucket tb = multiTb.get(n);
                try {
                    if (tb.takeToken()) System.out.printf("访问第%d台机器成功\n", n);
                    else System.out.printf("访问第%d台机器失败\n", n);
                } finally {
                    group.countDown();
                }
            });
            t.start();
        }
        try {
            group.await();
            //0 : TokenBucketCap(0)
            //1 : TokenBucketCap(2)
            //2 : TokenBucketCap(4)
            //3 : TokenBucketCap(3)
            //4 : TokenBucketCap(5)
            //5 : TokenBucketCap(2)
            //6 : TokenBucketCap(0)
            //7 : TokenBucketCap(0)
            //8 : TokenBucketCap(7)
            //9 : TokenBucketCap(0)
            for (Map.Entry<Integer, TokenBucket> entry : multiTb.entrySet()) {
                System.out.printf("%d : TokenBucketCap(%d)\n",entry.getKey(),entry.getValue().getOldCap());
            }
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}
