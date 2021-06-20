import http.server
import random
from prometheus_client import start_http_server, Counter

REQUEST_COUNT = Counter('http_requests_total', 'total app http request count', ['app_name', 'endpoint'])
RANDOM_COUNT = Counter('app_random_count', 'increment counter by random value')

APP_PORT = 8000
METRICS_PORT = 8001


class HandleRequests(http.server.BaseHTTPRequestHandler):

    def do_GET(self):
        REQUEST_COUNT.labels('prom_python_app', self.path).inc()
        random_val = random.random() * 10
        RANDOM_COUNT.inc(random_val)

        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()
        self.wfile.write(bytes(
            """<html>
                <head>
                    <title>First Application</title>
                </head>
                <body style='color: #333; margin-top: 30px;'>
                    <center>
                        <h2>Welcome to our first Prometheus-Python application.</h2>
                    </center>
                </body>
            </html>""", "utf-8"))


if __name__ == "__main__":
    start_http_server(METRICS_PORT)
    server = http.server.HTTPServer(('0.0.0.0', APP_PORT), HandleRequests)
    server.serve_forever()
