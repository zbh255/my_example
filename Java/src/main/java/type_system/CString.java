package type_system;

import java.util.stream.IntStream;

public final class CString implements java.io.Serializable, Comparable<CString>, CharSequence {
    @Override
    public int length() {
        return 0;
    }

    @Override
    public char charAt(int index) {
        return 0;
    }

    @Override
    public CharSequence subSequence(int start, int end) {
        return null;
    }

    @Override
    public IntStream chars() {
        return null;
    }

    @Override
    public IntStream codePoints() {
        return null;
    }

    @Override
    public int compareTo(CString o) {
        return 0;
    }

}
