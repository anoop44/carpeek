import urllib.request
import urllib.parse
import re

def search_unsplash(query):
    url = "https://unsplash.com/s/photos/" + urllib.parse.quote(query)
    req = urllib.request.Request(url, headers={'User-Agent': 'Mozilla/5.0'})
    try:
        html = urllib.request.urlopen(req).read().decode('utf-8')
        # match image url
        images = re.findall(r'src="(https://images\.unsplash\.com/photo-[^"]+)"', html)
        if images:
            return images[0]
    except Exception as e:
        print("Error:", e)
    return None

print(search_unsplash("bmw m3"))
