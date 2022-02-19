package algorithm.code_1_1_20;


import edu.princeton.cs.algs4.StdDraw;
import edu.princeton.cs.algs4.StdRandom;

public class RandomArray {
    public static void main(String[] args) {
        int N = 50;
        double[] a = new double[N];
        for (int i = 0; i < N; i++)
            a[i] = StdRandom.random();
        for (int i = 0; i < N; i++)
        {
            double x = 1.0 * i / N;
            double y = a[i] / 2.0;
            double rw = 0.5 / N;
            double rh = a[i] / 2.0;
            StdDraw.filledRectangle(x,y,rw,rh);
        }
    }
}
