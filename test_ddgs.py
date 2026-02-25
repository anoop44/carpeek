from duckduckgo_search import DDGS
import time

def test_search():
    with DDGS() as ddgs:
        print("Searching...")
        results = list(ddgs.images("BMW M3 E30 car site:unsplash.com", max_results=1))
        print("Results:", results)

if __name__ == "__main__":
    test_search()
