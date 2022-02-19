package high_group_test.api_limiting;

// 基于令牌桶的限流方式
public class TokenBucket {
    // 令牌桶的容量
    private final int cap;
    // 添加令牌的间隔时间，如: 间隔一秒钟添加一个令牌
    private final int addTime;
    // 添加令牌的个数
    private final int addTokens;

    // 上一次添加令牌的时间
    private volatile long oldAddTime;
    // 上一次的令牌容量
    private volatile int oldCap;

    public TokenBucket(int cap, int addTime, int addTokens) {
        this.oldCap = cap;
        this.cap = cap;
        this.addTime = addTime;
        this.addTokens = addTokens;
    }

    public int getCap() {
        return this.cap;
    }

    public int getOldCap() {
        return this.oldCap;
    }

    public synchronized boolean takeToken() {
        if (this.oldAddTime == 0) {
            this.oldAddTime = System.currentTimeMillis() / 1000;
        }
        long newTime = System.currentTimeMillis() / 1000;
        int cur = 0;
        try {
            cur = oldCap + ((int)(newTime - oldAddTime) / addTime) * addTokens;
            cur = cur > this.cap ? this.cap : cur;
            return cur > 0;
        } finally {
            if (cur > 0) cur--;
            this.oldCap = cur;
            this.oldAddTime = newTime;
        }
    }
}
