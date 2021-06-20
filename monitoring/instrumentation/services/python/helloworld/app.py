import http.server
import random
import time

from prometheus_client import start_http_server, Counter, Gauge, Summary, Histogram

# Counter
# rate(hello_worlds_total[1m])
# rate(hello_world_exceptions_total[1m])
# rate(hello_world_sales_euro_total[1m])
REQUESTS = Counter('hello_worlds_total', 'Hello Worlds requested', labelnames=['path', 'method'])
EXCEPTIONS = Counter('hello_world_exceptions_total', 'Exceptions serving Hello World.')
SALES = Counter('hello_world_sales_euro_total', 'Euros made serving Hello World.')

# Gauge
INPROGRESS = Gauge('hello_worlds_inprogress', 'Number of Hello Worlds in progress.')
LAST = Gauge('hello_world_last_time_seconds', 'The last time a Hello World was served.')

# summary
# LATENCY = Summary('hello_world_latency_seconds', 'Time for a request Hello World.')

# histogram
LATENCY = Histogram('hello_world_latency_seconds', 'Time for a request Hello World.')


class MyHandler(http.server.BaseHTTPRequestHandler):

    @INPROGRESS.track_inprogress()
    @EXCEPTIONS.count_exceptions()
    @LATENCY.time()
    def do_GET(self):
        REQUESTS.labels(self.path, self.command).inc()
        euros = random.random()
        # time.sleep(5)
        SALES.inc(euros)
        if random.random() < 0.2:
            raise Exception
        self.send_response(200)
        self.end_headers()
        self.wfile.write(b'Hello World')
        LAST.set_to_current_time()


if __name__ == '__main__':
    start_http_server(8000)
    server = http.server.HTTPServer(('0.0.0.0', 8001), MyHandler)
    server.serve_forever()
