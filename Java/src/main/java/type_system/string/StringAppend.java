package type_system.string;

import org.junit.Test;

// 研究Java的字符串拼接方式
public class StringAppend {
    @Test
    public void immutable() {}
    @Test
    public void StringBuilder() {
        StringBuilder sb = new StringBuilder();
        sb.append("hh");
    }
    @Test
    public void StringBuffer() {}

    public static void main(String[] args) {

    }
}
