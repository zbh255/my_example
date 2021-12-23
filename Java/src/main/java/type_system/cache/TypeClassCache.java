package type_system.cache;

public class TypeClassCache {
    // 演示Java包装类型对基本类型数据的缓存
    // -128~127
    public static void main(String[] args) {
        Long a = 127L;
        Long b = 127L;
        System.out.printf("type long cache is a==b :%s\n",a == b);

        Long a2 = 128L;
        Long b2 = 128L;
        System.out.printf("type long cache value=128 is a==b :%s\n",a2 == b2);
    }
}
