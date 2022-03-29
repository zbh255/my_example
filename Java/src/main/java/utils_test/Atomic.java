package utils_test;

import java.util.concurrent.*;
import java.util.concurrent.atomic.AtomicInteger;

public class Atomic {
    public static void main(String[] args) {
        AtomicInteger var = new AtomicInteger(1);
        var.compareAndSet()
    }
}
