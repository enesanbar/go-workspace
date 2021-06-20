package patterns.creational.builder.example;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class HtmlBuilderTest {

    @BeforeEach
    void setUp() {

    }

    @Test
    void buildAParagraphElementWithStringBuilder() {
        String hello = "hello";

        StringBuilder sb = new StringBuilder();
        sb.append("<p>")
                .append(hello)
                .append("</p>");

        System.out.println(sb);

        assertEquals("<p>hello</p>", sb.toString());
    }

    @Test
    void buildAListElementWithStringBuilder() {
        String [] words = {"hello", "world"};

        StringBuilder sb = new StringBuilder();

        sb.append("<ul>\n");
        for (String word: words) {
            // indentation management, line breaks and other evils
            sb.append(String.format("  <li>%s</li>\n", word));
        }
        sb.append("</ul>");
        System.out.println(sb);

        assertEquals("<ul>\n  <li>hello</li>\n  <li>world</li>\n</ul>", sb.toString());
    }

    @Test
    void buildAListElementWithHtmlBuilder() {
        // ordinary non-fluent builder
        HtmlBuilder builder = new HtmlBuilder("ul");
        builder.addChild("li", "hello");
        builder.addChild("li", "world");
        System.out.println(builder);

        assertEquals("" +
                "<ul>\n" +
                "    <li>\n" +
                "        hello\n" +
                "    </li>\n" +
                "    <li>\n" +
                "        world\n" +
                "    </li>\n" +
                "</ul>\n", builder.toString());
    }

    @Test
    void buildAListElementWithFluentHtmlBuilder() {
        HtmlBuilder builder = new HtmlBuilder("ul");

        // fluent builder
        builder.clear();
        builder.addChildFluent("li", "hello")
                .addChildFluent("li", "world");
        System.out.println(builder);

        assertEquals("" +
                "<ul>\n" +
                "    <li>\n" +
                "        hello\n" +
                "    </li>\n" +
                "    <li>\n" +
                "        world\n" +
                "    </li>\n" +
                "</ul>\n", builder.toString());
    }


}