package patterns.creational.builder.example5;

public class HttpRequest {

    public String url;
    public String method;
    public int timeout;

    public HttpRequest(String url, String method, int timeout) {
        this.url = url;
        this.method = method;
        this.timeout = timeout;
    }

    private HttpRequest(Builder builder) {
        url = builder.url;
        method = builder.method;
        timeout = builder.timeout;
    }

    public static class Builder{
        private String url;
        private String method;
        private int timeout;

        // url is required
        public Builder(String url) {
           this.url = url;
        }

        public Builder usingMethod(String val) {
            method = val;
            return this;
        }

        public Builder withTimeout(int timeout) {
            this.timeout = timeout;
            return this;
        }

        public HttpRequest build(){
            return new HttpRequest(this);
        }
    }

    public static void main(String[] args) {
        HttpRequest request = new HttpRequest.Builder("url")
                .usingMethod("GET")
                .withTimeout(4*100)
                .build();
    }

}
