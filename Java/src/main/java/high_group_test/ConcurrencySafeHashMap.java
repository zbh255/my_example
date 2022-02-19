package high_group_test;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class ConcurrencySafeHashMap {
    public static void main(String[] args) {
        Map<String,Integer> map = new ConcurrentHashMap<String,Integer>();
        map.put("hello",123);
    }
}
