package unsafe_method;

import sun.misc.*;

import java.lang.reflect.Field;

public class SimpleExample {
    public static void main(String[] args) throws IllegalAccessException, InterruptedException {
        Field unsafeField = Unsafe.class.getDeclaredFields()[0];
        unsafeField.setAccessible(true);
        Unsafe unsafe = (Unsafe) unsafeField.get(null);

        long address = unsafe.allocateMemory(4);
        unsafe.putInt(address,6666666);
        System.out.println(unsafe.getInt(address));
        // free
        unsafe.freeMemory(address);
        Thread.sleep(2000);
        System.out.println(unsafe.getInt(address));
    }
}
