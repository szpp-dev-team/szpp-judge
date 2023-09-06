import urllib.request

url = "https://example.com"

req = urllib.request.Request(url)
with urllib.request.urlopen(req) as res:
    body: bytes = res.read()

print(body.decode())
